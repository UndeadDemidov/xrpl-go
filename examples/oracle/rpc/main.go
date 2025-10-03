package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Peersyst/xrpl-go/examples/clients"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/rpc/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func safeInt64ToUint32(value int64) uint32 {
	if value < 0 {
		return 0
	}
	if value > int64(^uint32(0)) {
		return ^uint32(0) // max uint32 value
	}
	return uint32(value)
}

func printJSON(data interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error marshaling to JSON: %s\n", err)
		return
	}
	fmt.Println(string(jsonBytes))
}

func main() {
	//
	// Configure client
	//
	fmt.Println("⏳ Setting up testnet RPC client...")
	client := clients.GetTestnetRPCClient()

	//
	// Configure wallets
	//
	fmt.Println("⏳ Setting up wallets...")
	oracleIssuer, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Printf("❌ Error creating oracle issuer wallet: %s\n", err)
		return
	}
	err = client.FundWallet(&oracleIssuer)
	if err != nil {
		fmt.Printf("❌ Error funding oracle issuer wallet: %s\n", err)
		return
	}
	fmt.Println("💸 Oracle issuer wallet funded!")
	fmt.Println()

	//
	// Create oracle set transaction
	//
	fmt.Println("⏳ Creating oracle set transaction...")

	// 1 minute ago
	lastUpdatedTime := safeInt64ToUint32(time.Now().Add(-time.Second).Unix())
	oracleDocumentID := uint32(1)

	oracleSet := transaction.OracleSet{
		BaseTx: transaction.BaseTx{
			Account: oracleIssuer.ClassicAddress,
		},
		OracleDocumentID: oracleDocumentID,
		LastUpdatedTime:  lastUpdatedTime,
		URI:              hex.EncodeToString([]byte("https://example.com")),
		Provider:         hex.EncodeToString([]byte("Chainlink")),
		AssetClass:       hex.EncodeToString([]byte("currency")),
		PriceDataSeries: []ledger.PriceDataWrapper{
			{
				PriceData: ledger.PriceData{
					BaseAsset:  currency.ConvertStringToHex("ACGB"),
					QuoteAsset: "USD",
					AssetPrice: 123,
					Scale:      3,
				},
			},
		},
	}

	flatOracleSet := oracleSet.Flatten()

	fmt.Println("📄 Oracle Set Transaction JSON:")
	printJSON(flatOracleSet)
	fmt.Println()

	response, err := client.SubmitTxAndWait(flatOracleSet, &types.SubmitOptions{
		Wallet:   &oracleIssuer,
		Autofill: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Oracle set transaction submitted")
	fmt.Printf("🌐 Hash: %s\n", response.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", response.Validated)

	if !response.Validated {
		fmt.Println("❌ Oracle set transaction failed")
		return
	}
	fmt.Println()

	// Delete oracle
	fmt.Println("⏳ Deleting oracle...")

	oracleDelete := transaction.OracleDelete{
		BaseTx: transaction.BaseTx{
			Account: oracleIssuer.ClassicAddress,
		},
		OracleDocumentID: oracleDocumentID,
	}

	flatOracleDelete := oracleDelete.Flatten()

	fmt.Println("📄 Oracle Delete Transaction JSON:")
	printJSON(flatOracleDelete)
	fmt.Println()

	responseDelete, err := client.SubmitTxAndWait(flatOracleDelete, &types.SubmitOptions{
		Wallet:   &oracleIssuer,
		Autofill: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Oracle deleted")
	fmt.Printf("🌐 Hash: %s\n", responseDelete.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", responseDelete.Validated)
}
