package main

import (
	"fmt"
	"github.com/abstractpotato/potato-serialization-lib/psl"
)

func main() {
	paramTest()
	contextTest()
	requestTest()
	certificateTest()
	blockTest()
	txTest()
}

func paramTest() {
	params := psl.NewParams()

	fmt.Printf("Original Params:\n %+v\n\n", params)

	paramHex, _ := params.ToHex()
	fmt.Printf("Params Hex:\n %s\n\n", paramHex)

	paramJSON, _ := params.ToJSON()
	fmt.Printf("Params JSON:\n %s\n\n", paramJSON)

	paramCBOR, _ := params.ToCBOR()
	fmt.Printf("Params CBOR:\n %s\n\n", paramCBOR)

	newParamsCBOR, _ := psl.ParamsFromCBOR(paramCBOR)
	fmt.Printf("Reconstructed Params CBOR:\n %+v\n\n", newParamsCBOR)

	newParamsHex, _ := psl.ParamsFromHex(paramHex)
	fmt.Printf("Reconstructed Params HEX:\n %+v\n\n", newParamsHex)
}

func contextTest() {
	context := psl.NewContext()
	context.Params = psl.NewParams()
	context.Params.Network = 0
	context.Params.TxFeePerByte = 430
	context.Epoch = 1
	context.Slot = 42
	context.Tip = 7

	fmt.Printf("Original Context:\n %+v\n\n", context)

	contextHex, _ := context.ToHex()
	fmt.Printf("Context Hex:\n %s\n\n", contextHex)

	contextJSON, _ := context.ToJSON()
	fmt.Printf("Context JSON:\n %s\n\n", contextJSON)

	contextCBOR, _ := context.ToCBOR()
	fmt.Printf("Context CBOR:\n %s\n\n", contextCBOR)

	newContextCBOR, _ := psl.ContextFromCBOR(contextCBOR)
	fmt.Printf("Reconstructed Context CBOR:\n %+v\n\n", newContextCBOR)

	newContextHex, _ := psl.ContextFromHex(contextHex)
	fmt.Printf("Reconstructed Context HEX:\n %+v\n\n", newContextHex)

}

func requestTest() {
	request := psl.NewRequest()
	request.Ticker = "GENESIS"
	request.Url = "https://starch.one"
	request.Addr = "test_addr"
	request.Relays = []string{"http://0.0.0.0:5001"}

	fmt.Printf("Original Request:\n %+v\n\n", request)

	requestHex, _ := request.ToHex()
	fmt.Printf("Request Hex:\n %s\n\n", requestHex)

	requestJSON, _ := request.ToJSON()
	fmt.Printf("Request JSON:\n %s\n\n", requestJSON)

	requestCBOR, _ := request.ToCBOR()
	fmt.Printf("Request CBOR:\n %s\n\n", requestCBOR)

	newRequestCBOR, _ := psl.RequestFromCBOR(requestCBOR)
	fmt.Printf("Reconstructed Request CBOR:\n %+v\n\n", newRequestCBOR)

	newRequestHex, _ := psl.RequestFromHex(requestHex)
	fmt.Printf("Reconstructed Request HEX:\n %+v\n\n", newRequestHex)
}

func certificateTest() {
	certificate := psl.NewCertificate()
	certificate.RequestHash = "request_hash"
	certificate.Signature = []byte("signature")

	fmt.Printf("Original Certificate:\n %+v\n\n", certificate)

	certificateHex, _ := certificate.ToHex()
	fmt.Printf("Certificate Hex:\n %s\n\n", certificateHex)

	certificateJSON, _ := certificate.ToJSON()
	fmt.Printf("Certificate JSON:\n %s\n\n", certificateJSON)

	certificateCBOR, _ := certificate.ToCBOR()
	fmt.Printf("Certificate CBOR:\n %s\n\n", certificateCBOR)

	newCertificateCBOR, _ := psl.CertificateFromCBOR(certificateCBOR)
	fmt.Printf("Reconstructed Certificate CBOR:\n %+v\n\n", newCertificateCBOR)

	newCertificateHex, _ := psl.CertificateFromHex(certificateHex)
	fmt.Printf("Reconstructed Certificate HEX:\n %+v\n\n", newCertificateHex)
}

func blockTest() {
	block := psl.NewBlock()
	block.Hash()

	fmt.Printf("Original Block:\n %+v\n\n", block)

	blockHex, _ := block.ToHex()
	fmt.Printf("Block Hex:\n %s\n\n", blockHex)

	// to json
	blockJSON, _ := block.ToJSON()
	fmt.Printf("Block JSON:\n %s\n\n", blockJSON)

	// Convert struct to CBOR
	blockCBOR, _ := block.ToCBOR()
	fmt.Printf("Block CBOR:\n %s\n\n", blockCBOR)

	// rebuild struct from CBOR
	newBlockCBOR, _ := psl.BlockFromCBOR(blockCBOR)
	fmt.Printf("Reconstructed Block CBOR:\n %+v\n\n", newBlockCBOR)

	// rebuild block from HEX
	newBlockHex, _ := psl.BlockFromHex(blockHex)
	fmt.Printf("Reconstructed Block HEX:\n %+v\n\n", newBlockHex)
}

func txTest() {
	tx := psl.NewTransaction()
	tx.Body.Data = append(tx.Body.Data, psl.TxData{
		Tag:  "example",
		Data: []byte("payload"),
		Type: 0,
	})
	tx.Hash()

	fmt.Printf("Original Transaction:\n %+v\n\n", tx)

	txHex, _ := tx.ToHex()
	fmt.Printf("Transaction Hex:\n %s\n\n", txHex)

	// to json
	txJSON, _ := tx.ToJSON()
	fmt.Printf("Transaction JSON:\n %s\n\n", txJSON)

	// Convert struct to CBOR
	txCBOR, _ := tx.ToCBOR()
	fmt.Printf("Transaction CBOR:\n %s\n\n", txCBOR)

	// rebuild struct from CBOR
	newTransactionCBOR, _ := psl.TransactionFromCBOR(txCBOR)
	fmt.Printf("Reconstructed Transaction CBOR:\n %+v\n\n", newTransactionCBOR)

	// rebuild tx from HEX
	newTransactionHex, _ := psl.TransactionFromHex(txHex)
	fmt.Printf("Reconstructed Transaction HEX:\n %+v\n\n", newTransactionHex)
}
