package code_gen

import (
	"net"
	"testing"
	"time"
)

func requireLocalCodegenMysql(t *testing.T) {
	t.Helper()

	conn, err := net.DialTimeout("tcp", "127.0.0.1:3309", time.Second)
	if err != nil {
		t.Skipf("skip code_gen database test: %v", err)
		return
	}

	_ = conn.Close()
}

func TestGenStruct(t *testing.T) {
	requireLocalCodegenMysql(t)
	GenModel("member", "Member")
}
func TestGenRpc(t *testing.T) {
	rpcCommon := RpcCommon{
		PackageName: "mclient",
		ModuleName:  "Market",
		ServiceName: "Market",
		GrpcPackage: "market",
	}
	// rpc FindSymbolThumb(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolThumbTrend(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolInfo(MarketReq) returns(ExchangeCoin);
	rpc1 := Rpc{
		FunName: "FindSymbolThumbTrend",
		Resp:    "SymbolThumbRes",
		Req:     "MarketReq",
	}
	//rpc2 := Rpc{
	//	FunName: "FindSymbolThumbTrend",
	//	Resp:    "SymbolThumbRes",
	//	Req:     "MarketReq",
	//}
	//rpc3 := Rpc{
	//	FunName: "FindSymbolInfo",
	//	Resp:    "ExchangeCoin",
	//	Req:     "MarketReq",
	//}
	rpcList := []Rpc{}
	rpcList = append(rpcList, rpc1)
	result := RpcResult{
		RpcCommon: rpcCommon,
		Rpc:       rpcList,
	}
	GenZeroRpc(result)
}
