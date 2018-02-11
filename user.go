package pullhashi

import (
	"os"
	"os/user"
)

// UserBinDir returns back an os intelligent home directory.
func UserBinDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + "/bin"
}

// EnsureDirExists make sure we can write our bins out before we download them.
func EnsureDirExists(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
