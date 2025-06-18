package commands

import (
	"fmt"
	"slices"

	"github.com/fyllekanin/murl/internal/softwares"
)

func List(software softwares.Software, includeUnstable bool) {
	result, err := software.List()

	if err != nil {
		fmt.Println("Failed to fetch versions for " + software.Name())
		return
	}

	fmt.Println("Available versions (Ascending):")
	slices.Reverse(result)
	for _, item := range result {
		if !item.Stable && !includeUnstable {
			continue
		}
		fmt.Printf("%-15s %s%s%s\n", item.Name, getStableValue(item), getInstalledValue(item), getLtsValue(item))
	}
}

func getStableValue(item softwares.SoftwareVersion) string {
	if item.Stable {
		return " (Stable)"
	}
	return " (Unstable)"
}

func getInstalledValue(item softwares.SoftwareVersion) string {
	if item.IsInstalled {
		return " (Installed)"
	}
	return ""
}

func getLtsValue(item softwares.SoftwareVersion) string {
	if item.IsLts {
		if item.LtsName != "" {
			return " (LTS - " + item.LtsName + ")"
		}
		return " (LTS)"
	}
	return ""
}
