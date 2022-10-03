package files

import (
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	perm := os.FileMode(0644)             // 0644 or 0777?
	err := os.WriteFile(path, data, perm) // ignore_security_alert
	if err != nil {
		return err
	}
	return err
}
