package models

import "time"

type TomlConfig struct {
	Title string
	Owner OwnerInfo
	DB    Database `toml:"database"`
}

type OwnerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type Database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}
