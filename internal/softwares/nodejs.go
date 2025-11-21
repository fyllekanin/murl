package softwares

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	"github.com/fyllekanin/murl/internal/persistance"
)

type ConvertibleLts string

func (bit *ConvertibleLts) UnmarshalJSON(data []byte) error {
	if string(data) == "false" {
		*bit = ""
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*bit = ConvertibleLts(str)
	return nil
}

type NodeVersion struct {
	Version string         `json:"version"`
	Files   []string       `json:"files"`
	Lts     ConvertibleLts `json:"lts"`
}

type NodeJs struct{}

func (s *NodeJs) Name() string {
	return "nodejs"
}

func (s *NodeJs) Install(version string) error {
	// empty
	return nil
}

func (s *NodeJs) List() ([]SoftwareVersion, error) {
	resp, err := http.Get("https://nodejs.org/dist/index.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []NodeVersion
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	versions := make([]SoftwareVersion, 0)
	for _, item := range result {
		if !s.isVersionApplicable(item) {
			continue
		}

		versions = append(versions, SoftwareVersion{
			Name:        item.Version,
			Stable:      true,
			IsLts:       item.Lts != "",
			LtsName:     string(item.Lts),
			IsInstalled: persistance.IsVersionInstalled(s.Name(), item.Version),
		})
	}
	return versions, nil
}

func (s *NodeJs) isVersionApplicable(version NodeVersion) bool {
	platform := runtime.GOOS + "-" + s.getArch()
	for _, item := range version.Files {
		if strings.HasPrefix(item, platform) {
			return true
		}
	}
	return false
}

func (s *NodeJs) getArch() string {
	switch runtime.GOOS {
	case "linux":
		if runtime.GOARCH == "amd64" {
			return "x64"
		}
		return runtime.GOARCH
	default:
		return runtime.GOARCH
	}
}

func NewNodeJs() Software {
	return &NodeJs{}
}
