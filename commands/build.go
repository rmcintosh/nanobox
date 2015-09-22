// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package commands

//
import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pagodabox/nanobox-cli/config"
	"github.com/pagodabox/nanobox-cli/util"
	"github.com/pagodabox/nanobox-golang-stylish"
)

//
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Rebuilds/compiles your project",
	Long: `
Description:
  Rebuilds/compiles your project`,

	Run: nanoBuild,
}

// nanoBuild
func nanoBuild(ccmd *cobra.Command, args []string) {
	fmt.Printf(stylish.Bullet("Building codebase..."))

	//
	build := util.Sync{
		Model:   "build",
		Path:    fmt.Sprintf("http://%s/builds", config.ServerURI),
		Verbose: fVerbose,
	}

	//
	build.Run(args)

	//
	switch build.Status {

	// sync completed successfully
	case "complete":
		fmt.Printf(stylish.Bullet("Build complete... Navigate to %v.nano.dev to view your app.", config.App))

		// failed sync's wont be handled here, as the server or hooks should already
		// be providing output
	case "errored":
		// fmt.Printf(stylish.Error("Build failed", "Your build failed to well... build"))

		// if a build is run w/o having first run a deploy
	case "unavailable":
		fmt.Printf(stylish.ErrBullet("Before you can run a build, you must first deploy."))
	}
}
