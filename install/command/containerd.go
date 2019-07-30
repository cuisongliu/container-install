package command

import (
	"bytes"
	"fmt"
	"github.com/wonderivan/logger"
	"text/template"
)

type Containerd struct{}

func NewContainerd() stepInterface {
	var stepInterface stepInterface
	stepInterface = &Containerd{}
	return stepInterface
}

const containerdFileName = "containerd.tgz"

func (d Containerd) SendPackage(host string) {
	SendPackage(host, PkgUrl, containerdFileName)
}

func (d Containerd) Tar(host string) {
	cmd := fmt.Sprintf("tar --strip-components=1 -xvzf /root/%s -C /usr/local/bin", containerdFileName)
	Cmd(host, cmd)
}

func (d Containerd) Config(host string) {
	cmd := "mkdir -p " + Lib
	Cmd(host, cmd)
	cmd = "mkdir -p /etc/containerd"
	Cmd(host, cmd)
	cmd = "containerd config default > /etc/containerd/config.toml"
	//cmd = "echo \"" + string(d.configFile()) + "\" > /etc/docker/daemon.json"
	Cmd(host, cmd)
}

func (d Containerd) Enable(host string) {
	cmd := "echo \"" + string(d.serviceFile()) + "\" > /usr/lib/systemd/system/containerd.service"
	Cmd(host, cmd)
	cmd = "systemctl enable  containerd.service && systemctl restart  containerd.service"
	Cmd(host, cmd)
}

func (d Containerd) Version(host string) {
	cmd := "containerd --version"
	Cmd(host, cmd)
}

func (d Containerd) Uninstall(host string) {
	cmd := "systemctl stop  containerd.service && systemctl disable containerd.service"
	Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/ctr && rm -rf /usr/local/bin/containerd* "
	Cmd(host, cmd)
	cmd = "rm -rf /var/lib/containerd && rm -rf /etc/containerd/* "
	Cmd(host, cmd)
	if Lib != "" {
		cmd = "rm -rf " + Lib
		Cmd(host, cmd)
	}
}

func (d Containerd) serviceFile() []byte {
	var templateText = string(`[Unit]
Description=containerd container runtime
Documentation=https://containerd.io
After=network.target
  
[Service]
ExecStart=/usr/local/bin/containerd
Restart=always
RestartSec=5
Delegate=yes
KillMode=process
OOMScoreAdjust=-999
LimitNOFILE=1048576
LimitNPROC=infinity
LimitCORE=infinity
  
[Install]
WantedBy=multi-user.target
`)
	return []byte(templateText)
}
func (d Containerd) configFile() []byte {
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
	envMap["DOCKER_LIB"] = Lib
	envMap["ZERO"] = 0
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	return buffer.Bytes()
}

func (d Containerd) Print() {
	urlPrefix := "https://github.com/containerd/containerd/releases/download/v%s/containerd-%s.linux-amd64.tar.gz"
	versions := []string{
		"1.1.0",
		"1.1.1",
		"1.1.2",
		"1.1.3",
		"1.1.4",
		"1.1.5",
		"1.1.6",
		"1.1.7",

		"1.2.0",
		"1.2.1",
		"1.2.2",
		"1.2.3",
		"1.2.4",
		"1.2.5",
		"1.2.6",
		"1.2.7",
	}

	for _, v := range versions {
		println(fmt.Sprintf(urlPrefix, v, v))
	}
}
