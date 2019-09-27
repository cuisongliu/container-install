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
	"text/template"
)

type Containerd struct{}

func NewContainerd() stepInterface {
	var stepInterface stepInterface
	stepInterface = &Containerd{}
	return stepInterface
}

const containerdFileName = "containerd.tgz"

func (d *Containerd) lib() string {
	if Lib == "" {
		return "/var/lib/containerd"
	} else {
		return Lib
	}
}
func (d *Containerd) SendPackage(host string) {
	sendPackage(host, PkgUrl, containerdFileName)
}

func (d *Containerd) Tar(host string) {
	cmd := fmt.Sprintf("tar --strip-components=1 -xvzf /root/%s -C /usr/local/bin", containerdFileName)
	sealos.Cmd(host, cmd)
}
func (d *Containerd) Config(host string) {
	cmd := "mkdir -p " + d.lib()
	sealos.Cmd(host, cmd)
	cmd = "mkdir -p /etc/containerd"
	sealos.Cmd(host, cmd)
	//cmd = "containerd config default > /etc/containerd/config.toml"
	cmd = "echo \"" + string(d.configFile()) + "\" > /etc/containerd/config.toml"
	sealos.Cmd(host, cmd)
}

func (d *Containerd) Enable(host string) {
	cmd := "echo \"" + string(d.serviceFile()) + "\" > /usr/lib/systemd/system/containerd.service"
	sealos.Cmd(host, cmd)
	cmd = "systemctl enable  containerd.service && systemctl restart  containerd.service"
	sealos.Cmd(host, cmd)
}

func (d *Containerd) Version(host string) {
	cmd := "containerd --version"
	sealos.Cmd(host, cmd)
	logger.Warn("pull docker hub command. ex: ctr images pull docker.io/library/alpine:3.8")
	logger.Warn("pull http registry command. ex:  ctr images pull 10.0.45.222/library/alpine:3.8 --plain-http")
}

func (d *Containerd) Uninstall(host string) {
	cmd := "systemctl stop  containerd.service && systemctl disable containerd.service"
	sealos.Cmd(host, cmd)
	cmd = "rm -rf /usr/local/bin/ctr && rm -rf /usr/local/bin/containerd* "
	sealos.Cmd(host, cmd)
	cmd = "rm -rf /var/lib/containerd && rm -rf /etc/containerd/* "
	sealos.Cmd(host, cmd)
	if d.lib() != "" {
		cmd = "rm -rf " + d.lib()
		sealos.Cmd(host, cmd)
	}
}

func (d *Containerd) serviceFile() []byte {
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
func (d *Containerd) configFile() []byte {
	var templateText = string(`root = \"{{.CONTAINERD_LIB}}\"
state = \"/run/containerd\"
oom_score = 0

[grpc]
  address = \"/run/containerd/containerd.sock\"
  uid = 0
  gid = 0
  max_recv_message_size = 16777216
  max_send_message_size = 16777216

[debug]
  address = \"\"
  uid = 0
  gid = 0
  level = \"\"

[metrics]
  address = \"\"
  grpc_histogram = false

[cgroup]
  path = \"\"

[plugins]
  [plugins.cgroups]
    no_prometheus = false
  [plugins.cri]
    stream_server_address = \"127.0.0.1\"
    stream_server_port = \"0\"
    enable_selinux = false
    sandbox_image = \"k8s.gcr.io/pause:3.1\"
    stats_collect_period = 10
    systemd_cgroup = false
    enable_tls_streaming = false
    max_container_log_line_size = 16384
    [plugins.cri.containerd]
      snapshotter = \"overlayfs\"
      no_pivot = false
      [plugins.cri.containerd.default_runtime]
        runtime_type = \"io.containerd.runtime.v1.linux\"
        runtime_engine = \"\"
        runtime_root = \"\"
      [plugins.cri.containerd.untrusted_workload_runtime]
        runtime_type = \"\"
        runtime_engine = \"\"
        runtime_root = \"\"
    [plugins.cri.cni]
      bin_dir = \"/opt/cni/bin\"
      conf_dir = \"/etc/cni/net.d\"
      conf_template = \"\"
    [plugins.cri.registry]
      [plugins.cri.registry.mirrors]
        [plugins.cri.registry.mirrors.\"docker.io\"]
          endpoint = [\"https://registry-1.docker.io\"]
        {{range .CONTAINERD_REGISTRY -}}[plugins.cri.registry.mirrors.\"{{.}}\"]
          endpoint = [\"{{.}}\"]
    {{end -}}
    [plugins.cri.x509_key_pair_streaming]
      tls_cert_file = \"\"
      tls_key_file = \"\"
  [plugins.diff-service]
    default = [\"walking\"]
  [plugins.linux]
    shim = \"containerd-shim\"
    runtime = \"runc\"
    runtime_root = \"\"
    no_shim = false
    shim_debug = false
  [plugins.opt]
    path = \"/opt/containerd\"
  [plugins.restart]
    interval = \"10s\"
  [plugins.scheduler]
    pause_threshold = 0.02
    deletion_threshold = 0
    mutation_threshold = 100
    schedule_delay = \"0s\"
    startup_delay = \"100ms\"
`)
	tmpl, err := template.New("text").Parse(templateText)
	if err != nil {
		logger.Error("template parse failed:", err)
		panic(1)
	}
	var envMap = make(map[string]interface{})
	envMap["CONTAINERD_REGISTRY"] = RegistryArr
	envMap["CONTAINERD_LIB"] = d.lib()
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	return buffer.Bytes()
}

func (d *Containerd) Print() {
	data, err := Asset("install/command/containerd.json")
	if err != nil {
		logger.Error(err)
	}
	var versions []string
	_ = json.Unmarshal(data, &versions)
	logger.Debug(string(data))
	for _, v := range versions {
		println(v)
	}
}

func (d *Containerd) Fetch() {
	url := "https://containerd.io/downloads/"
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		logger.Error("http code is not 200")
		os.Exit(1)
	}
	var versions []string
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	ahtml := doc.Find("a[class=\"button is-dark is-clipboard\"]")
	for _, html := range ahtml.Nodes {
		attr := html.Attr
		if len(attr) == 3 {
			if attr[2].Key == "data-clipboard-text" {
				versions = append(versions, attr[2].Val)
				logger.Debug("加入缓存值：%s", attr[2].Val)
			}
		}
	}
	logger.Debug("写入json文件")
	dockerJson, err := os.OpenFile("install/command/containerd.json", os.O_WRONLY, 0755)
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
