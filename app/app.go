package app

import "io/fs"

const (
	EnvConfigDir = "ADR_CONFIG_DIR"

	DefaultDirPerms  fs.FileMode = 0664
	DefaultFilePerms fs.FileMode = 0644
)

var Version = "v0.0.1"
