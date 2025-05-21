package helpers

import "golang.org/x/exp/rand"

func GenerateRandomOTP() string {
	const otpChars = "1234567890"
	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(otp)
}
