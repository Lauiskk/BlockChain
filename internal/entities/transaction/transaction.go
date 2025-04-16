package transaction

const subsidy = 50

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
type TXOutput struct {
	Value        int
	ScriptPubKey string
}
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}
