package core

import (
	"github.com/mesogii/svld/util"

	"io/ioutil"
	"strings"
	"path"
	"crypto/rand"
)



func Key_Gen_Intf(kn string) error {

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

func generateKey() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		util.Error(KEY_GEN_ERROR)
	}

	return key
}

func keyExists(kn string) bool {
	fn := ""
	keys_dir := util.KeysDir()
	new_key_name := keyFileName(kn)

	keys, _ := ioutil.ReadDir(keys_dir)
	for _, f := range keys {
		fn = f.Name()
		if string(fn) == new_key_name {
			return true
		}
	}

	return false
}

func keyFileName(fp string) string {
	sp := strings.Split(path.Base(fp), ".")

	new_name := sp[0] + ".key"
	return new_name
}
