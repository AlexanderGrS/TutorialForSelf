package main

import (
	"localImports/pack1"
	"localImports/pack1/subpack1"
	"localImports/pack2"
	"localImports/pack3"
)

func main() {
	pack1.SayHello()
	subpack1.SayCustomHello()
	pack2.SayGoodBye()
	pack3.SayHello()
}
