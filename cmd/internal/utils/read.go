package utils

import (
	"backuper-tool/cmd/internal/types"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFileAsJson(cfgStruct *types.ConfigStructure) error {
	cfgBytes, err := os.ReadFile(types.ConfigPath)
	if err != nil {
		fmt.Println("Can't read the file", err)
		return err
	}
	json.Unmarshal(cfgBytes, cfgStruct)

	return nil
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
