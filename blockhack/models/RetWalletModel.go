package models

type RetWalletError struct {
	Error string
}

type RetWallet struct {
	WalPubKey     string
	WalPrivateKey string
	Money         int
}
