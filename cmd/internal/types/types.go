package types

import "time"

const ConfigPath = "/etc/backuper/config.json"

const (
	InitCommand = "init"
	AddCommand  = "add"
	ListCommand = "list"
)

type ConfigStructure struct {
	Instances *[]ConfigElement `json:"instances"`
}

type ConfigElement struct {
	Name     string         `json:"name"`
	Source   string         `json:"source"`
	Output   string         `json:"output"`
	Interval *time.Duration `json:"intervals"`
	ToDelete *bool          `json:"toDelete"`
}
