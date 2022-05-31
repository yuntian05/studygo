package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	Id         string      `json:"id"`
	Items      []*OrderItem `json:"item"`
	Quality    int         `json:"quality"`
	TotalPrice float64     `json:"total_price"`
}

func main() {
	unmarshal()
}

func marshal()  {
	o := Order{
		Id:         "1234",
		Quality:    3,
		TotalPrice: 20,
		Items: []*OrderItem{
			{
				Id:    "item_1",
				Name:  "learn go",
				Price: 20,
			},
			{
				Id:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
	}
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func unmarshal()  {
	s := `{"id":"1234","item":[{"id":"item_1","name":"learn go","price":20},{"id":"item_2","name":"interview","price":10}],"quality":3,"total_price":20}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}
