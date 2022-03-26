package service

import (
	"blockhack/models"
	"time"
)

/*first block in the network*/
func InitGenesisBlock(blockchain models.BlockchainNetwork) models.BlockchainNetwork {
	var FirstBlock models.Block

	blockchain.Wallet = make([]models.Wallet, 0)
	t := time.Now()
	FirstBlock.PrevBlock = ""
	FirstBlock.Timestamp = t.String()
	FirstBlock.HashDevice = GenerateHash(FirstBlock, 0)
	blockchain.Blocks = append(blockchain.Blocks, FirstBlock)
	FirstBlock.MoreDeviceBlocks = make([]models.Block, 0)
	FirstBlock.ModifiedData = make([]models.ModData, 0)
	return blockchain
}
