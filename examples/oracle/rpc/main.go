package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/rpc/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

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
	fmt.Println("⏳ Setting up client...")
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(cfg)
	fmt.Println("✅ Client configured!")
	fmt.Println()

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
	lastUpdatedTime := time.Now().Add(-time.Second).Unix()

	oracleSet := transaction.OracleSet{
		BaseTx: transaction.BaseTx{
			Account: oracleIssuer.ClassicAddress,
		},
		OracleDocumentID: 1,
		LastUpdatedTime:  uint32(lastUpdatedTime),
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
}
