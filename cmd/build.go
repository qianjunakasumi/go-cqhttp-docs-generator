// Copyright (c) 2022 qianjunakasumi <i@qianjunakasumi.ren>
//                    qianjunakasumi <qianjunakasumi@outlook.com>
//                    [qianjunakasumi] 千橘雫霞(https://github.com/qianjunakasumi)
//
//      This Source Code Form is subject to the terms of the Mozilla Public
//      License, v. 2.0. If a copy of the MPL was not distributed with this
//      file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "构建文档",
	Long:  `构建文档`,
	Run:   buildCmdRun,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildCmdRun(cmd *cobra.Command, args []string) {

}
