package core

import (
	"github.com/mesogii/svalbard/util"

	"io/ioutil"
	"path"
	"crypto/rand"
)



func KeyListIntf() error {

	util.Notice("List all key files")
	//Loops all files in key directory, prints to terminal using util
	func KeyListIntf() error {

		util.Notice("List all key files")
		files, _ := ioutil.ReadDir(util.KeysDir())
		for _, f := range files {
			util.PrintKeyFile(f.Name())

		return nil
	}

func KeyGenIntf(kn string) error {

	if kn == "" {
		return NO_KEY_NAME_PROVIDED
	}

	if keyExists(kn) {
		return KEY_ALREADY_EXISTS
	}
	key_data := generateKey()
	key_file_name := keyFileName(kn)
	keys_dir := util.KeysDir()

	new_key_file_path := path.Join(keys_dir, key_file_name)

	err := ioutil.WriteFile(new_key_file_path, key_data, 0644)

	if err != nil {
		return CREATE_ENCRYPTED_FILE_ERROR
	}

	return nil
}

func KeyRemoveIntf(kn string) error {

	util.Notice("Remove key file, " + kn)

	return nil
}


// Generate a new AES Key
func generateKey() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		util.Error(KEY_GEN_ERROR)
	}

	return key
}

// Check if key file already exists in Svalbard base dir
func keyExists(kn string) bool {
	fn := ""
	keys_dir := util.KeysDir()
	new_key_name := keyFileName(kn)


	// Read from keys directory
	keys, _ := ioutil.ReadDir(keys_dir)
	for _, f := range keys {
		fn = f.Name()
		if string(fn) == new_key_name {
			return true
		}
	}

	return false
}


// Generate name for newly created key file
func keyFileName(fp string) string {
	new_name := fp + ".hem"
	return new_name
}
