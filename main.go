package main

import (
	"flag"
	"fmt"
	"version-next/pkg/entity"
	"version-next/pkg/handler"
)

func main() {
	// versionフラグを定義
	versionInput := flag.String("v", "", "An array of semantic versioning strings (e.g., 2.3.1, 2.4.0-rc.1, 2.3.2-beta.1)")
	bumpUpReleaseType := flag.String("r", "", "Release type to bump up (rc, beta, alpha, release)")

	flag.Parse()

	if *versionInput == "" {
		fmt.Println("Please provide version strings with the -v flag.")
		return
	}

	if *bumpUpReleaseType == "release" {
		*bumpUpReleaseType = "None"
	}

	versionHandler := handler.NewVersionHandler()

	vh := handler.NewVersionHandler()

	version, err := vh.ParseVersion(*versionInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if *bumpUpReleaseType != "" {
		parsedType := entity.ParsePreReleaseType(*bumpUpReleaseType)
		if version.PreRelease.Type == parsedType {
			version = versionHandler.BumpUpVersion(version)
		} else {
			version = versionHandler.BumpUpPreReleaseType(version, parsedType)
		}
	} else {
		version = versionHandler.BumpUpVersion(version)
	}

	fmt.Println(version.String())

}
