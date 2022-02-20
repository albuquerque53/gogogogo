package address_test

import (
	. "automated-tests/address"
	"testing"
)

type testScenery struct {
	addressToInsert string
	expectedAddress string
}

func TestGetAddressType(test *testing.T) {
	testSceneries := []testScenery{
		{"Rua dos Bobos", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Alameda México", "Alameda"},
		{"Praça das Rosas", "Invalid address type!"},
		{"RUA MILÃO", "Rua"},
		{"AVENIDA PRIMAVERA DE CAIENA", "Avenida"},
		{"ALAMEDA POR-DO-SOL", "Alameda"},
		{"Estrada Santos", "Invalid address type!"},
		{"", "Invalid address type!"},
	}

	for _, sceneryToTest := range testSceneries {
		expectedAddressType := sceneryToTest.expectedAddress
		receivedAddressType := GetAddressType(sceneryToTest.addressToInsert)

		if expectedAddressType != receivedAddressType {
			test.Errorf("The received type (%s) is not equal to expected (%s)",
				receivedAddressType,
				expectedAddressType,
			)
		}
	}

}
