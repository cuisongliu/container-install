package install

import (
	"github.com/cuisongliu/container-install/install/command"
	"sync"
)

var Docker bool
var Hosts []string

func NewInstaller() mainInterface {
	var mainInterface mainInterface
	mainInterface = &Installer{}
	return mainInterface
}

type Installer struct {
}

type mainInterface interface {
	Install()
	UnInstall()
	Print()
}

var docker = command.NewDocker()
var containerd = command.NewContainerd()

func (s Installer) Install() {
	var wg sync.WaitGroup
	for _, host := range Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			docker.SendPackage(host)
			docker.Tar(host)
			docker.Config(host)
			docker.Enable(host)
			docker.Version(host)
		}(host)
	}
	wg.Wait()
}

func (s Installer) UnInstall() {
	var wg sync.WaitGroup
	for _, host := range Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			docker.Uninstall(host)
		}(host)
	}
	wg.Wait()
}

func (s Installer) Print() {
	if Docker {
		docker.Print()
	} else {
		containerd.Print()
	}
}
