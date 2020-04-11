module github.com/yanlong-li/HelloWorldClient

go 1.14

require (
	github.com/ProtonMail/gopenpgp/v2 v2.0.0
	github.com/yanlong-li/HelloWorld-GO v0.0.0-20200308072225-efccc577e911
	github.com/yanlong-li/HelloWorldServer v0.0.0-20200308072454-e28492f0b3c9
)
replace (
    golang.org/x/crypto => github.com/ProtonMail/crypto v0.0.0-20191122234321-e77a1f03baa0
    github.com/yanlong-li/HelloWorld-GO => ../HelloWorld-GO
    github.com/yanlong-li/HelloWorldServer => ../HelloWorldServer
)
