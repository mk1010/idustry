package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mk1010/idustry/modules/industry_identification_center/model"
)

var jsonString = `
{
	"ID":10
}
`

func TestModule(t *testing.T) {
	f, err := os.OpenFile("./server.crt", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	fileByte, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(fileByte), string(fileByte))
}

func TestJson(t *testing.T) {
	s := model.Student{}
	fmt.Println(json.Unmarshal([]byte(jsonString), &s))
	q, _ := json.Marshal(&s)
	v := string(q)
	t.Log(v)
}
