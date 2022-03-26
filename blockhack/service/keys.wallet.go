package service

import (
	"crypto/sha512"
	"encoding/hex"
)

func GeneratePrivateKey(getData string) string {

	dataToPut := getData
	hash := sha512.New()

	hash.Write([]byte(dataToPut))
	ret := hash.Sum(nil)
	return hex.EncodeToString(ret)
}

func GeneratePubKey(privateKey string) string {

	dataToPut := privateKey
	hash := sha512.New()

	hash.Write([]byte(dataToPut))
	ret := hash.Sum(nil)
	return hex.EncodeToString(ret)
}
