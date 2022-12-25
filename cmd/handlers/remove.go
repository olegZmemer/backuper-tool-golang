package handlers

import (
	types "backuper-tool/cmd/internal/types"
	utils "backuper-tool/cmd/internal/utils"
	"encoding/json"
	"os"
)

func (b *Backuper) Remove(name string) error {
	// read file
	configStrct := &types.ConfigStructure{}
	utils.ReadFileAsJson(configStrct)

	instances := *configStrct.Instances

	for i, v := range instances {
		if v.Name == name {
			instances = append(instances[:i], instances[i+1:]...)
		}
	}

	configStrct.Instances = &instances
	// rewrite file
	cfgFile, _ := os.OpenFile(types.ConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer cfgFile.Close()

	cfgBytes, _ := json.Marshal(configStrct)
	cfgFile.Write(cfgBytes)

	return nil
}
