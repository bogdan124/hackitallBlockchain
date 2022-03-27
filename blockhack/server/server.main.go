package server

import (
	"blockhack/models"
	"blockhack/service"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var BlocksChains models.BlockchainNetwork

func SetUpServer() {
	if len(BlocksChains.Blocks) == 0 {
		BlocksChains = service.InitGenesisBlock(BlocksChains)
	}
	http.HandleFunc("/get", ShowBlocks)
	http.HandleFunc("/mine", SetBlock)
	http.HandleFunc("/createWallet", CreateWallet)
	http.HandleFunc("/wallet", ViewWallet)
	http.HandleFunc("/get/device/name", GetDeviceName)
	http.HandleFunc("/put/device/data", PutDeviceData)
	//	http.HandleFunc("/transaction", CreateTransaction)

	http.ListenAndServe(":8080", nil)
}

func ShowBlocks(w http.ResponseWriter, req *http.Request) {
	fmt.Print(len(BlocksChains.Blocks))
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)

	if len(BlocksChains.Blocks) == 0 {
		BlocksChains = service.InitGenesisBlock(BlocksChains)
	}

	response, err := json.Marshal(BlocksChains)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500"))
	}
	w.Write(response)

	fmt.Print(BlocksChains)
}

func SetBlock(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	var viewSetBlock models.BlockData
	viewBlock := json.NewDecoder(req.Body)
	viewBlock.Decode(&viewSetBlock)
	///if verify != nil {
	///	return
	//}
	defer req.Body.Close()

	getVerification := service.SearchWallets(BlocksChains.Wallet, viewSetBlock.Wallet)

	if getVerification == true {
		for i := 0; ; i++ {
			iteratorBlock := i
			hashBlock, actualBlock := service.NetworkBlock(BlocksChains.Blocks, viewSetBlock, iteratorBlock, viewSetBlock.Wallet, BlocksChains)
			//fmt.Println(service.ValidateBlock(hashBlock, 2, "0"))
			if service.ValidateBlock(hashBlock, 4, "0") {
				BlocksChains.Blocks = append(BlocksChains.Blocks, actualBlock)
				//	data := BlocksChains.Wallets[viewSetBlock.Wallet]
				//getCoins := BlocksChains.Wallets[viewSetBlock.Wallet].Coins
				//data.Coins = getCoins + 10
				//BlocksChains.Wallets[viewSetBlock.Wallet] = data
				break
			} else {
				fmt.Println(hashBlock)
			}
		}
	} else {
		fmt.Print("bad wallet")
	}

}

func CreateWallet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	//var viewSetBlock models.BlockData
	var walroute models.WalletAPI
	var wallet models.Wallet
	viewBlock := json.NewDecoder(req.Body)
	viewBlock.Decode(&walroute)
	privateKEY := service.GeneratePrivateKey(walroute.DataToGenerateKey)
	publicKey := service.GeneratePubKey(privateKEY)
	fmt.Println("wallet")

	checkIfKeyIsGood := service.SearchWallets(BlocksChains.Wallet, publicKey)
	fmt.Println("WalletValidator:", checkIfKeyIsGood)
	if checkIfKeyIsGood == false {
		//	viewSetBlock.Coins = 0
		//BlocksChains.Wallets[viewSetBlock.Wallet] = viewSetBlock
		wallet.PublicKey = publicKey
		wallet.DataAccess = make([]models.PermList, 0)
		fmt.Println("okkk")
		BlocksChains.Wallet = append(BlocksChains.Wallet, wallet)
		retWallet := models.RetWallet{}
		privateKey := privateKEY
		retWallet.WalPrivateKey = privateKey
		retWallet.WalPubKey = publicKey
		json.NewEncoder(w).Encode(retWallet)
	} else {
		RetWallError := models.RetWalletError{}
		RetWallError.Error = "Wallet already created"
		json.NewEncoder(w).Encode(RetWallError)
	}

}

func ViewWallet(w http.ResponseWriter, req *http.Request) {

}

func GetDeviceName(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var walroute models.GetDeviceName

	viewBlock := json.NewDecoder(req.Body)
	viewBlock.Decode(&walroute)
	publicKey := service.GeneratePubKey(walroute.PrivateKey)
	fmt.Println("Pereche chei:", publicKey, walroute.PrivateKey)
	find := 0
	RetBlock := models.Block{}
	for j := 0; j < len(BlocksChains.Wallet); j++ {
		if BlocksChains.Wallet[j].PublicKey == publicKey {
			find = 1
		}
	}
	if find == 1 {
		find = 0
		for i := 0; i < len(BlocksChains.Blocks); i++ {
			if BlocksChains.Blocks[i].DeviceName == walroute.DeviceToAccess {
				find = 1
				RetBlock = BlocksChains.Blocks[i]
			}
		}
		if find == 1 {
			json.NewEncoder(w).Encode(RetBlock)
		} else {
			json.NewEncoder(w).Encode(RetBlock)
		}
	} else {
		json.NewEncoder(w).Encode(RetBlock)
	}

}

