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
		b.Add("third", "hello!!!!!!", "31213")
	case types.ListCommand:
		b.List()
	case types.RemoveCommand:
		b.Remove("new")
	default:
		fmt.Println("There is no such command.")
	}
}
