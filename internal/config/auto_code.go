package config

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transferRestart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-equipment" json:"serverApi" yaml:"server-equipment"`
	SInitialize     string `mapstructure:"server-initialize" json:"serverInitialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"serverModel" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"serverRequest"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"serverRouter" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"serverService" yaml:"server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-equipment" json:"webApi" yaml:"web-equipment"`
	WForm           string `mapstructure:"web-form" json:"webForm" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"webTable" yaml:"web-table"`
	WFlow           string `mapstructure:"web-flow" json:"webFlow" yaml:"web-flow"`
}
