package main

import (
	"VersionNext/pkg/handler"
	"flag"
	"fmt"
)

func main() {
	// versionフラグを定義
	versionInput := flag.String("v", "", "An array of semantic versioning strings (e.g., 2.3.1, 2.4.0-rc.1, 2.3.2-beta.1)")

	flag.Parse()

	if *versionInput == "" {
		fmt.Println("Please provide version strings with the -v flag.")
		return
	}

	versionHandler := handler.NewVersionHandler()

	vh := handler.NewVersionHandler()

	version, err := vh.ParseVersion(*versionInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	version = versionHandler.BumpUpVersion(version)

	fmt.Println("New version:", version.String())

}
