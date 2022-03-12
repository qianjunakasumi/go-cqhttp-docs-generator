// Copyright (c) 2022 qianjunakasumi <i@qianjunakasumi.ren>
//                    qianjunakasumi <qianjunakasumi@outlook.com>
//                    [qianjunakasumi] 千橘雫霞(https://github.com/qianjunakasumi)
//
//      This Source Code Form is subject to the terms of the Mozilla Public
//      License, v. 2.0. If a copy of the MPL was not distributed with this
//      file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新 API 描述文件",
	Long:  `更新 API 描述文件`,
	Run:   updateCmdRun,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("mirror", "https://ghproxy.com/https://raw.githubusercontent.com/Mrs4s/go-cqhttp/dev/coolq/api.go", "源文件镜像")
}

func updateCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("开始更新 API... 正在下载中")
	url, _ := cmd.Flags().GetString("mirror")
	download(url)
	fmt.Println("更新源文件成功，正在分析...")
}

func download(url string) {
	//https://raw.githubusercontent.com/Mrs4s/go-cqhttp/dev/coolq/api.go
	//https://ghproxy.com/https://raw.githubusercontent.com/Mrs4s/go-cqhttp/dev/coolq/api.go
	rsp, err := http.Get(url)
	if err != nil {
		panic("下载源文件文件失败：" + err.Error())
	}
	defer rsp.Body.Close()

	_ = os.MkdirAll(".cache", 0755)
	wd, _ := os.Getwd()
	file, err := os.Create(filepath.Join(wd, ".cache", "api.go"))
	if err != nil {
		panic("创建源文件缓存文件失败：" + err.Error())
	}

	_, err = io.Copy(file, rsp.Body)
	if err != nil {
		panic("保存源文件失败：" + err.Error())
	}
}
