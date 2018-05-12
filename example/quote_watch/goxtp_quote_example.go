package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pseudocodes/goxtp"
)

type QuoteConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	AccountKey string
	LogPath    string
}

var config QuoteConfig

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	flag.StringVar(&config.Host, "host", "120.27.164.138", "行情服务器地址")
	flag.IntVar(&config.Port, "port", 6002, "行情服务器端口")
	flag.StringVar(&config.Username, "username", "pseudo", "行情服务登录用户名")
	flag.StringVar(&config.Password, "password", "******", "行情服务登录密码")
	flag.StringVar(&config.AccountKey, "accountkey", "xxxxxxxx", "api key")
	flag.StringVar(&config.LogPath, "logpath", "./test", "相关日志输出路径")

	flag.Parse()
}

type GoXTPClient struct {
	Host string
	Port int

	Username   string
	Password   string
	AccountKey string

	ClientID int
	Path     string
}

func NewDirectorQuoteSpi(v interface{}) goxtp.QuoteSpi {
	return goxtp.NewDirectorQuoteSpi(v)
}

type GoXTPQuoteSpi struct {
	Client GoXTPClient
}

//OnDisconnected 当客户端与行情后台通信连接断开时，该方法被调用
func (p *GoXTPQuoteSpi) OnDisconnected(reason int) {
	log.Printf("OnDisconnected: %v\n", reason)
}

//OnError 错误应答
func (p *GoXTPQuoteSpi) OnError(errorInfo goxtp.XTPRI) {
	log.Printf("ErrorID %v\n", errorInfo.GetError_id())
	log.Printf("ErrorMsg %v\n", errorInfo.GetError_msg())
}

//OnSubMarketData 订阅行情应答
func (p *GoXTPQuoteSpi) OnSubMarketData(ticker goxtp.XTPST, errorInfo goxtp.XTPRI, isLast bool) {
}

//OnUnSubMarketData 退订行情应答
func (p *GoXTPQuoteSpi) OnUnSubMarketData(ticker goxtp.XTPST, errorInfo goxtp.XTPRI, isLast bool) {
}

//OnDepthMarketData 深度行情通知，包含买一卖一队列
func (p *GoXTPQuoteSpi) OnDepthMarketData(marketData goxtp.XTPMD, bid1Qty []int64, maxBid1Count int, ask1Qty []int64, maxAsk1Count int) {

}

//OnSubOrderBook 订阅行情订单簿应答
func (p *GoXTPQuoteSpi) OnSubOrderBook(ticker goxtp.XTPST, errorInfo goxtp.XTPRI, isLast bool) {
}

//OnUnSubOrderBook 退订行情订单簿应答
func (p *GoXTPQuoteSpi) OnUnSubOrderBook(ticker goxtp.XTPST, errorInfo goxtp.XTPRI, isLast bool) {
}

//OnOrderBook 退订逐笔行情应答
func (p *GoXTPQuoteSpi) OnOrderBook(orderBook goxtp.XTPOB) {

}

//OnSubTickByTick  逐笔行情通知
func (p *GoXTPQuoteSpi) OnSubTickByTick(ticker goxtp.XTPST, errorInfo goxtp.XTPRI, isLast bool) {
}

//OnUnSubTickByTick 退订逐笔行情应答
func (p *GoXTPQuoteSpi) OnUnSubTickByTick(ticker goxtp.XTPST, errorInfo goxtp.XTPRI, isLast bool) {

}

//OnTickByTick 逐笔行情通知
func (p *GoXTPQuoteSpi) OnTickByTick(tbtData goxtp.XTPTBT) {}

//OnSubscribeAllMarketData 订阅全市场的行情应答
func (p *GoXTPQuoteSpi) OnSubscribeAllMarketData(errorInfo goxtp.XTPRI) {}

//OnUnSubscribeAllMarketData 退订全市场的行情应答
func (p *GoXTPQuoteSpi) OnUnSubscribeAllMarketData(errorInfo goxtp.XTPRI) {}

//OnSubscribeAllOrderBook 订阅全市场的行情订单簿应答
func (p *GoXTPQuoteSpi) OnSubscribeAllOrderBook(errorInfo goxtp.XTPRI) {}

//OnUnSubscribeAllOrderBook 退订全市场的行情订单簿应答
func (p *GoXTPQuoteSpi) OnUnSubscribeAllOrderBook(errorInfo goxtp.XTPRI) {}

//OnSubscribeAllTickByTick 订阅全市场的逐笔行情应答
func (p *GoXTPQuoteSpi) OnSubscribeAllTickByTick(errorInfo goxtp.XTPRI) {}

//OnUnSubscribeAllTickByTick 退订全市场的逐笔行情应答
func (p *GoXTPQuoteSpi) OnUnSubscribeAllTickByTick(errorInfo goxtp.XTPRI) {}

//OnQueryAllTickers 查询可交易合约的应答
func (p *GoXTPQuoteSpi) OnQueryAllTickers(tickerInfo goxtp.XTPQSI, errorInfo goxtp.XTPRI, isLast bool) {
}

//OnQueryTickersPriceInfo 查询合约的最新价格信息应答
func (p *GoXTPQuoteSpi) OnQueryTickersPriceInfo(tickerInfo goxtp.XTPTPI, errorInfo goxtp.XTPRI, isLast bool) {
}

func main() {
	fmt.Printf("xtp-go config: %+v\n", config)

	xtp := GoXTPClient{
		Host:       config.Host,
		Port:       config.Port,
		Username:   config.Username,
		Password:   config.Password,
		AccountKey: config.AccountKey,

		ClientID: 1,
	}
	quoteAPI := goxtp.QuoteApiCreateQuoteApi(uint8(xtp.ClientID), config.LogPath)
	if quoteAPI.Swigcptr() == 0 {
		fmt.Println("here!")

	}

	pQuoteSPI := goxtp.NewDirectorQuoteSpi(&GoXTPQuoteSpi{Client: xtp})
	log.Println("here!")
	fmt.Println("here!!!!!")
	quoteAPI.SetHeartBeatInterval(15)
	quoteAPI.SetUDPBufferSize(128)
	quoteAPI.RegisterSpi(pQuoteSPI)
	log.Println("create QuoteApi success")
	ret := quoteAPI.Login(xtp.Host, xtp.Port, xtp.Username, xtp.Password, goxtp.XTP_PROTOCOL_TCP)
	log.Printf("login return: %v", ret)

	// select {}

}
