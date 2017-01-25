package core

import (
	"github.com/mesogii/svalbard/util"
	"io/ioutil"
	"path"
	"strings"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)



func EncryptIntf(fp string) error {

	efn := encryptedFileName(fp)

	randomSymmetricKey := []byte("my secret 256 bit symmetric key ")


	if fp == "" {
		return NO_FILE_ERROR
	}

	dat, err := ioutil.ReadFile(fp)

	if err != nil {
		return READ_FILE_ERROR
	}

	encryptedFile, _ := encryptCBC(randomSymmetricKey, dat)

	err = ioutil.WriteFile(efn, encryptedFile, 0644)

	if err != nil {
		return CREATE_ENCRYPTED_FILE_ERROR
	} else {
		util.Notice("The encryption was successful.\nWould you like to delete the original file? (y/n)")
		answer := util.RecurCompareInput("y","n")
		if answer == "y" {
			var err = os.Remove(fp)
			if err != nil {
				return DELETE_UNENCRYPTED_FILE_ERROR
			}
		}
	}

	// decrypt and print the original file
	// _, _ := core.DecryptCBC(randomSymmetricKey, encryptedFile)


	return nil
}

func encryptCBC(symmetricKey []byte, inBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(symmetricKey)
	if err != nil {
		return nil, err
	}

	padLen := aes.BlockSize - len(inBytes)%aes.BlockSize
	padding := make([]byte, padLen)
	_, err = rand.Reader.Read(padding)
	if err != nil {
		return nil, err
	}
	padding[0] = byte(padLen)
	inBytes = append(padding, inBytes...)
	inBytesLen := len(inBytes)
	if inBytesLen%aes.BlockSize != 0 {
		return nil, ERR_CBC_PAD
	}

	ciphertext := make([]byte, aes.BlockSize+inBytesLen)
	initializationVector := ciphertext[:aes.BlockSize]
	_, err = rand.Reader.Read(initializationVector)
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCBCEncrypter(block, initializationVector)
	cfb.CryptBlocks(ciphertext[aes.BlockSize:], inBytes)

	return ciphertext, nil
}

func encryptedFileName(fp string) string {
	sp := strings.Split(path.Base(fp), ".")

	new_name := sp[0] + ".vault"
	return new_name
}
