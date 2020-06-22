package main

import (
	"fmt"
	"log"
)

func main()  {
	c,err:=UnmarshalConfig("./test.toml")
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(c.GetTestkey())
	fmt.Println(c.GetTestArray()[0])
	fmt.Println(c.getA("b").A1)
	fmt.Println(c.GetPostgresql().PassWord)
}
