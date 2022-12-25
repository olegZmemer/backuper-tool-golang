package main

import (
	handlers "backuper-tool/cmd/handlers"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		backuper := handlers.Backuper{}

		backuper.CommandHandler(args)
	} else {
		fmt.Println("use arguments:\ninit \nadd \nlist ")
		return
	}

}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
