package handlers

import (
	types "backuper-tool/cmd/internal/types"
	"fmt"
)

type Backuper struct{}

func (b *Backuper) CommandHandler(firstArg string, args ...[]string) {
	switch firstArg {
	case types.InitCommand:
		b.InitConfig()
	case types.AddCommand:
		b.Add("first", "hello", "213")
	case types.ListCommand:
		fmt.Println("List")
	default:
		fmt.Println("There is no such command.")
	}
}
