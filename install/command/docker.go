package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	sealos "github.com/fanux/sealos/install"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Docker struct{}

func NewDocker() stepInterface {
	var stepInterface stepInterface
	stepInterface = &Docker{}
	return stepInterface
}

const dockerFileName = "docker.tgz"

func (d *Docker) lib() string {
	if Lib == "" {
		return "/var/lib/docker"
	} else {
		return Lib
	}
}

func (d *Docker) SendPackage(host string) {
	sendPackage(host, PkgUrl, dockerFileName)
}

func (d *Docker) Tar(host string) {
	cmd := fmt.Sprintf("tar --strip-components=1 -xvzf /root/%s -C /usr/local/bin", dockerFileName)
	sealos.Cmd(host, cmd)
}

func (d *Docker) Config(host string) {
	cmd := "mkdir -p " + d.lib()
	sealos.Cmd(host, cmd)
	cmd = "mkdir -p /etc/docker"
	sealos.Cmd(host, cmd)
	cmd = "echo \"" + string(d.configFile()) + "\" > /etc/docker/daemon.json"
	sealos.Cmd(host, cmd)
}

func (d *Docker) Enable(host string) {
	cmd := "echo \"" + string(d.serviceFile()) + "\" > /usr/lib/systemd/system/docker.service"
	sealos.Cmd(host, cmd)
	cmd = "systemctl enable  docker.service && systemctl restart  docker.service"
	sealos.Cmd(host, cmd)
}

func (d *Docker) Version(host string) {
	cmd := "docker version"
	sealos.Cmd(host, cmd)
}

func (d *Docker) Uninstall(host string) {
	cmd := "systemctl stop  docker.service && systemctl disable docker.service"
	sealos.Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/runc && rm -rf /usr/local/bin/ctr && rm -rf /usr/local/bin/containerd* "
	sealos.Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/docker* && rm -rf /var/lib/docker && rm -rf /etc/docker/* "
	sealos.Cmd(host, cmd)
	if d.lib() != "" {
		cmd = "rm -rf " + d.lib()
		sealos.Cmd(host, cmd)
	}
}

func (d *Docker) serviceFile() []byte {
	var templateText = string(`[Unit]
Description=Docker Application Container Engine
Documentation=https://docs.docker.com
After=network.target

[Service]
Type=notify
ExecStart=/usr/local/bin/dockerd
ExecReload=/bin/kill -s HUP $MAINPID
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
TimeoutStartSec=0
Delegate=yes
KillMode=process

[Install]
WantedBy=multi-user.target
`)
	return []byte(templateText)
}
func (d *Docker) configFile() []byte {
	var templateText = string(`{
  \"registry-mirrors\": [
     \"http://373a6594.m.daocloud.io\"
  ],
  {{if len .DOCKER_REGISTRY}}
  \"insecure-registries\":
        [{{range $i,$v :=.DOCKER_REGISTRY}}{{if eq $i  0}}\"{{$v}}\"{{else}},\"{{$v}}\"{{end}}{{end}}],
  {{end}}
  \"graph\":\"{{.DOCKER_LIB}}\"
}`)
	tmpl, err := template.New("text").Parse(templateText)
	if err != nil {
		logger.Error("template parse failed:", err)
		panic(1)
	}
	var envMap = make(map[string]interface{})
	envMap["DOCKER_REGISTRY"] = RegistryArr
	envMap["DOCKER_LIB"] = d.lib()
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	return buffer.Bytes()
}

func (d *Docker) Print() {
	urlPrefix := "https://download.docker.com/linux/static/stable/x86_64/%s"
	data, err := Asset("install/command/docker.json")
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	var versions []string
	_ = json.Unmarshal(data, &versions)
	for _, v := range versions {
		println(fmt.Sprintf(urlPrefix, v))
	}
}

func (d *Docker) Fetch() {
	url := "https://download.docker.com/linux/static/stable/x86_64/"
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		logger.Error("http code is not 200")
		os.Exit(1)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	ahtml := doc.Find("a")
	var versions []string
	for _, html := range ahtml.Nodes {
		attr := html.Attr
		if len(attr) > 0 {
			if strings.Contains(attr[0].Val, "docker") {
				versions = append(versions, attr[0].Val)
				logger.Debug("加入缓存值：%s", attr[0].Val)
			}
		}
	}
	//
	logger.Debug("写入json文件")
	dockerJson, err := os.OpenFile("install/command/docker.json", os.O_WRONLY, 0755)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	defer dockerJson.Close()
	//json解析
	encoder := json.NewEncoder(dockerJson)
	err = encoder.Encode(&versions)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	} else {
		logger.Info("构建成功")
	}

}
