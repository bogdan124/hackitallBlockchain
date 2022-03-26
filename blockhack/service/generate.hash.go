package service

import (
	"blockhack/models"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func GenerateHash(blk models.Block, iteratorBlock int) string {

	dataToPut := blk.Timestamp + fmt.Sprint(blk.HashDevice) + blk.PrevBlock + string(iteratorBlock)
	hash := sha512.New()

	hash.Write([]byte(dataToPut))
	ret := hash.Sum(nil)
	return hex.EncodeToString(ret)
}

func GenHashOfAString(val string) string {
	dataToPut := val
	hash := sha512.New()

	hash.Write([]byte(dataToPut))
	ret := hash.Sum(nil)
	return hex.EncodeToString(ret)
}
