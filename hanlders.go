package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (b *Backuper) InitConfig() error {
	if _, err := os.Stat("/etc/backuper"); err == nil {
		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("Config already exists, exiting program...")

			return nil
		}
	} else {
		fmt.Println("Initializing config in '/etc/backuper/config.json'...")

		err = os.Mkdir("/etc/backuper", os.FileMode(4096))
		if err != nil {
			fmt.Println("Can't create folder", err)
		}

		f, err := os.Create(configPath)
		if err != nil {
			fmt.Println("err with creating", err)
		}
		defer f.Close()

		_, err = f.WriteString("{\"instances\": []}")
		if err != nil {
			fmt.Println("Can't create file", err)

			return err
		}

		fmt.Println("Config file created!")
	}

	return nil
}

func (b *Backuper) Add(name string, source string, output string) error {
	// read file
	cfgBytes, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Can't read the file", err)
		return err
	}

	// append new instance
	configStrct := &ConfigStructure{}
	json.Unmarshal(cfgBytes, configStrct)

	instances := *configStrct.Instances
	instances = append(instances, ConfigElement{
		Name:   name,
		Source: source,
		Output: output,
	})

	configStrct.Instances = &instances
	fmt.Println(configStrct.Instances)

	// rewrite file
	cfgFile, _ := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer cfgFile.Close()

	cfgBytes, _ = json.Marshal(configStrct)
	cfgFile.Write(cfgBytes)

	return nil
}
