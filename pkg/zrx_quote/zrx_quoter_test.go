package zrx_quote

import (
	"IDP/internal/logger"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func Test(t *testing.T) {
	var zrx = new(ZrxQuoter)
	if err := logger.Init("ZrxQuoter.log", zapcore.WarnLevel); err != nil {
		fmt.Println("logger.Init failed", "err", err)
		os.Exit(1)
	}
	res, err := zrx.getOutputAmount("ETH", "10000000000000", "0x6b175474e89094c44da98b954eedeac495271d0f")
	//res, _ := zrx.getOutputAmount("ETH", "1", "0x6b175474e89094c44da98b954eedeac495271d0f")
	fmt.Println(res)
	//assert.Nil(t, res)
	assert.Nil(t, err)
}
