package util

import (
	"os"
	"os/user"
	"path"
)

var base_dir_name = ".svalbard"
var keys_dir_name = "keys"

func BaseDir() string {
	usr, err := user.Current()
	if err != nil {
		Error(err)
	}
	dir := path.Join(usr.HomeDir, base_dir_name)
	keys_dir := path.Join(dir, keys_dir_name)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	if _, err := os.Stat(keys_dir); os.IsNotExist(err) {
		os.Mkdir(keys_dir, os.ModePerm)
	}

	return dir
}

func KeysDir() string {
	base := BaseDir()
	return path.Join(base,keys_dir_name)
}
