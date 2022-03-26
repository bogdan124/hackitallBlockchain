package models

type Block struct {
	Timestamp        string
	MoreDeviceBlocks []Block
	HashDevice       string
	DeviceName       string
	PrevBlock        string
	ModifiedData     []ModData
	DataHold         []DataHold
}
