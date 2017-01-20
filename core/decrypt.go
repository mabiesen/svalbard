package core

import(
	"io/ioutil"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
)

func DecryptIntf(fp string) error {

	randomSymmetricKey := []byte("my secret 256 bit symmetric key ")

	if fp == "" {
		return NO_FILE_ERROR
	}

	dat, err := ioutil.ReadFile(fp)

	if err != nil {
		return READ_FILE_ERROR
	}

	// decrypt and print the original file
	superLargeFile2, _ := decryptCBC(randomSymmetricKey, dat)

	fmt.Print(string(superLargeFile2))



	return nil
}

func decryptCBC(symmetricKey []byte, inBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(symmetricKey)
	if err != nil {
		return nil, err
	}

	initializationVector := inBytes[:aes.BlockSize]
	inBytes = inBytes[aes.BlockSize:]
	cfb := cipher.NewCBCDecrypter(block, initializationVector)
	cfb.CryptBlocks(inBytes, inBytes)
	padLen := uint8(inBytes[0])
	inBytes = inBytes[padLen:]
	return inBytes, nil
}
