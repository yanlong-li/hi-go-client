module github.com/yanlong-li/hi-go-client

go 1.14

require (
	github.com/yanlong-li/hi-go-logger v0.0.0-20201019104050-b1e94d395fee
	github.com/yanlong-li/hi-go-server v0.0.0-20201019112221-294a5c985b81
	github.com/yanlong-li/hi-go-socket v0.0.0-20201019105643-c29816f01818
)

replace (
	github.com/yanlong-li/hi-go-logger => ../hi-go-logger
	github.com/yanlong-li/hi-go-server => ../hi-go-server
	github.com/yanlong-li/hi-go-socket => ../hi-go-socket
	golang.org/x/crypto => github.com/ProtonMail/crypto v0.0.0-20201016191319-576ad9c42ffa
)
