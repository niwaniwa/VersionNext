package handler

import (
	"errors"
	"strconv"
	"strings"
	"version-next/pkg/entity"
)

type VersionHandler struct {
}

func NewVersionHandler() *VersionHandler {
	return &VersionHandler{}
}

func (vh *VersionHandler) ParseVersion(versionInput string) (entity.Version, error) {
	versions := strings.Split(versionInput, "-")

	if len(versions) == 2 {
		intVersion, err := vh.parseVersionInt(versions[0])
		if err != nil {
			return entity.Version{}, err
		}
		preReleaseRaw := strings.Split(versions[1], ".")
		preReleaseVer, err := strconv.Atoi(preReleaseRaw[1])
		if err != nil {
			return entity.Version{}, err
		}
		preRelease := entity.PreRelease{Type: entity.ParsePreReleaseType(preReleaseRaw[0]), Index: preReleaseVer}
		return entity.Version{Major: intVersion[0], Minor: intVersion[1], Patch: intVersion[2], PreRelease: preRelease}, nil
	}

	intVersion, err := vh.parseVersionInt(versionInput)
	if err != nil {
		return entity.Version{}, err
	}

	if len(intVersion) != 3 {
		return entity.Version{}, errors.New("invalid version format")
	}

	return entity.Version{Major: intVersion[0], Minor: intVersion[1], Patch: intVersion[2], PreRelease: entity.PreRelease{Type: entity.None, Index: 0}}, nil
}

func (vh *VersionHandler) parseVersionInt(versionInput string) ([]int, error) {
	versions := strings.Split(versionInput, ".")
	intVersion := make([]int, len(versions))

	for i, v := range versions {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intVersion[i] = num
	}

	return intVersion, nil
}

func (vh *VersionHandler) BumpUpVersion(version entity.Version) entity.Version {
	if version.PreRelease.Type != entity.None {
		version.PreRelease.Index++
		return version
	}

	version.Patch++
	return version
}

func (vh *VersionHandler) BumpUpPreReleaseType(version entity.Version, preReleaseType entity.PreReleaseType) entity.Version {
	if vh.ValidateNoVersionRollback(version, preReleaseType) {
		version = vh.bumpUpPreReleaseType(version, preReleaseType)
		version = vh.bumpUpPreReleaseIndex(version)
		return version
	}

	if version.PreRelease.Type == entity.None && preReleaseType != entity.None {
		version = vh.BumpUpVersion(version)
		version = vh.bumpUpPreReleaseType(version, preReleaseType)
		return version
	}

	return entity.Version{}
}

func (vh *VersionHandler) bumpUpPreReleaseIndex(version entity.Version) entity.Version {
	version.PreRelease.Index++
	return version
}

func (vh *VersionHandler) resetPreReleaseIndex(version entity.Version) entity.Version {
	version.PreRelease.Index = 0
	return version
}

func (vh *VersionHandler) bumpUpPreReleaseType(version entity.Version, releaseType entity.PreReleaseType) entity.Version {
	switch releaseType {
	case entity.Alpha:
		version.PreRelease.Type = entity.Alpha
	case entity.Beta:
		version.PreRelease.Type = entity.Beta
	case entity.Rc:
		version.PreRelease.Type = entity.Rc
	case entity.None:
		version.PreRelease.Type = entity.None
	}
	version = vh.resetPreReleaseIndex(version)
	return version
}

func (vh *VersionHandler) ValidateNoVersionRollback(version entity.Version, targetReleaseType entity.PreReleaseType) bool {
	allowedTransitions := map[entity.PreReleaseType][]entity.PreReleaseType{
		entity.Alpha: {entity.Alpha, entity.Beta, entity.Rc, entity.None},
		entity.Beta:  {entity.Beta, entity.Rc, entity.None},
		entity.Rc:    {entity.Rc, entity.None},
	}

	allowed, exists := allowedTransitions[version.PreRelease.Type]
	if !exists {
		return false
	}

	for _, t := range allowed {
		if t == targetReleaseType {
			return true
		}
	}

	return false
}
