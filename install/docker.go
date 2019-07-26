package install

import (
	"bytes"
	"fmt"
	"github.com/wonderivan/logger"
	"text/template"
)

type Docker struct{}

func newDocker() stepInterface {
	var stepInterface stepInterface
	stepInterface = &Docker{}
	return stepInterface
}

func (d Docker) tar(host string) {
	cmd := fmt.Sprintf("tar --strip-components=1 -xvzf /root/%s -C /usr/local/bin", dockerFileName)
	Cmd(host, cmd)
}

func (d Docker) config(host string) {
	cmd := "mkdir -p " + Lib
	Cmd(host, cmd)
	cmd = "mkdir -p /etc/docker"
	Cmd(host, cmd)
	cmd = "echo \"" + string(d.configFile()) + "\" > /etc/docker/daemon.json"
	Cmd(host, cmd)
}

func (d Docker) enable(host string) {
	cmd := "echo \"" + string(d.serviceFile()) + "\" > /usr/lib/systemd/system/docker.service"
	Cmd(host, cmd)
	cmd = "systemctl enable  docker.service && systemctl restart  docker.service"
	Cmd(host, cmd)
}

func (d Docker) version(host string) {
	cmd := "docker version"
	Cmd(host, cmd)
}

func (d Docker) uninstall(host string) {
	cmd := "systemctl stop  docker.service && systemctl disable docker.service"
	Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/runc && rm -rf /usr/local/bin/ctr && rm -rf /usr/local/bin/containerd* "
	Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/docker* && rm -rf /var/lib/docker && rm -rf /etc/docker/* "
	Cmd(host, cmd)
	if Lib != "" {
		cmd = "rm -rf " + Lib
		Cmd(host, cmd)
	}
}

func (d Docker) serviceFile() []byte {
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
func (d Docker) configFile() []byte {
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
