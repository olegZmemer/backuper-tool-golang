package handlers

import (
	types "backuper-tool/cmd/internal/types"
	"encoding/json"
	"fmt"
	"os"
)

func (b *Backuper) Add(name string, source string, output string) error {
	// read file
	cfgBytes, err := os.ReadFile(types.ConfigPath)
	if err != nil {
		fmt.Println("Can't read the file", err)
		return err
	}

	// append new instance
	configStrct := &types.ConfigStructure{}
	json.Unmarshal(cfgBytes, configStrct)

	instances := *configStrct.Instances
	instances = append(instances, types.ConfigElement{
		Name:   name,
		Source: source,
		Output: output,
	})

	configStrct.Instances = &instances
	fmt.Println(configStrct.Instances)

	// rewrite file
	cfgFile, _ := os.OpenFile(types.ConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer cfgFile.Close()

	cfgBytes, _ = json.Marshal(configStrct)
	cfgFile.Write(cfgBytes)

	return nil
}
