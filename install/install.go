package install

import "sync"

//username
var (
	User        string
	Passwd      string
	Hosts       []string
	RegistryArr []string
	Lib         string
	PkgUrl      string
)

const dockerFileName = "docker.tgz"

func NewInstaller() *Installer {
	return &Installer{}
}

type Installer struct {
}

type mainInterface interface {
	Install()
	UnInstall()
}
type stepInterface interface {
	tar(host string)
	config(host string)
	enable(host string)
	version(host string)
	uninstall(host string)
	//
	serviceFile() []byte
	configFile() []byte
}

var docker = newDocker()

func (s *Installer) Install() {
	var wg sync.WaitGroup
	for _, host := range Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			sendPackage(PkgUrl)
			docker.tar(host)
			docker.config(host)
			docker.enable(host)
			docker.version(host)
		}(host)
	}
	wg.Wait()
}

func (s *Installer) UnInstall() {
	var wg sync.WaitGroup
	for _, host := range Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			docker.uninstall(host)
		}(host)
	}
	wg.Wait()
}
