package explorer

//SendAddressInput SendAddressInput
type SendAddressInput struct {
	Address               string
	Amount                float32
	Comment               string
	CommentTo             string
	SubtractFeeFromAmount bool
	Replaceable           bool
	ConfTarget            int64
	EstimateMode          string
	AvoidReuse            bool
}

//SendAddressOutput SendAddressOutput
type SendAddressOutput struct {
	TxID      string //hex
	FeeReason string
}

//QueryOptions QueryOptions
type QueryOptions struct {
	MinimumAmount    interface{}
	MaximumAmount    interface{}
	MaximumCount     interface{}
	MinimumSumAmount interface{}
}

//ListUnspentInput ListUnspentInput
type ListUnspentInput struct {
	Minconf       int64
	Maxconf       int64
	Addresses     []string
	IncludeUnsafe bool
	QueryOptions  QueryOptions
}

//ListUuspentOutput ListUuspentOutput
type ListUuspentOutput struct {
	TxID          string //hex
	Vout          float64
	Address       string
	Label         string
	ScriptPubKey  string
	Amount        float64
	Confirmations int32
	RedeemScript  string //hex
	WitnessScript string
	Spendable     bool
	Solvable      bool
	Reused        bool
	Desc          string
	Safe          bool
}

//RawTxInputs RawTxInputs
type RawTxInputs struct {
	TxID     string //hex
	Vout     float64
	Sequence int32
}

//RawTxOutputs RawTxOutputs
type RawTxOutputs struct {
	AddressAmount map[string]interface{}
	Data          string //hex
}

//CreateRawTxInput CreateRawTxInput
type CreateRawTxInput struct {
	Inputs      RawTxInputs
	Outputs     RawTxOutputs
	Locktime    int32
	Replaceable bool
}

//Explorer Explorer
type Explorer interface {
	GetNewAddress(label string, atype string) string
	SendToAddress(inp SendAddressInput) *SendAddressOutput
	ListUnspent(inp *ListUnspentInput) *[]ListUuspentOutput
	CreateRawTransaction(inp *CreateRawTxInput) string //hex
	//DecodeRawTransaction()
	//SignRawTransactionWithKey()
	//SignRawTransactionWithWallet()
	//ValidateAddress()
}

//GO111MODULE=on go mod init github.com/Ulbora/BitcoinCoreExplorer
