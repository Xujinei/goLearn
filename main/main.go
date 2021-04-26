package main

import (
	"fmt"
	"github.com/Xujinei/goLearn/service"
	"github.com/objcoding/testmod"
)

func main() {
	fmt.Println(testmod.Hi("nn"))
	service.SayHi("name")
}
