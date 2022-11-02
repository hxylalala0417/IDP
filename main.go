package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	BuyAmount string `json:"buyAmount"`
}

type ZrxQuoter struct {
}

func (p *ZrxQuoter) getOutputAmount(sellToken, sellAmount, buyToken string) string {
	url := fmt.Sprintf("https://api.0x.org/swap/v1/quote?sellToken=%v&sellAmount=%v&buyToken=%v", sellToken, sellAmount, buyToken)

	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	res, _ := client.Do(req)

	//res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	response := Response{} // 解析json数据并将数据存储在response结构体中
	if err := json.Unmarshal([]byte(body), &response); err == nil {
		//fmt.Println(response)
	} else {
		fmt.Println(err)
	}
	//json.Unmarshal(body, &response)
	//fmt.Println(response)

	//fmt.Println(string(body))

	// body -> json -> 取出字段
	return response.BuyAmount
}

func main() {
	var zrx = new(ZrxQuoter)
	fmt.Println(zrx.getOutputAmount("ETH", "100000000000", "0x6b175474e89094c44da98b954eedeac495271d0f"))

}
