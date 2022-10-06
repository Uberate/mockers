package files

import (
	"os"
)

// IsFileExists return true when file is exists
func IsFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsPermission(err) {
			return true
		}
		return false
	}
	return true
}

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
