package install

import "testing"

func TestDockerInstaller_SendPackage(t *testing.T) {
	User = "root"
	Passwd = "admin"
	Hosts = []string{"172.16.213.131"}
	sendPackage("https://download.docker.com/linux/static/stable/x86_64/docker-19.03.0.tgz")
}
func TestDockerInstaller_SendPackageLocalFile(t *testing.T) {
	User = "root"
	Passwd = "admin"
	Hosts = []string{"172.16.213.131"}
	sendPackage("/home/cuisongliu/Documents/kubernetes-doc/docker-19.03.0.tgz")
}
