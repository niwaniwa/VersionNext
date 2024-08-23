package main

import (
	"VersionNext/pkg/entity"
	"VersionNext/pkg/handler"
	"flag"
	"fmt"
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
		version = versionHandler.BumpUpPreReleaseType(version, parsedType)
	} else {
		version = versionHandler.BumpUpVersion(version)
	}

	fmt.Println("New version:", version.String())

}