func PutDeviceData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	var walroute models.PutDeviceData
	var dataToBeReturned int
	var getPermissions int
	viewBlock := json.NewDecoder(req.Body)
	viewBlock.Decode(&walroute)
	RetBlock := models.DefaultRet{}
	publicKey := service.GeneratePubKey(walroute.PrivateKey)
	for i := 0; i < len(BlocksChains.Wallet); i++ {
		if BlocksChains.Wallet[i].PublicKey == publicKey {
			dataToBeReturned = service.CheckDataAccess(i, walroute.FileData, BlocksChains)
			for j := 0; j < len(BlocksChains.Blocks); j++ {
				if BlocksChains.Blocks[j].DeviceName == walroute.DeviceName {
					dHold := models.DataHold{}
					dHold.Filename = append(dHold.Filename, walroute.Filename)
					dHold.Path = append(dHold.Path, walroute.Path)
					BlocksChains.Blocks[j].DataHold = append(BlocksChains.Blocks[j].DataHold, dHold)
					writeDataFileDisk(dHold.Filename[0], walroute.FileData)
				}
			}
			/*
				//full permissions
				if BlocksChains.Wallet[i].UserType == 1 {
					writeDataFileDisk(walroute.Filename, walroute.FileData)
					RetBlock.Message = "File is writted"
					json.NewEncoder(w).Encode(RetBlock)
					getPermissions = 1
				}
				//read write permissions
				if dataToBeReturned == 3 {
					writeDataFileDisk(walroute.Filename, walroute.FileData)
					RetBlock.Message = "File is writted"
					json.NewEncoder(w).Encode(RetBlock)
					getPermissions = 1
				}
				//write permissions
				if dataToBeReturned == 1 {
					writeDataFileDisk(walroute.Filename, walroute.FileData)
					RetBlock.Message = "File is writted"
					json.NewEncoder(w).Encode(RetBlock)
					getPermissions = 1
				}

			*/
		}
	}
	if getPermissions != 0 {
		RetBlock.Message = "Can't write on disk no"
		json.NewEncoder(w).Encode(RetBlock)
	}
	fmt.Println(walroute, dataToBeReturned)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func writeDataFileDisk(filenameTowrite string, content string) {
	f, _ := os.Create("files/" + filenameTowrite)

	// Create a new writer.
	w := bufio.NewWriter(f)

	// Write a string to the file.
	w.WriteString(content)

	// Flush.
	w.Flush()
}

/*
func CreateTransaction(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var viewSetBlock models.TransactionRequest

	viewBlock := json.NewDecoder(req.Body)
	viewBlock.Decode(&viewSetBlock)

	getWallet := service.GeneratePubKey(viewSetBlock.PrivateKey)

	_, err := BlocksChains.Wallets[getWallet]
	if err {

		data := BlocksChains.Wallets[getWallet]
		data.Coins = data.Coins - viewSetBlock.Money
		if data.Coins < 0 {

			RetWallError := models.RetWalletError{}
			RetWallError.Error = "You don't have enough funds"
			json.NewEncoder(w).Encode(RetWallError)
		} else if data.Coins >= 0 {

			var tx1 models.InitTrans
			tx1.Coins = viewSetBlock.Money
			tx1.Wallet = getWallet

			var tx2 models.DestTrans
			tx2.Wallet = viewSetBlock.DestinationWallet

			var tx models.Transaction
			tx.Id = 1
			tx.WalletSource = tx1
			tx.WalletDest = tx2
			_, err := BlocksChains.Wallets[tx2.Wallet]
			if err {

				data1 := BlocksChains.Wallets[tx2.Wallet]
				data1.Coins += viewSetBlock.Money
				BlocksChains.Wallets[tx2.Wallet] = data1
				BlocksChains.Wallets[getWallet] = data
				putData := BlocksChains.Blocks[0].Transactions
				BlocksChains.Blocks[0].Transactions = append(putData, tx)
			}

		}

	} else {
		RetWallError := models.RetWalletError{}
		RetWallError.Error = "No Wallet created"
		json.NewEncoder(w).Encode(RetWallError)
	}
}
*/
