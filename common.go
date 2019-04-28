package common

const (
	//1000~2000 base
	RpcSuccessful          int32 = 1000
	RpcFailed              int32 = 1001
	ServerInternalError	   int32 = 1002

	//0~1000 judge
	AC int32 = 0
	EA int32 = 1
	TL int32 = 2
	ML int32 = 3
	CE int32 = 4
)


var Message = map[int32]string{
	AC:                  "通过",
	EA:                  "错误的解答",
	TL:                  "超时",
	ML:                  "内存溢出",
	CE:                  "编译错误",
	RpcFailed:           "Rpc Failed",
	ServerInternalError: "ServerInternalError",
}


