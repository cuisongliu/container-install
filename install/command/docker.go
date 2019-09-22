package command

import (
	"bytes"
	"fmt"
	sealos "github.com/fanux/sealos/install"
	"github.com/wonderivan/logger"
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
	urlPrefix := "https://download.docker.com/linux/static/stable/x86_64/docker-%s.tgz"
	versions := []string{
		"17.03.0-ce",
		"17.03.1-ce",
		"17.03.2-ce",
		"17.06.0-ce",
		"17.06.1-ce",
		"17.06.2-ce",
		"17.09.0-ce",
		"17.09.1-ce",
		"17.12.0-ce",
		"17.12.1-ce",
		"18.03.0-ce",
		"18.03.1-ce",
		"18.06.0-ce",
		"18.06.1-ce",
		"18.06.2-ce",
		"18.06.3-ce",
		"18.09.0",
		"18.09.1",
		"18.09.2",
		"18.09.3",
		"18.09.4",
		"18.09.5",
		"18.09.6",
		"18.09.7",
		"18.09.8",
		"19.03.0",
	}

	for _, v := range versions {
		println(fmt.Sprintf(urlPrefix, v))
	}
}
