package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}
// $ go run main.go -name=eddycjy -n=煎鱼
//name: 煎鱼

// 我们可以发现输出的结果是最后一个赋值的变量，也就是 -n。
//
//你可能会有一些疑惑，为什么长短选项要分开两次调用，一个命令行参数的标志位有长短选项，是常规需求，这样子岂不是重复逻辑，有没有优化的办法呢。
//
//实际上标准库 flag 并不直接支持该功能，但是我们可以通过其它第三方库来实现这个功能，这块我们在后续也会使用到。