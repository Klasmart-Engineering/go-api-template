// Copyright 2022 james wilson. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var config string

var rootCmd = &cobra.Command{
	Use: "go-api-template",
	Short: `🐺 go api template

If you have any suggestions, bug reports, or annoyances please report
them to our issue tracker at <https://github.com/james wilson/go-api-template/issues>`,
}

// Execute runs cmd tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
