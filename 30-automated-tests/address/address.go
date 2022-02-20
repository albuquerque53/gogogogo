package address

import "strings"

// GetAddressType must verify if especific address is valid
func GetAddressType(address string) string {
	validTypes := []string{"rua", "avenida", "alameda"}

	lowerAddress := strings.ToLower(address)
	firstWordOfAddress := strings.Split(lowerAddress, " ")[0]

	isAValidAddressType := false

	for _, addressType := range validTypes {
		if validAddressType(addressType, firstWordOfAddress) {
			isAValidAddressType = true
		}
	}

	if isAValidAddressType {
		return strings.Title(firstWordOfAddress)
	}

	return "Invalid address type!"
}

func validAddressType(expectedAddressType, realAddressType string) bool {
	return expectedAddressType == realAddressType
}
