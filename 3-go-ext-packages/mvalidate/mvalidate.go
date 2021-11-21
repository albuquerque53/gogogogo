package mvalidate

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// Verify Email
func CheckEmail() {
	error := checkmail.ValidateFormat("gabrielalbuquerque-dev@.com")

	fmt.Println(error)
}
