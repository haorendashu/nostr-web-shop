package utils

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/lightningnetwork/lnd/zpay32"
)

func LightningInvoiceParse(invoice string) (*zpay32.Invoice, error) {
	return zpay32.Decode(invoice, &chaincfg.MainNetParams)
}
