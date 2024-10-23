package utils

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
)

func ValidateEmail(email string) error {
	var (
		verifier = emailverifier.NewVerifier()
	)

	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return err
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return fmt.Errorf("invalid email address syntax")
	}

	fmt.Println("email validation result", ret)
	return nil
}

