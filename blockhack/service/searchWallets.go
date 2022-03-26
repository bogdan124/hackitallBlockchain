package service

import (
	"blockhack/models"
	"fmt"
)

func SearchWallets(BlkChains []models.Wallet, PublicKey string) bool {
	fin := 0
	fmt.Println("sdf", BlkChains)
	if len(BlkChains) == 0 {
		fmt.Println("sdewf", BlkChains)
		return false
	}
	for i := 0; i < len(BlkChains); i++ {
		fmt.Println(BlkChains[i].PublicKey, PublicKey)
		if BlkChains[i].PublicKey == PublicKey {
			fin = 1
			fmt.Println(BlkChains[i].PublicKey, PublicKey)
			return true
		} else {
			fmt.Println("debugging")
		}
	}
	if fin == 1 {
		return true
	} else {
		return false
	}
}
