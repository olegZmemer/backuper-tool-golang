package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	// determine is there device to write
	cmd := exec.Command("df")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	if len(args) != 0 {
		commandHandler(args[0])
	} else {
		fmt.Println("use arguments:\ninit \nadd \nlist ")
		return
	}

	// disks := out.String()

	// disks = strings.ReplaceAll(disks, "(MISSING)", "")
	// disks = strings.ReplaceAll(disks, "!", "")

	// if !strings.Contains(disks, "/media/bratka/My Passport") {
	// 	fmt.Println("I can't find any disk")
	// }

	// // copy data to device
	// if !Exists(args[0]) {
	// 	fmt.Println("Dir does not exist")
	// 	return
	// }

	// copyCmd := exec.Command("cp", "--recursive", args[0], fmt.Sprintf("/media/bratka/My Passport/backup_%d", time.Now().Unix()))
	// copyCmd.Run()
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func commandHandler(firstArg string, args ...[]string) {
	backuper := Backuper{}

	switch firstArg {
	case InitCommand:
		backuper.InitConfig()
	case AddCommand:
		backuper.Add("first", "hello", "213")
	case ListCommand:
		fmt.Println("List")
	default:
		fmt.Println("There is no such command.")
	}
}
