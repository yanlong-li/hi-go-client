module github.com/yanlong-li/hi-go-client

go 1.15

require (
	github.com/yanlong-li/hi-go-logger v0.0.0-20201019104050-b1e94d395fee
	github.com/yanlong-li/hi-go-server v0.0.0-20201021022935-8e3cc6f1fd77
	github.com/yanlong-li/hi-go-gateway v0.0.0-20201119075128-0a84a9b658ce
	github.com/yanlong-li/hi-go-socket v0.0.0-20201118092047-1f8816f4990c
)

replace (
	github.com/yanlong-li/hi-go-logger => ../hi-go-logger
	github.com/yanlong-li/hi-go-server => ../hi-go-server
	github.com/yanlong-li/hi-go-socket => ../hi-go-socket
	github.com/yanlong-li/hi-go-gateway => ../hi-go-gateway
)
