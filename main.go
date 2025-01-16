package main

import (
	"github.com/google/go-tdx-guest/client"
	labi "github.com/google/go-tdx-guest/client/linuxabi"
	"github.com/google/go-tdx-guest/verify"
)

func main() {
	device, err := client.OpenDevice()
	if err != nil {
		panic(err)
	}
	defer device.Close()

	var reportData64 [labi.TdReportDataSize]byte
	tdxQuoteProvider, err := client.GetQuoteProvider()
	if err != nil {
		panic(err)
	}

	quote, err := client.GetQuote(tdxQuoteProvider, reportData64)
	if err != nil {
		panic(err)
	}
	err = verify.TdxQuote(quote, nil)
	if err != nil {
		panic(err)
	}
}
