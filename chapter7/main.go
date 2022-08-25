package main

import (
	"flag"
	"fmt"
	"time"
)

var endpoint = flag.String(
	"endpoint",
	"hoge.prd.com",
	"The endpoint",
)

var flagbool = flag.Bool(
	"flagbool",
	false,
	"flag bool",
)

var flagint = flag.Int(
	"flagint",
	100,
	"flag int",
)

var flagduration = flag.Duration(
	"flagduration",
	time.Duration(100),
	"flag duration",
)

func main() {
	flag.Parse()

	fmt.Println("endpoint: ", *endpoint)
	fmt.Println("flagbool: ", *flagbool)
	fmt.Println("flagint: ", *flagint)
	fmt.Println("flagduration: ", *flagduration)
}
