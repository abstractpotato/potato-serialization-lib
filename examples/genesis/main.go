package main

import (
	"fmt"
	"github.com/abstractpotato/potato-serialization-lib/psl"
)

// this is an example of a hard coded genesis block
// does not yet include signatures

func main() {
	// genesis validator request used to identify the initial validator
	validatorRequest := psl.NewRequest()
	validatorRequest.Ticker = "GENESIS"
	validatorRequest.Url = "https://starch.one"
	validatorRequest.Addr = "test_addr"
	validatorRequest.Relays = []string{
		"http://0.0.0.0:5001",
		"http://0.0.0.0:5002",
	}
	validatorRequestJSON, _ := validatorRequest.ToJSON()
	fmt.Printf("Genesis Validator Request:\n %s\n\n", validatorRequestJSON)

	// initial protocol parameters
	params := psl.NewParams()
	params.Network = 0
	params.MaxBlockHeaderSize = 256
	params.MaxBlockBodySize = 2048000000
	params.MaxTxSize = 4096
	params.TxFeePerByte = 430
	params.SlotsPerEpoch = 432000
	params.SlotTimeInMs = 1000
	params.ProtocolVersion = 0
	paramsJSON, _ := params.ToJSON()
	fmt.Printf("Genesis Params:\n %+s\n\n", paramsJSON)

	context := psl.NewContext()
	context.Params = params
	context.Epoch = 0
	context.Slot = 0
	context.Tip = 0
	contextJSON, _ := context.ToJSON()
	fmt.Printf("Genesis Context:\n %s\n\n", contextJSON)

	// convert the validator request to cbor and format into TxData
	validatorRequestData := psl.TxData{}
	validatorRequestData.Tag = "genesis_validator_request"
	validatorRequestCBOR, _ := validatorRequest.ToCBOR()
	validatorRequestData.Data = validatorRequestCBOR
	validatorRequestData.Type = 0

	// convert parameters to cbor and format into TxData
	paramsData := psl.TxData{}
	paramsData.Tag = "genesis_params"
	paramsCBOR, _ := params.ToCBOR()
	paramsData.Data = paramsCBOR
	paramsData.Type = 0

	// include a placeholder certificate for the genesis request
	certificate := psl.NewCertificate()
	certificate.RequestHash = "genesis"
	certificate.Signature = []byte("genesis_signature")
	certificateData := psl.TxData{}
	certificateData.Tag = "genesis_certificate"
	certificateCBOR, _ := certificate.ToCBOR()
	certificateData.Data = certificateCBOR
	certificateData.Type = 0

	// add data into transaction body
	transaction := psl.NewTransaction()
	transaction.Body.Epoch = context.Epoch
	transaction.Body.TTL = context.Slot + params.SlotsPerEpoch
	transaction.Body.Timestamp = 0
	transaction.Body.Network = params.Network
	transaction.Body.Data = append(transaction.Body.Data, validatorRequestData)
	transaction.Body.Data = append(transaction.Body.Data, paramsData)
	transaction.Body.Data = append(transaction.Body.Data, certificateData)
	transaction.Hash() // generate transaction hash

	txJSON, _ := transaction.ToJSON()
	fmt.Printf("Genesis Transaction:\n %s\n\n", txJSON)

	// add transaction into block body
	block := psl.NewBlock()
	block.Header.ID = 0
	block.Header.Validator = validatorRequest.Addr
	block.Header.Witnesses = append(block.Header.Witnesses, psl.Witness{
		Validator: validatorRequest.Addr,
		Signature: []byte("genesis_block_signature"),
	})
	// add validator signature as witness
	block.Body.PreviousHash = "0000"
	block.Body.VRF = "genesis"
	block.Body.Epoch = context.Epoch
	block.Body.Slot = context.Slot
	block.Body.Transactions = append(block.Body.Transactions, transaction)
	block.Body.Fees = 0
	block.Body.Timestamp = 0
	block.Hash() // generate block hash

	blockJSON, _ := block.ToJSON()
	fmt.Printf("Genesis Block:\n %s\n\n", blockJSON)
}
