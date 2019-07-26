package install

import "sync"

//username
var (
	User        string
	Passwd      string
	Hosts       []string
	RegistryArr []string
	DockerLib   string
	PkgUrl      string
)

const fileName = "docker.tgz"

func NewDockerInstaller() *DockerInstaller {
	return &DockerInstaller{}
}

type DockerInstaller struct {
}

type DockerInstallInterface interface {
	Install()
	UnInstall()
}

func (s *DockerInstaller) Install() {
	var wg sync.WaitGroup
	for _, host := range Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			sendPackage(PkgUrl)
			tarDocker(host)
			configDocker(host)
			enableDocker(host)
			versionDocker(host)
		}(host)
	}
	wg.Wait()
}

func (s *DockerInstaller) UnInstall() {
	var wg sync.WaitGroup
	for _, host := range Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			uninstallDocker(host)
		}(host)
	}
	wg.Wait()
}
