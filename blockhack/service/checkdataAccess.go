package service

import "blockhack/models"

func CheckDataAccess(iterator int, filenamePath string, BlocksChains models.BlockchainNetwork) int {
	var writeType int
	writeType = 0 //no permission
	if len(BlocksChains.Wallet[iterator].DataAccess) != 0 {
		for i := 0; i < len(BlocksChains.Wallet[iterator].DataAccess); i++ {
			if BlocksChains.Wallet[iterator].DataAccess[i].Filename == filenamePath {
				//check if you have permissions
				if BlocksChains.Wallet[iterator].DataAccess[i].Write == 1 {
					writeType = 1 //have only write permissions
				}
				if BlocksChains.Wallet[iterator].DataAccess[i].Read == 1 {
					writeType = 2 //have only read permissions
				}
				if BlocksChains.Wallet[iterator].DataAccess[i].Read == 1 {
					if BlocksChains.Wallet[iterator].DataAccess[i].Write == 1 {
						writeType = 3 //have read write permissions
					}
				}

			}
		}
	} else {
		writeType = 0
	}

	return writeType
}
