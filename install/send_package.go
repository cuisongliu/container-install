package install

import (
	"fmt"
	"path"
	"strings"
	"sync"
)

//SendPackage is
func (s *DockerInstaller) SendPackage(url string) {
	pkg := path.Base(url)
	//only http
	isHttp := strings.HasPrefix(url, "http")
	wgetCommand := ""
	if isHttp {
		wgetParam := ""
		if strings.HasPrefix(url, "https") {
			wgetParam = "--no-check-certificate"
		}
		wgetCommand = fmt.Sprintf(" wget %s ", wgetParam)
	}
	remoteCmd := fmt.Sprintf("cd /root &&  %s %s ", wgetCommand, url)
	kubeLocal := fmt.Sprintf("/root/%s", pkg)
	var wm sync.WaitGroup
	for _, master := range Hosts {
		wm.Add(1)
		go func(master string) {
			defer wm.Done()
			if isHttp {
				go WatchFileSize(master, kubeLocal, GetFileSize(url))
				Cmd(master, remoteCmd)
			} else {
				Copy(master, url, kubeLocal)
			}
		}(master)
	}
	wm.Wait()
}
