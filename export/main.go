package main

import (
	"github.com/orestonce/korm"
)

func main() {
	korm.MustCreateCode(korm.MustCreateCode_Req{
		ModelPkgDir:      "./model",
		ModelPkgFullPath: "way_m3u8/model",
		ModelNameList: []string{
			"Work_D",
		},
		OutputFileName: `./model/generated_korm.go`,
		GenMustFn:      true,
	})
}
