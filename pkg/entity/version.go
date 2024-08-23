package entity

import "fmt"

type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease PreRelease
}

func NewVersion(major, minor, patch int, preRelease PreRelease) *Version {
	return &Version{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: preRelease,
	}
}

func (v *Version) String() string {
	if v.PreRelease.Type != None {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.PreRelease.String())
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
