package gracemain

import (
	"os"
	"path/filepath"
	"time"
)

func WriteLastRun(dir string) error {
	return os.WriteFile(
		filepath.Join(dir, "last-run.txt"),
		[]byte(time.Now().Format(time.RFC3339)),
		0640,
	)
}
