package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unsafe"
)


func main()  {
	printNTHeaders()
}

func readPEFile(fileName string)[]byte {
	//fp, err := os.Open(fileName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer fp.Close()
	//
	//chunks := make([]byte, 0)
	//buf := make([]byte, 1024)
	//for {
	//	n, err := fp.Read(buf)
	//	if err != nil && err != io.EOF {
	//		log.Fatal(err)
	//	}
	//	if 0 == n {
	//		break
	//	}
	//	chunks = append(chunks, buf[:n]...)
	//}
	chunks, err := ioutil.ReadFile(fileName);
	if err != nil {
		log.Fatal(err)
	}
	return chunks
}

type ImageDosHeader struct {
	eMagic    uint16
	eCblp     uint16
	eCp       uint16
	eCrlc     uint16
	eCparhdr  uint16
	eMinalloc uint16
	eMaxalloc uint16
	eSs      uint16
	eSp      uint16
	eCsum    uint16
	eIp      uint16
	eCs      uint16
	eLfarlc  uint16
	eOvno    uint16
	eRes     [4]uint16
	eOemid   uint16
	eOeminfo uint16
	eRes2    [10]uint16
	eLfanew  uint32
}

type ImageNTHeader struct {

}
func printNTHeaders() {
	fileName := "/Users/yuntian/Downloads/1147873423_yuntian0105/computer/安全/32位汇编学习环境/MASM8.0/ml.exe"
	fileData := readPEFile(fileName)
	pDosHeader := (*ImageDosHeader)(unsafe.Pointer(&fileData[0]))
	fmt.Println("========DOC头===========")
	fmt.Printf("MZ标志 %x\n", pDosHeader.eMagic)
}
