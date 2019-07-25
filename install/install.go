package install

//username
var (
	User        string
	Passwd      string
	Hosts       []string
	RegistryArr []string
	DockerLib   string
	PkgUrl      string
)

type DockerInstaller struct {
}
