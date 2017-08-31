package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Pers struct {
	Name     string
	Age      int
	Children string
}

func (p *Pers) setName(n string) {
	p.Name = n
}

func main() {
	fmt.Print("hello")
	if err := structure(); err != nil {
		fmt.Println(err)
	}
}

func structure() error {
	fmt.Print("hi")
	j, err := ioutil.ReadFile("sample-yaml")
	if err != nil {
		return err
	}
	//fmt.Println(jsn)

	//j := []byte(jsn)

	var p Pers
	p.setName("bummer")
	fmt.Printf("%v", p.Name)

	if err := json.Unmarshal(j, &p); err != nil {
		return err
	}

	fmt.Printf("%v", p.Name)
	return nil
}
