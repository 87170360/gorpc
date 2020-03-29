package common

var ServerAddr = "127.0.0.1:10011"

type Args struct {
	A int
	B int
}

type Reply struct {
	A *int32
	B *int32
	C *int
	D *Args
}
