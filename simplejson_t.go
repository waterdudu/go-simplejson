package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"reflect"
)

func simplejson_t() {
	js, err := simplejson.NewJson([]byte(`{
    "test": {
        "array": [1, "2", 3, 9, 1.23],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
        }
    }`))

	if err != nil {
		fmt.Println("error decode json using simplejson")
		return
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Printf("arr : %v, i : %v, ms : %v\n", arr, i, ms)

	for _, a := range arr {
		t := reflect.TypeOf(a)
		fmt.Printf("kind : %v, value : %v\n", t.Kind(), a)
	}
}

func main() {
	simplejson_t()
}
