package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"myName"`
	Sex  string `info:"sex"`
}

type Movie struct {
	Name   string   `json:"name"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	var re resume
	finfTag(&re)
	fmt.Println("========struct to json=======")
	moves := Movie{"xijuzhiwang", 2000,
		38, []string{"zhouxingchi", "zhangbaozhi"}}
	jsonStr, err := json.Marshal(moves)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr:%s", jsonStr)

	fmt.Println("=======json to struct========")
	jsonString := "{\"name\":\"xijuzhiwang\",\"year\":2000,\"rmb\":38,\"actors\":[\"zhouxingchi\",\"zhangbaozhi\"]}"
	myMovices := Movie{}
	erros := json.Unmarshal([]byte(jsonString), &myMovices)
	if erros != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovices)

}

func finfTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println("info:", tagString)
	}
}
