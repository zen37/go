package main

import (
	"flag"
	"fmt"
)

func main() {
	var str *string
	str = flag.String("str", "foo", "the flag is a string")

	var num *int
	num = flag.Int("num", 1, "the flag is an integer")

	var err_show *bool
	err_show = flag.Bool("err", false, "this is a boolean for printing or not errors. default is false")

	flag.Parse()

	fmt.Println(*str, *num, *err_show)

	if *err_show {
		fmt.Println("errors are printed")
	} else {
		fmt.Println("errors are not printed")
	}
}
