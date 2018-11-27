package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()

	a := [3]int{1, 2, 3}
	fmt.Printf("I am array[数组] .%d--%T---%v---%#x", a, a, a, a)

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("\r\nI am slice[切片] .%d--%T---%v---%#x", s, s, s, s)

	m := map[string]string{"name": "123", "ages": "18"}

	fmt.Printf("\r\nI am map [映射] .-%d--%T---%v---%#x", m, m, m, m)

	if help {
		usage()
	}
}

var (
	flagVal            *string
	help               bool
	user, home, gopath string
)

func init() {
	flag.StringVar(&user, "user", "user", "help message for flagname")
	flag.StringVar(&home, "home", "home", "home message for flagname")
	flag.BoolVar(&help, "help", false, "this help")
	flag.Bool("person", false, "suppress non-error messages during configuration testing")
	flag.Usage = usage

	if user == "" {
		log.Fatalln("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	fmt.Println(flagVal)
	// gopath 可通过命令行中的 --gopath 标记覆盖掉。
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
