package main

import (
	"flag"
	"fmt"
	"net/url"
	"reflect"
	"time"
)

// Custom flag
// flag.Value interface is this:
// type Value interface {
// 	String() string
// 	Set(string) error
// }

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func init() {
	flag.Var(&URLValue{u}, "url", "URL to parse")
}

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

	if reflect.ValueOf(*u).IsZero() {
		panic("did not pass an URL")
	}

	fmt.Printf("{scheme: %q, host: %q, path: %q}\n", u.Scheme, u.Host, u.Path)
	fmt.Println("endpoint: ", *endpoint)
	fmt.Println("flagbool: ", *flagbool)
	fmt.Println("flagint: ", *flagint)
	fmt.Println("flagduration: ", *flagduration)
}
