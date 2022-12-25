package handlers

import (
	types "backuper-tool/cmd/internal/types"
	utils "backuper-tool/cmd/internal/utils"
	"fmt"
)

func (b *Backuper) List() error {
	// read file
	configStrct := &types.ConfigStructure{}
	utils.ReadFileAsJson(configStrct)

	// print config data
	var list []byte
	for _, inst := range *configStrct.Instances {
		list = append(list, fmt.Sprintf("Name: %s, Source: %s, Output: %s \n", inst.Name, inst.Source, inst.Output)...)
	}
	fmt.Println(string(list))

	return nil
}
