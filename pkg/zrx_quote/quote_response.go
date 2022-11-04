package zrx_quote

//Define struct according to the response body

type Response struct {
	ChainID            int       `json:"chainId"`
	Price              string    `json:"price"`
	GuaranteedPrice    string    `json:"guaranteedPrice"`
	To                 string    `json:"to"`
	Data               string    `json:"data"`
	Value              string    `json:"value"`
	Gas                string    `json:"gas"`
	EstimatedGas       string    `json:"estimatedGas"`
	GasPrice           string    `json:"gasPrice"`
	ProtocolFee        string    `json:"protocolFee"`
	MinimumProtocolFee string    `json:"minimumProtocolFee"`
	BuyTokenAddress    string    `json:"buyTokenAddress"`
	SellTokenAddress   string    `json:"sellTokenAddress"`
	BuyAmount          string    `json:"buyAmount"`
	SellAmount         string    `json:"sellAmount"`
	Sources            []Sources `json:"sources"`
	Orders             []Orders  `json:"orders"`
	AllowanceTarget    string    `json:"allowanceTarget"`
	SellTokenToEthRate string    `json:"sellTokenToEthRate"`
	BuyTokenToEthRate  string    `json:"buyTokenToEthRate"`
}
type Sources struct {
	Name       string `json:"name"`
	Proportion string `json:"proportion"`
}
type FillData struct {
	TokenAddressPath []string `json:"tokenAddressPath"`
	Router           string   `json:"router"`
}
type Orders struct {
	MakerToken   string   `json:"makerToken"`
	TakerToken   string   `json:"takerToken"`
	MakerAmount  string   `json:"makerAmount"`
	TakerAmount  string   `json:"takerAmount"`
	FillData     FillData `json:"fillData"`
	Source       string   `json:"source"`
	SourcePathID string   `json:"sourcePathId"`
	Type         int      `json:"type"`
}
