package install

import (
	"fmt"
	"strings"
	"sync"
)

func sendPackage(url string) {
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
	var wm sync.WaitGroup
	for _, host := range Hosts {
		wm.Add(1)
		go func(host string) {
			defer wm.Done()
			if isHttp {
				go WatchFileSize(host, localFile, GetFileSize(url))
				Cmd(host, remoteCmd)
			} else {
				Copy(host, url, localFile)
			}
		}(host)
	}
	wm.Wait()
}
