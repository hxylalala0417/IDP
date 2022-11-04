package zrx_quote

import (
	"IDP/internal/logger"
	"IDP/pkg/consts"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type ZrxQuoter struct {
}

//Communicates with API and returns the response
func (p *ZrxQuoter) requestQuoteAPI(queryParams string) (*Response, error) {
	// queryParams: ?sellToken=%v&sellAmount=%v&buyToken=%v
	url := consts.API_ENGPOINT_Mainnet + consts.API_QUOTE + queryParams

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error: GET failed")
		logger.Logger.Warnw("Get error", "err", "GET failed")
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error: Sending request or receiving respond failed")
		logger.Logger.Warnw("Do error", "err", "Do function failed")
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error: Read failed")
		logger.Logger.Warnw("Read error", "err", "Read failed")
		return nil, err
	}

	response := &Response{}
	if err := json.Unmarshal([]byte(body), response); err != nil {
		fmt.Println("error: Unmarshal failed")
		logger.Logger.Warnw("Unmarshal error", "err", "Unmarshal failed")
		return nil, err
	}
	//fmt.Println(string(body))
	return response, nil
}

//Take the response, extract the BuyAmount field and returns it as a big.int
func (p *ZrxQuoter) getOutputAmount(sellToken, sellAmount, buyToken string) (*big.Int, error) {
	resp, err := p.requestQuoteAPI(fmt.Sprintf("?sellToken=%v&sellAmount=%v&buyToken=%v", sellToken, sellAmount, buyToken))
	if err != nil {
		fmt.Println("error: Unmarshal failed")
		logger.Logger.Warnw("Unmarshal error", "err", "Unmarshal failed")
		return nil, err
	}
	i := new(big.Int)
	i.SetString(resp.BuyAmount, 10)
	return i, nil
}

//func (p *ZrxQuoter) getInputAmount(sellToken, , buyToken string,) (*big.Int, error) {
//	resp, err := p.requestQuoteAPI(fmt.Sprintf("?sellToken=%v&buyAmount=%v&buyToken=%v", sellToken, sellAmount, buyToken))
//	if err!=nil{
//		return nil, err
//	}
//	return resp.BuyAmount, nil
//}
