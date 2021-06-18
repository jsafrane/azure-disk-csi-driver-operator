package generated

import (
	"embed"
)

//go:embed assets
var f embed.FS

// ReadFile reads and returns the content of the named file.
func ReadFile(name string) ([]byte, error) {
	return f.ReadFile("	assets/" + name)
}
