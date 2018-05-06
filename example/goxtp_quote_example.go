package main

import (
	"flag"
	"fmt"
	"goxtp"
	"log"
)

var (
	quoteServerHost = flag.String("host", "120.27.164.138", "行情服务器地址")
	quoteServerPort = flag.Int("port", 6002, "行情服务器端口")
	quoteUsername   = flag.String("username", "pseudo", "行情服务登录用户名")
	quotePassword   = flag.String("password", "******", "行情服务登录密码")
	accountKey      = flag.String("accountkey", "xxxxxxxx", "api key")
	logPath         = flag.String("logpath", "./test", "相关日志输出路径")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	flag.Parse()
}

type GoXTPClient struct {
	Host string
	Port int

	Username string
	Password string

	ClientId int
	Path     string
}

func NewDirectorQuoteSpi(v interface{}) goxtp.QuoteSpi {
	return goxtp.NewDirectorQuoteSpi(v)
}

type GoXTPQuoteSpi struct {
	Client GoXTPClient
}

func (p *GoXTPQuoteSpi) OnDisconnected(arg2 int) {
	log.Printf("OnDisconnected: %v\n", arg2)
}
func (p *GoXTPQuoteSpi) OnError(arg2 goxtp.XTPRI) {
	log.Printf("ErrorID %v\n", arg2.GetError_id())
	log.Printf("ErrorMsg %v\n", arg2.GetError_msg())
}

func (p *GoXTPQuoteSpi) OnSubMarketData(arg2 goxtp.XTPST, arg3 goxtp.XTPRI, arg4 bool) {
}
func (p *GoXTPQuoteSpi) OnUnSubMarketData(arg2 goxtp.XTPST, arg3 goxtp.XTPRI, arg4 bool) {
}
func (p *GoXTPQuoteSpi) OnDepthMarketData(arg2 goxtp.XTPMD, arg3 []int64, arg4 int, arg5 []int64, arg6 int) {

}
func (p *GoXTPQuoteSpi) OnSubOrderBook(arg2 goxtp.XTPST, arg3 goxtp.XTPRI, arg4 bool) {
}
func (p *GoXTPQuoteSpi) OnUnSubOrderBook(arg2 goxtp.XTPST, arg3 goxtp.XTPRI, arg4 bool) {}
func (p *GoXTPQuoteSpi) OnOrderBook(arg2 goxtp.XTPOB)                                   {}

func (p *GoXTPQuoteSpi) OnSubTickByTick(arg2 goxtp.XTPST, arg3 goxtp.XTPRI, arg4 bool) {}

func (p *GoXTPQuoteSpi) OnUnSubTickByTick(arg2 goxtp.XTPST, arg3 goxtp.XTPRI, arg4 bool) {}

func (p *GoXTPQuoteSpi) OnTickByTick(arg2 goxtp.XTPTBT) {}

func (p *GoXTPQuoteSpi) OnSubscribeAllMarketData(arg2 goxtp.XTPRI) {}

func (p *GoXTPQuoteSpi) OnUnSubscribeAllMarketData(arg2 goxtp.XTPRI) {}

func (p *GoXTPQuoteSpi) OnSubscribeAllOrderBook(arg2 goxtp.XTPRI) {}

func (p *GoXTPQuoteSpi) OnUnSubscribeAllOrderBook(arg2 goxtp.XTPRI) {}

func (p *GoXTPQuoteSpi) OnSubscribeAllTickByTick(arg2 goxtp.XTPRI) {}

func (p *GoXTPQuoteSpi) OnUnSubscribeAllTickByTick(arg2 goxtp.XTPRI) {}

func (p *GoXTPQuoteSpi) OnQueryAllTickers(arg2 goxtp.XTPQSI, arg3 goxtp.XTPRI, arg4 bool) {}

func (p *GoXTPQuoteSpi) OnQueryTickersPriceInfo(arg2 goxtp.XTPTPI, arg3 goxtp.XTPRI, arg4 bool) {}

func main() {
	fmt.Println("vim-go")

	xtp := GoXTPClient{
		Host:     *quoteServerHost,
		Port:     *quoteServerPort,
		Username: *quoteUsername,
		Password: *quotePassword,

		ClientId: 1,
	}

	quoteApi := goxtp.QuoteApiCreateQuoteApi(xtp.ClientId, *logPath)
	pQuoteSpi := goxtp.NewDirectorQuoteSpi(&GoXTPQuoteSpi{Client: xtp})
	quoteApi.SetHeartBeatInterval(15)
	quoteApi.SetUDPBufferSize(128)
	quoteApi.RegisterSpi(pQuoteSpi)

	quoteApi.Login(xtp.Host, xtp.Port, xtp.Username, xtp.Password, 1)

	select {}

}
