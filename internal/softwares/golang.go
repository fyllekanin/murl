package softwares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/fyllekanin/murl/internal/persistance"
)

type GoLangVersion struct {
	Version string `json:"version"`
	Stable  bool   `json:"stable"`
	Files   []struct {
		Os   string `json:"os"`
		Arch string `json:"arch"`
		Kind string `json:"kind"`
	} `json:"files"`
}

type Golang struct{}

func (s *Golang) Name() string {
	return "go"
}

func (s *Golang) List() ([]SoftwareVersion, error) {
	resp, err := http.Get("https://go.dev/dl/?mode=json&include=all")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to fetch golang versions")
	}

	var result []GoLangVersion
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to parse golang versions")
	}

	versions := make([]SoftwareVersion, 0)
	for _, item := range result {
		if !s.isVersionApplicable(item) {
			continue
		}

		versions = append(versions, SoftwareVersion{
			Name:        item.Version,
			Stable:      item.Stable,
			IsInstalled: persistance.IsVersionInstalled(s.Name(), item.Version),
		})
	}
	return versions, nil
}

func (s *Golang) isVersionApplicable(version GoLangVersion) bool {
	for _, item := range version.Files {
		if item.Os == runtime.GOOS && item.Arch == runtime.GOARCH && item.Kind == "archive" {
			return true
		}
	}
	return false
}

func NewGoLang() Software {
	return &Golang{}
}
