package main

import (
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	//Loads the .env file and throws an error if it cannot load the variables from that file correctly
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

	//Grab your testnet account ID and private key from the .env file
	myAccountId, err := hedera.AccountIDFromString(os.Getenv("TESTNET_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	myPrivateKey, err := hedera.PrivateKeyFromStringEd25519(os.Getenv("TESTNET_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

  accountToFund, err := hedera.AccountIDFromString("0.0.3490973")
	//Transfer hbar from your testnet account to the new account
	transaction := hedera.NewTransferTransaction().
		AddHbarTransfer(myAccountId, hedera.HbarFrom(-9000, hedera.HbarUnits.Hbar)).
		AddHbarTransfer(accountToFund, hedera.HbarFrom(9000, hedera.HbarUnits.Hbar))

	//Submit the transaction to a Hedera network
	txResponse, err := transaction.Execute(client)

	if err != nil {
		panic(err)
	}
  fmt.Printf("tx: %v\n", txResponse)

}
