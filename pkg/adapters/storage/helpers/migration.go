package helpers

import (
	"fmt"
	"time"
)

func GenerateMigrationID(name string) string {
	timestamp := time.Now().Format("20060102_150405")
	return fmt.Sprintf("%s_%s", timestamp, name)
}
