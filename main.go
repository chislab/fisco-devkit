package main

import (
	"fisco/check"
	"flag"
)

func main() {
	module := flag.String("module", "test", "选择测试的模块, 默认为test")

	flag.Parse()

	if *module == "test" {
		check.Test()
	}
}
