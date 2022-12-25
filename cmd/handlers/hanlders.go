package handlers

import (
	types "backuper-tool/cmd/internal/types"
	"fmt"
)

type Backuper struct{}

func (b *Backuper) CommandHandler(args []string) {
	switch args[0] {
	case types.InitCommand:
		b.InitConfig()
	case types.AddCommand:
		b.Add(args)
	case types.ListCommand:
		b.List()
	case types.RemoveCommand:
		b.Remove(args)
	case types.ExecCommnad:
		b.Exec(args)
	default:
		fmt.Println("There is no such command.")
	}
}
