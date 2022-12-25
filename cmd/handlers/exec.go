package handlers

import (
	"backuper-tool/cmd/internal/types"
	"backuper-tool/cmd/internal/utils"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func (b *Backuper) Exec(args []string) error {
	var flags []string

	for i := 0; i < len(args); i++ {
		byteStr := []byte(args[i])

		if string(byteStr[0]) == "-" {
			flags = append(flags, strings.Split(string(byteStr[1:]), "")...)
		}
	}
	cfgStruct := types.ConfigStructure{}
	utils.ReadFileAsJson(&cfgStruct)

	choosenInstance := &types.ConfigElement{}
	for _, inst := range *cfgStruct.Instances {
		if inst.Name == args[1] {
			choosenInstance = &inst
		}
	}

	if len(choosenInstance.Name) == 0 {
		fmt.Println("There is no such name, use 'add' to add backup instances")
		return errors.New("no such name")
	}

	if utils.IndexOf("m", flags) != -1 {
		if _, err := os.Stat(choosenInstance.Source); err != nil {
			fmt.Println("Source folder doesn't exist! Update your instance to currect state")
			return err
		}
		if _, err := os.Stat(choosenInstance.Output); err != nil {
			fmt.Println("Output folder doesn't exist! Update your instance to currect state")
			return err
		}
		cmdRsync := exec.Command("rsync", "-a", choosenInstance.Source, choosenInstance.Output)

		if err := cmdRsync.Run(); err != nil {
			fmt.Println("Error while merge two folder, exiting proccess")
			return err
		}
		fmt.Println("Synced successfully!")
	} else if utils.IndexOf("r", flags) != -1 {
		// remove old folder
		fmt.Println("Delete old backup folder from output path..")

		srcPathArr := strings.Split(choosenInstance.Source, "/")
		lastPathName := srcPathArr[len(srcPathArr)-1]
		outputString := fmt.Sprintf("%s/%s", choosenInstance.Output, lastPathName)

		if _, err := os.Stat(outputString); err != nil {
			err := os.RemoveAll(outputString)
			if err != nil {
				fmt.Println("Error while deleting old backup folder.")
				return err
			}
		}

		// copy source folder
		fmt.Println("Copying files, wait until it ends ... ")
		cpCmd := exec.Command("cp", "-r", choosenInstance.Source, choosenInstance.Output)

		if err := cpCmd.Run(); err != nil {
			fmt.Printf("ERROR WHILE COPIYNG. DELETING COPIED FILES (in output)...")
			os.RemoveAll(choosenInstance.Output)

			return err
		}

		fmt.Println("Files successfully backuped!")
	} else if len(flags) == 0 {
		fmt.Println("No flags given - backup to new folder with unix_time")

		bckpString := fmt.Sprintf(
			"%s/%s_%d",
			choosenInstance.Output,
			choosenInstance.Name,
			time.Now().Unix(),
		)

		copyCmd := exec.Command("cp", "--recursive", choosenInstance.Source, bckpString)
		if err := copyCmd.Run(); err != nil {
			fmt.Println("ERROR WHILE COPIYNG. DELETING COPIED FILES (in output)...", err)
			os.RemoveAll(bckpString)

			return err
		}
	} else {
		fmt.Println("No such flags, use -r, -m or don't put flags to create time_unix copy")
	}

	return nil
}
