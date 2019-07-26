package command

var (
	User   string
	Passwd string

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
	//
	serviceFile() []byte
	configFile() []byte
}
