// Copyright (c) 2022 qianjunakasumi <i@qianjunakasumi.ren>
//                    qianjunakasumi <qianjunakasumi@outlook.com>
//                    [qianjunakasumi] 千橘雫霞(https://github.com/qianjunakasumi)
//
//      This Source Code Form is subject to the terms of the Mozilla Public
//      License, v. 2.0. If a copy of the MPL was not distributed with this
//      file, You can obtain one at http://mozilla.org/MPL/2.0/.

package steel

type DescriptionDocument struct {
	ApiVersion string              `yaml:"apiVersion"`
	Metadata   DescriptionMetadata `yaml:"metadata"`
	Spec       DescriptionSpec     `yaml:"spec"`
}

type DescriptionMetadata struct {
	AvailableVersion string        `yaml:"availableVersion"`
	Family           string        `yaml:"family"`
	Genus            string        `yaml:"genus"`
	Species          string        `yaml:"species"`
	SpeciesI18n      string        `yaml:"speciesI18n"`
	Annotations      []Annotations `yaml:"annotations"`
	CommitId         string        `yaml:"commitId"`
	UpdateTimestamp  string        `yaml:"updateTimestamp"`
}

type DescriptionSpec struct {
	Name            DescriptionName     `yaml:"name"`
	Endpoint        string              `yaml:"endpoint"`
	Description     string              `yaml:"description"`
	Payload         DescriptionPayload  `yaml:"payload"`
	PayloadExample  string              `yaml:"payloadExample"`
	Response        DescriptionResponse `yaml:"response"`
	ResponseExample string              `yaml:"responseExample"`
	Extra           string              `yaml:"extra"`
}

type DescriptionName struct {
	Metadata struct {
		AutoFill `yaml:"autoFill"`
	} `yaml:"metadata"`
	Spec string `yaml:"spec"`
}

type DescriptionPayload struct {
	Metadata struct {
		Annotations []Annotations `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec []DescriptionPayloadSpec `yaml:"spec"`
}

type DescriptionPayloadSpec struct {
	Name string `yaml:"name"`
	Type struct {
		Metadata struct {
			AutoFill `yaml:"autoFill"`
			Link     string `yaml:"link"`
		}
		Spec string `yaml:"spec"`
	} `yaml:"type"`
	Default     string `yaml:"default"`
	Optional    bool   `yaml:"optional"`
	Description string `yaml:"description"`
}

type DescriptionResponse struct {
	Metadata struct {
		Annotations []Annotations `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec []DescriptionResponseSpec `yaml:"spec"`
}

type DescriptionResponseSpec struct {
	Name string `yaml:"name"`
	Type struct {
		Metadata struct {
			AutoFill `yaml:"autoFill"`
			Link     string `yaml:"link"`
		}
		Spec string `yaml:"spec"`
	}
	Description string `yaml:"description"`
}
