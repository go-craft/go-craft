package version

import (
	"fmt"
)

var (
	VersionMajor = 0
	VersionMinor = 0
	VersionPatch = 1
	VersionName = "go-craft"
)

type GoCraftVersion struct {
	major int
	minor int
	patch int
	name string
}

func NewGoCraftVersion() *GoCraftVersion {
	return &GoCraftVersion{
		major: VersionMajor,
		minor: VersionMinor,
		patch: VersionPatch,
		name:  VersionName,
	}
}

func (v *GoCraftVersion) Name() string {
	return v.name
}

func (v *GoCraftVersion) SemVer() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}