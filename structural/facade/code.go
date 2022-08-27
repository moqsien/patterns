package facade

import "fmt"

type SecurityCode struct {
	code int
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}
