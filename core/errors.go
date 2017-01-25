package core

import (
	"errors"
)

var NO_FILE_ERROR = errors.New("No file specified.")
var READ_FILE_ERROR = errors.New("Could not read file provided.")
var DELETE_UNENCRYPTED_FILE_ERROR = errors.New("Could not delete the original file.")
var CREATE_ENCRYPTED_FILE_ERROR = errors.New("Could not write encrypted file. Most likely due to lack of permissions.")
var NO_KEY_NAME_PROVIDED = errors.New("No key name provided.")
var KEY_GEN_ERROR = errors.New("Could not generate 256 bit key.")
var KEY_ALREADY_EXISTS = errors.New("A key with that name already exists.")
var ERR_CBC_PAD = errors.New("inBytes not a multiple of aes.BlockSize")
var ERR_NOT_PEM = errors.New("not a PEM")
