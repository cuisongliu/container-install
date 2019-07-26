package install

import "testing"

func Test_InstallDocker(t *testing.T) {
	User = "root"
	Passwd = "admin"
	host := "172.16.213.131"
	RegistryArr = []string{}
	Lib = "/var/lib/docker"
	tarDocker(host)
	configDocker(host)
	enableDocker(host)
	versionDocker(host)
}

func Test_UnInstallDocker(t *testing.T) {
	User = "root"
	Passwd = "admin"
	host := "172.16.213.131"
	Lib = "/var/lib/docker"
	uninstallDocker(host)
	versionDocker(host)
}

func Test_DockerConfig(t *testing.T) {
	RegistryArr = []string{"127.0.0.1", "127.0.0.2"}
	Lib = "/var/lib/docker"
	s := dockerConfig(RegistryArr, Lib)
	t.Log(string(s))
}
