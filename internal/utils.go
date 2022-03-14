// Copyright (c) 2022 qianjunakasumi <i@qianjunakasumi.ren>
//                    qianjunakasumi <qianjunakasumi@outlook.com>
//                    [qianjunakasumi] 千橘雫霞(https://github.com/qianjunakasumi)
//
//      This Source Code Form is subject to the terms of the Mozilla Public
//      License, v. 2.0. If a copy of the MPL was not distributed with this
//      file, You can obtain one at http://mozilla.org/MPL/2.0/.

package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"

	"github.com/qianjunakasumi/go-cqhttp-docs-generator/internal/steel"

	"github.com/spf13/viper"
)

func ParseGolangFile() (file *ast.File) {
	wd, _ := os.Getwd()
	file, err := parser.ParseFile(
		token.NewFileSet(),
		filepath.Join(wd, ".cache", "api.go"),
		nil,
		parser.ParseComments,
	)
	if err != nil {
		panic("解析源文件失败：" + err.Error())
	}

	return
}

func ParseAPIFile(b io.Reader) (a *steel.DescriptionDocument) {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(b)
	if err != nil {
		panic("读取 API 定义文档错误：" + err.Error())
	}

	a = new(steel.DescriptionDocument)
	err = viper.Unmarshal(a)
	if err != nil {
		panic("解析 API 定义文档错误：" + err.Error())
	}

	return
}
