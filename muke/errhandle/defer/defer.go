package main

import (
	"bufio"
	"fmt"
	"os"
	"studygo/muke/functional/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occured")
	fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL | os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibnocci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main() {
	dir, _ := os.Getwd()
	writeFile(dir + "/muke/errhandle/defer/fib.txt")
}
