package main

import (
	"context"
	"fmt"

	"github.com/KyberNetwork/service-framework/pkg/client/grpcclient"
	onchainpricev1 "github.com/KyberNetwork/service-proto-public/gen/onchainprice/v1"
	"google.golang.org/protobuf/encoding/prototext"
)

const (
	HeaderXChainId  = "X-Chain-Id"
	ChainIdEthereum = "1"
)

func main() {
	ctx := context.Background()

	cli, err := grpcclient.New(onchainpricev1.NewOnchainPriceServiceClient,
		grpcclient.WithConfig(&grpcclient.Config{
			BaseURL:  "onchain-price.kyberengineering.io",
			Headers:  map[string]string{HeaderXChainId: ChainIdEthereum},
			ClientID: "service-proto-public",
		}))
	if err != nil {
		panic(err)
	}

	fmt.Println("Fetching price of wstETH by WETH")

	resp, err := cli.C.ListPrices(ctx, &onchainpricev1.ListPricesRequest{
		Tokens: []string{"0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"}, // wstETH
		Quotes: []string{"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"}, // WETH
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(prototext.Format(resp))
}
