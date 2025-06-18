package softwares

type SoftwareVersion struct {
	Name        string
	Stable      bool
	IsInstalled bool
	IsLts       bool
	LtsName     string
}

type Software interface {
	// name of the software to be used in command + persistance
	Name() string
	// list of versions in decending order
	List() ([]SoftwareVersion, error)
}
