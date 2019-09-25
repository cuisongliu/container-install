package command

import (
	"fmt"
	sealos "github.com/fanux/sealos/install"
	"strings"
)

var (
	RegistryArr []string
	Lib         string
	PkgUrl      string
)

type stepInterface interface {
	SendPackage(host string)
	Tar(host string)
	Config(host string)
	Enable(host string)
	Version(host string)
	Uninstall(host string)
	Print()
	Fetch()
	//
	lib() string
	serviceFile() []byte
	configFile() []byte
}

func sendPackage(host, url, fileName string) {
	//only http
	isHttp := strings.HasPrefix(url, "http")
	downloadCmd := ""
	if isHttp {
		downloadParam := ""
		if strings.HasPrefix(url, "https") {
			downloadParam = "--no-check-certificate"
		}
		downloadCmd = fmt.Sprintf(" wget %s -O %s", downloadParam, fileName)
	}
	remoteCmd := fmt.Sprintf("cd /root &&  %s %s ", downloadCmd, url)
	localFile := fmt.Sprintf("/root/%s", fileName)
	if isHttp {
		go sealos.WatchFileSize(host, localFile, sealos.GetFileSize(url))
		sealos.Cmd(host, remoteCmd)
	} else {
		sealos.Copy(host, url, localFile)
	}
}
