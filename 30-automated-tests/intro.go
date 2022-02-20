package main

import (
	"automated-tests/address"
	"fmt"
)

func main() {
	addressType := address.GetAddressType("Alameda México")

	fmt.Println(addressType)
}
