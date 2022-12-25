package handlers

import (
	types "backuper-tool/cmd/internal/types"
	"fmt"
	"os"
)

func (b *Backuper) InitConfig() error {
	if _, err := os.Stat("/etc/backuper"); err == nil {
		if _, err := os.Stat(types.ConfigPath); err == nil {
			fmt.Println("Config already exists, exiting program...")

			return nil
		}
	} else {
		fmt.Println("Initializing config in '/etc/backuper/config.json'...")

		err = os.Mkdir("/etc/backuper", os.FileMode(4096))
		if err != nil {
			fmt.Println("Can't create folder", err)
		}

		f, err := os.Create(types.ConfigPath)
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
