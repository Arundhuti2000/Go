package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var data [][]byte
	var err error
	for index, item:= range items{
		data[index], err= json.Marshal(item)
	}
	var myZero [][]byte
	if err != nil{
		return myZero,err
	}
	return data,nil
}
