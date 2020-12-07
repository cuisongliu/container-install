package command

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/cuisongliu/container-install/pkg"
	"github.com/cuisongliu/container-install/pkg/filesize"
	"github.com/wonderivan/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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
		go pkg.SSHConfig.LoggerFileSize(host, localFile, int(filesize.Do(url)))
		pkg.SSHConfig.Cmd(host, remoteCmd)
	} else {
		pkg.SSHConfig.Copy(host, url, localFile)
	}
}

func getUrl(rawurl string) ([]byte, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("解析url为空")
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logger.Debug("client do http status:", resp.Status)
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, resp.Body)
	_ = resp.Body.Close()
	return ioutil.ReadAll(buf)
}
