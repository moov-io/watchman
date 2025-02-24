package watchman

import (
	"embed"
)

//go:embed configs/config.default.yml
var ConfigDefaults embed.FS

//go:embed pkg/search/models.go
var ModelsFilesystem embed.FS
