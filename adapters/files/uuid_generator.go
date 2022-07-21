package files

import (
	"os"
	"os/exec"
)

const envAuthKeysPathname = "AUTH_KEYS_PATHNAME"

func UUIDGenerator() error {
	authKeysPathname := os.Getenv(envAuthKeysPathname)
	newUUID, err := exec.Command("uuidgen").Output()
	WriteFile(authKeysPathname, string(newUUID))
	return err
}
