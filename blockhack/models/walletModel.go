package models

type Wallet struct {
	PublicKey  string
	UserType   int
	DataAccess []PermList
}
