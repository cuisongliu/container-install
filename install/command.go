package install

import (
	"bytes"
	"fmt"
	"github.com/wonderivan/logger"
	"strings"
	"text/template"
)

func tarDocker(host string) {
	cmd := fmt.Sprintf("tar --strip-components=1 -xvzf /root/%s -C /usr/local/bin", fileName)
	Cmd(host, cmd)
}
func configDocker(host string) {
	cmd := "mkdir -p " + DockerLib
	Cmd(host, cmd)
	cmd = "mkdir -p /etc/docker"
	Cmd(host, cmd)
	cmd = "echo \"" + string(dockerConfig(RegistryArr, DockerLib)) + "\" > /etc/docker/daemon.json"
	Cmd(host, cmd)
}
func enableDocker(host string) {
	cmd := "echo \"" + string(dockerServiceFile()) + "\" > /usr/lib/systemd/system/docker.service"
	Cmd(host, cmd)
	cmd = "systemctl enable  docker.service && systemctl restart  docker.service"
	Cmd(host, cmd)
}

func versionDocker(host string) {
	cmd := "docker version"
	Cmd(host, cmd)
}

func uninstallDocker(host string) {
	cmd := "systemctl stop  docker.service && systemctl disable docker.service"
	Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/docker* && rm -rf /var/lib/docker && rm -rf /etc/docker/* "
	Cmd(host, cmd)
	if DockerLib != "" {
		cmd = "rm -rf " + DockerLib
		Cmd(host, cmd)
	}
}

func dockerServiceFile() []byte {
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

func dockerConfig(registryArr []string, dir string) []byte {
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
	envMap["DOCKER_REGISTRY"] = registryArr
	envMap["DOCKER_LIB"] = dir
	envMap["ZERO"] = 0
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	return buffer.Bytes()
}

func registryJoin(registryArr []string) string {
	var sb strings.Builder
	for i, v := range registryArr {
		if i != 0 {
			sb.Write([]byte(","))
		}
		sb.Write([]byte("\""))
		sb.Write([]byte(v))
		sb.Write([]byte("\""))
	}
	return sb.String()
}
