package queries

import (
	"embed"
	"fmt"
)

//go:embed *.sql
var SQLFiles embed.FS

// Load lê uma query pelo nome do arquivo dentro do embed.FS
func Load(filename string) (string, error) {
	data, err := SQLFiles.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("erro ao ler SQL %s: %w", filename, err)
	}
	return string(data), nil
}
