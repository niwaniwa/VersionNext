package handler_test

import (
	"testing"
	"version-next/pkg/entity"
	"version-next/pkg/handler"
)

func TestParseVersion(t *testing.T) {
	testCases := map[string]struct {
		versionInput string
		expected     string
	}{
		"MajorMinorPatch": {
			versionInput: "2.3.1",
			expected:     "2.3.1",
		},
		"MajorMinorPatchPreRelease": {
			versionInput: "2.4.0-rc.1",
			expected:     "2.4.0-rc.1",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			vh := handler.NewVersionHandler()

			version, err := vh.ParseVersion(tc.versionInput)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if version.String() != tc.expected {
				t.Errorf("expected: %s, but got: %s", tc.expected, version.String())
			}
		})
	}
}

func TestParseVersion_InvalidInput(t *testing.T) {
	testCases := map[string]struct {
		versionInput string
		expectedErr  bool
	}{
		"InvalidFormat": {
			versionInput: "invalid_version",
			expectedErr:  true,
		},
		"MissingPatchVersion": {
			versionInput: "1.0",
			expectedErr:  true,
		},
		"InvalidPreRelease": {
			versionInput: "1.0.0-beta.x",
			expectedErr:  true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			vh := handler.NewVersionHandler()
			_, err := vh.ParseVersion(tc.versionInput)
			if (err != nil) != tc.expectedErr {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err != nil)
			}
		})
	}
}

func TestBumpUpVersion(t *testing.T) {
	testCases := map[string]struct {
		versionInput string
		expected     string
		releaseType  entity.PreReleaseType
	}{
		"MajorMinorPatch": {
			versionInput: "2.3.1",
			expected:     "2.3.2",
			releaseType:  entity.None,
		},
		"BumpUpPreReleaseIndex": {
			versionInput: "2.4.0-rc.1",
			expected:     "2.4.0-rc.2",
			releaseType:  entity.None,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			vh := handler.NewVersionHandler()

			version, err := vh.ParseVersion(tc.versionInput)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			version = vh.BumpUpVersion(version)

			if version.String() != tc.expected {
				t.Errorf("expected: %s, but got: %s", tc.expected, version.String())
			}
		})
	}
}

func TestBumpUpPreRelease(t *testing.T) {
	testCases := map[string]struct {
		versionInput string
		expected     string
		releaseType  entity.PreReleaseType
	}{
		"BumpUpPreReleaseBeta": {
			versionInput: "2.3.1-alpha.3",
			expected:     "2.3.1-beta.1",
			releaseType:  entity.Beta,
		},
		"BumpUpPreReleaseRc": {
			versionInput: "2.4.0-beta.1",
			expected:     "2.4.0-rc.1",
			releaseType:  entity.Rc,
		},
		"BumpUpPreReleaseRc-2": {
			versionInput: "2.5.0-beta.3",
			expected:     "2.5.0-rc.1",
			releaseType:  entity.Rc,
		},
		"BumpUpRelease": {
			versionInput: "2.6.0-rc.2",
			expected:     "2.6.0",
			releaseType:  entity.None,
		},
		"BumpUpNextPreRelease": {
			versionInput: "2.7.0",
			expected:     "2.7.1-beta.0",
			releaseType:  entity.Beta,
		},
		"BumpUpNextPreRelease-2": {
			versionInput: "2.8.0",
			expected:     "2.8.1-alpha.0",
			releaseType:  entity.Alpha,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			vh := handler.NewVersionHandler()
			version, err := vh.ParseVersion(tc.versionInput)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			version = vh.BumpUpPreReleaseType(version, tc.releaseType)

			if version.String() != tc.expected {
				t.Errorf("expected: %s, but got: %s", tc.expected, version.String())
			}
		})
	}
}
