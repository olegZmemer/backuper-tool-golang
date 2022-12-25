package handlers

import (
	types "backuper-tool/cmd/internal/types"
	utils "backuper-tool/cmd/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func (b *Backuper) Add(args []string) error {
	// check correctness of arguments
	if len(args) < 3 {
		fmt.Println("not enough args, use 'backuper add <name> <sourceDir> <outputDir>'")
		return errors.New("wrong arguments")
	}

	if _, err := os.Stat(args[2]); err != nil {
		fmt.Println("This source folder doesn't exist. Specify existing folder.")
		return errors.New("wrong arguments")
	}
	if _, err := os.Stat(args[3]); err != nil {
		fmt.Println("This output folder doesn't exist. Specify existing folder.")
		return errors.New("wrong arguments")
	}

	// read file
	configStrct := &types.ConfigStructure{}
	utils.ReadFileAsJson(configStrct)

	for _, inst := range *configStrct.Instances {
		if args[1] == inst.Name {
			fmt.Println("Instance with this name already exists, try another name.")
			return nil
		}
	}

	instances := *configStrct.Instances
	instances = append(instances, types.ConfigElement{
		Name:   args[1],
		Source: args[2],
		Output: args[3],
	})

	configStrct.Instances = &instances

	// rewrite file
	cfgFile, _ := os.OpenFile(types.ConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer cfgFile.Close()

	cfgBytes, _ := json.Marshal(configStrct)
	if _, err := cfgFile.Write(cfgBytes); err != nil {
		fmt.Println("error while writing data")
		return err
	}

	fmt.Printf("Instance %s was created", args[0])
	return nil
}
