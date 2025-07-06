package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type smsOtp struct {
}

func (sms *smsOtp) genRandomOTP(otpLength int) string {

	minOtp := 1

	for i := 1; i < otpLength; i++ {
		minOtp = minOtp * 10
	}
	maxOtp := minOtp*10 - 1

	randomNumber := rand.Intn(maxOtp-minOtp+1) + minOtp
	fmt.Println("Generated OTP ", randomNumber)
	return strconv.Itoa(randomNumber)

}

func (sms *smsOtp) saveOTPToCache(otpValue string) {
	fmt.Println("OTP" + otpValue + " saved to sms cache")
}

func (sms *smsOtp) getMessage(otpValue string) string {
	return fmt.Sprintf("Please use %s otp to login into the app", otpValue)
}

func (sms *smsOtp) sendNotification(message string) error {
	fmt.Println("Sent OTP Notification Message through SMS: ", message)
	return nil
}
