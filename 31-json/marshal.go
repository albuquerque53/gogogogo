package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	rex := dog{"Rex", 3}

	jsonRex := marshalData(rex)

	fmt.Printf("This is my dog (json): %v\n", jsonRex)

	var unmarshalRex dog

	unmarshalData(jsonRex, &unmarshalRex)

	fmt.Printf("This is my dog (struct): %#v", unmarshalRex)
}

func marshalData(data dog) string {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
	// return bytes.NewBuffer(jsonData)
}

func unmarshalData(jsonData string, output *dog) {
	if err := json.Unmarshal([]byte(jsonData), &output); err != nil {
		log.Fatal(err)
	}
}
