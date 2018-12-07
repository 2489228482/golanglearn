package main

import (
	"fmt"
	"io/ioutil"
	"github.com/go-xmlpath/xmlpath"
)

func readFile(p string) string{
	f,err := ioutil.ReadFile(p)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	return string(f)
}

func 

func main() {
	path :="./file/baiduweb.html"
	html := readFile(path)
	fmt.Println(html)

}