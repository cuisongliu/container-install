package install

import "testing"

func Test_InstallDocker(t *testing.T) {
	User = "root"
	Passwd = "admin"
	host := "172.16.213.131"
	RegistryArr = []string{}
	DockerLib = "/var/lib/docker"
	tarDocker(host)
	configDocker(host)
	enableDocker(host)
	versionDocker(host)
}

func Test_UnInstallDocker(t *testing.T) {
	User = "root"
	Passwd = "admin"
	host := "172.16.213.131"
	DockerLib = "/var/lib/docker"
	uninstallDocker(host)
	versionDocker(host)
}

func Test_RegistryJoin(t *testing.T) {
	hosts := []string{"172.16.213.131", "172.16.213.131"}
	t.Log(registryJoin(hosts))
}

func Test_DockerConfig(t *testing.T) {
	RegistryArr = []string{"127.0.0.1", "127.0.0.2"}
	DockerLib = "/var/lib/docker"
	s := dockerConfig(RegistryArr, DockerLib)
	t.Log(string(s))
}
