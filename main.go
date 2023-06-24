package main

import (
	"fmt"
	"os"
	"yaml_v3/types"

	"gopkg.in/yaml.v3"
)

func main() {

	oamYaml, err := os.ReadFile("oam.yaml")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	oamStruct := &types.OAM{}
	yaml.Unmarshal(oamYaml, oamStruct)

	backstageStruct := types.NewBackstage()
	backstageStruct.Metadata.Name = oamStruct.Metadata.Name
	backstageStruct.Spec.Type = "service"

	backstageYaml, err := yaml.Marshal(&backstageStruct)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	err = os.WriteFile("result.yaml", backstageYaml, 0777)
	if err == nil {
		fmt.Print("File Created Successfully")
	} else {
		fmt.Print(err.Error())
	}
}
