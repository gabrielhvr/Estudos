package main

import (
		"fmt"
)
func TratamentoErro(err error){
	if err != nil {
		fmt.Println("Error: ", err)
	}
}