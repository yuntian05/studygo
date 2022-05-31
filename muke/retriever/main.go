package main

import (
	"fmt"
	"studygo/muke/retriever/mock"
	real2 "studygo/muke/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.mooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "this is another fake mooc",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is fake retriever"}
	r = &retriever
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	fmt.Println("try a session")
	fmt.Println(session(&retriever))
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v \n", r, r)
	fmt.Print(" > type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:" + v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:" + v.UserAgent)
	}
}
