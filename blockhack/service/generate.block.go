package service

import (
	"blockhack/models"
	"time"
)

/*first block in the network*/
func NetworkBlock(blockchain []models.Block, addData models.BlockData, iteratorBlock int, miner string, network models.BlockchainNetwork) (string, models.Block) {
	var FirstBlock models.Block

	t := time.Now()
	FirstBlock.Timestamp = t.String()
	if len(addData.DeviceName) != 0 {
		FirstBlock.DeviceName = addData.DeviceName
	}

	if len(blockchain) != 0 {
		FirstBlock.PrevBlock = blockchain[len(blockchain)-1].HashDevice
		FirstBlock.HashDevice = GenerateHash(blockchain[len(blockchain)-1], iteratorBlock)
	} else {
		FirstBlock.PrevBlock = network.GenesisBlock.HashDevice
		FirstBlock.HashDevice = GenerateHash(network.GenesisBlock, iteratorBlock)
	}
	FirstBlock.MoreDeviceBlocks = make([]models.Block, 0)
	FirstBlock.ModifiedData = make([]models.ModData, 0)
	FirstBlock.DataHold = make([]models.DataHold, 0)

	//	blockchain = append(blockchain, FirstBlock)

	return FirstBlock.HashDevice, FirstBlock
}
