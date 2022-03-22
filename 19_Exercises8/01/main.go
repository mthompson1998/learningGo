package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct{
	Name string
	Age int
}

func main() {
	u1 := user{
		Name: "James",
		Age: 33,
	}
	u2 := user{
		Name: "Moneypenny",
		Age: 27,
	}
	u3 := user{
		Name: "Dr. Evil",
		Age: 35,
	}
users := []user{u1, u2, u3}
fmt.Println(users)

b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
os.Stdout.Write(b)
}