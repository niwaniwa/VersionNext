package entity

import "strconv"

type PreReleaseType string

const (
	Alpha PreReleaseType = PreReleaseType("alpha")
	Beta  PreReleaseType = PreReleaseType("beta")
	Rc    PreReleaseType = PreReleaseType("rc")
	None  PreReleaseType = PreReleaseType("")
)

func ParsePreReleaseType(preRelease string) PreReleaseType {
	switch preRelease {
	case "alpha":
		return Alpha
	case "beta":
		return Beta
	case "rc":
		return Rc
	default:
		return None
	}
}

type PreRelease struct {
	Type  PreReleaseType
	Index int
}

func NewPreRelease(preRelease PreReleaseType, index int) *PreRelease {
	return &PreRelease{
		Type:  preRelease,
		Index: index,
	}
}

func (pr *PreRelease) String() string {
	if pr.Type != None {
		return string(pr.Type) + "." + strconv.Itoa(pr.Index)
	}
	return ""
}
