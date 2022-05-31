package main

import (
	"fmt"
	"regexp"
)

const text = `my email is yuntian0105@gmail.com
email1 is xx@qq.com
email2 is ss@outlook.com
email3 is jj@dd.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}

