// Copyright (c) 2022 qianjunakasumi <i@qianjunakasumi.ren>
//                    qianjunakasumi <qianjunakasumi@outlook.com>
//                    [qianjunakasumi] 千橘雫霞(https://github.com/qianjunakasumi)
//
//      This Source Code Form is subject to the terms of the Mozilla Public
//      License, v. 2.0. If a copy of the MPL was not distributed with this
//      file, You can obtain one at http://mozilla.org/MPL/2.0/.

package yml2md

import (
	"strings"

	"github.com/qianjunakasumi/go-cqhttp-docs-generator/internal/steel"
)

type DescriptionGenerator struct {
	source *steel.DescriptionDocument

	payloadText  string
	responseText string
}

func NewDescriptionParser(d *steel.DescriptionDocument) *DescriptionGenerator {
	return &DescriptionGenerator{source: d}
}

func (d *DescriptionGenerator) Build() {
	d.Payload()
	d.Response()
}

func (d *DescriptionGenerator) Print() string {

	d.Build()

	format := new(strings.Builder)
	format.Grow(len(d.payloadText) + len(d.responseText))

	format.WriteString(`# `)
	format.WriteString(d.source.Spec.Name.Spec)
	format.WriteString(`

`)
	format.WriteString(d.source.Spec.Description)
	format.WriteString(`

## 接口 / 终结点

`)
	format.WriteString(d.source.Spec.Endpoint)
	format.WriteString(d.payloadText)
	format.WriteString(d.responseText)

	return format.String()
}

func (d *DescriptionGenerator) Payload() {

	payloadSpec := d.source.Spec.Payload.Spec
	format := new(strings.Builder)
	format.Grow(cap(payloadSpec))

	format.WriteString(`

## 请求载荷 / 参数

| 字段名 | 类型 | 默认值 | 可选 | 描述 |
|-------|-----|-------|-----|------|
`,
	)
	for _, spec := range payloadSpec {

		format.WriteString(`|`)
		format.WriteString(spec.Name)

		format.WriteString(`|`)
		format.WriteString(spec.Type.Spec)

		format.WriteString(`|`)
		if spec.Default == "none" {
			format.WriteString(`无`)
		} else {
			format.WriteString(spec.Default)
		}

		format.WriteString(`|`)
		if !spec.Optional {
			format.WriteString(`×`)
		} else {
			format.WriteString(`√`)
		}

		format.WriteString(`|`)
		format.WriteString(spec.Description)
		format.WriteString(`|
`)
	}

	d.payloadText = format.String()
}

func (d *DescriptionGenerator) Response() {

	responseSpec := d.source.Spec.Response.Spec
	format := new(strings.Builder)
	format.Grow(cap(responseSpec))

	format.WriteString(`

## 响应数据

| 字段名 | 类型 | 描述 |
|-------|-----|------|
`,
	)
	for _, spec := range responseSpec {

		format.WriteString(`|`)
		format.WriteString(spec.Name)

		format.WriteString(`|`)
		format.WriteString(spec.Type.Spec)

		format.WriteString(`|`)
		format.WriteString(spec.Description)
		format.WriteString(`|
`)
	}

	d.responseText = format.String()
}
