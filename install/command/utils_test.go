package command

import (
	"github.com/wonderivan/logger"
	"golang.org/x/crypto/ssh"
	"net"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	session, err := connect("root", "127.0.0.1:2222")
	if err != nil {
		logger.Error("	Error create ssh session failed", err)
		panic(1)
	}
	defer session.Close()

	b, err := session.CombinedOutput("ls /root")
	logger.Debug("command result is:", string(b))
	if err != nil {
		logger.Error("	Error exec command failed", err)
		panic(1)
	}
}

func connect(user, host string) (*ssh.Session, error) {
	//sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	//if err != nil {
	//	return nil,err
	//}
	//extendedAgent := agent.NewClient(sock)
	//signers, err := extendedAgent.Signers()
	//if err != nil {
	//	return nil,err
	//}
	data := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAuKmOJtSMjbH1kXhdRE9qYFzOTb1rO5wKQ12uRAmZgGrzzDDX
p4x9nkGn/GrTIqP8ra+HcCMDuwZtoKC1JR/oIyl/dJ1FETCOd8LILgJw/+00rPAY
0RgTHbYQX/M3m++P2YdVyu9WBC1Ah8Jfmp/odjGT7EPPOaRqC40JtRxWpT4Jnl1M
mzDYJ0fYMHzwRZDlenqK/Ce4riKsCMO1apiXuDthTUY56EHZv+igXdOWuZUe+ppL
00bLsTpBuQZtpOcoCtc0LuaI63eOjErWIIkkJDOQzouILFXXZtXlCZQ7glHTM4DW
E6DxM/vBzCvuuuNyRbhGvUyVoHkfuSVcQCAv3wIDAQABAoIBAQCzlbm4C6cxOete
8JaLk0wZsMe1lMwPaZ4Vi6qpYkiVOe4lGy7vM9MKsEFlWqJAowheBUGLDZJYNVUy
DHh+RTxlzbq1NylvITC8SYKSNC+exRogQVNxLZ+RmnjsR0VCckUMQPBvbjjR0Qxb
uu1tG6xgHbEE6aFDJqE24I+bQZcGFNhJO0b4hFcZA/VSjooijGVoyUc1RrTyuMYY
YsbCx/q7e3uZbxbJ5I/NQQtv9/ubVIv03lkdnF9THayaYLOyCpfxhogXVavIh7po
ITK4VMxXqorzpNxdPz5IbsF0uDdylILefKg5n7uhBhYWODmE067yMgxrdNx20wkA
ahgps9xhAoGBAPR1oflbKR7AyVjgUIBZRxPysTmrByOtTjM96VZSqYYSb8SpsnNR
fTm2+eqOHPseBu/nVWZ9oeKtjmYm8mDohv8VeuwTHm2qcbTMX+0Y2f8dKoS9nT7V
WDSPygeDc7K9250y9QYeAMsOIJ9LhSow1t68GzQWOGOimJn5ney/LdIxAoGBAMFh
QWplD/CVXCZVrlH6iYfL36JSdQXVLKAmldnlTy1EpfOchnnMjrDnzcyV65rgFy+e
r2eoiI+jaKLwnt0+hXduTq/ZwdoitAk8ICCR4uZ2ar/MLtJYv2Ni9y4cmXVR2Ov/
7fKlXpOU/Ds7UbFIh24PEPAq98YfUED4NhvHIA8PAoGBAJRNIJwzj3iWoA+I1Y1n
m9UgMB+5/7THGF/BuWKi9zDc0m1OPXH0B7IRrP98g1xcVP0JLCfnI2Ruwap7CiN8
LRlmoJHC73y8IAr8yVz+7JD10quAlHpf1wjcCkYQmwneX/K3zSmO3hBRW70HhZuY
0WGCYEZzDHZ1V3phkkrjmBqhAoGBALtyMFyXVdoYjVhDWVPxjHprLn1DfFeJCVa7
0CmEUzlH/6yiHt/VXsMwDpavA8/+Q7tPECtke+rvtK+smfFPd0QLUo62f2eYl/cT
pvirMMvAIT2FCCWxDOOjvIGgC0hja+dnDxlTHtfjZJNtroQwD8apJ/wFSmNrWvl/
H8PRQswVAoGBAONA7xTYP/PNkCQbO13IVgnVMq/IqmmK0dSbYgQ9sTyKmJF67P64
ODTRE4U2v3SyLQcD5toPKDPntS/WKe4xTXMLs+BkBCEBxrQRcNAAfRXcL2xPfXBb
X5tdaSV5v1DveV+c6t34CVle2lkSALmRurxYIob3k7Ld2jQd8D2Z94JI
-----END RSA PRIVATE KEY-----
`)
	s, _ := ssh.ParsePrivateKey(data)
	auths := []ssh.AuthMethod{ssh.PublicKeys(s)}

	config := ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}

	clientConfig := &ssh.ClientConfig{
		User:    user,
		Auth:    auths,
		Timeout: time.Duration(5) * time.Minute,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := AddrReformat(host)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}
