package install

import (
	"github.com/cuisongliu/container-install/install/command"
	"sync"
)

var Docker bool
var Hosts []string

const (
	Install = 1 + iota
	Uninstall
	Print
)

func NewInstaller() mainInterface {
	var mainInterface mainInterface
	mainInterface = &Installer{}
	if IsDocker() {
	} else {

	}
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
			factory(Install, host)
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
			factory(Uninstall, host)
		}(host)
	}
	wg.Wait()
}

func (s Installer) Print() {
	factory(Print, "")
}

func IsDocker() bool {
	return Docker
}

func factory(itype int, host string) {
	switch itype {
	case Install:
		if IsDocker() {
			docker.SendPackage(host)
			docker.Tar(host)
			docker.Config(host)
			docker.Enable(host)
			docker.Version(host)
		} else {
			containerd.SendPackage(host)
			containerd.Tar(host)
			containerd.Config(host)
			containerd.Enable(host)
			containerd.Version(host)
		}
	case Uninstall:
		if IsDocker() {
			docker.Uninstall(host)
		} else {
			containerd.Uninstall(host)
		}
	case Print:
		if IsDocker() {
			docker.Print()
		} else {
			containerd.Print()
		}
	}
}
