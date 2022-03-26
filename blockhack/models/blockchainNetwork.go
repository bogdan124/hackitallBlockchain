package models

type BlockchainNetwork struct {
	Blocks       []Block
	GenesisBlock Block
	Wallet       []Wallet
}
