package main

func main() {

	smsOtp := &smsOtp{}

	otp := &otp{
		iOTP: smsOtp,
	}
	otp.genAndSendOTP(4)

	otp.iOTP = &emailOtp{}

	otp.genAndSendOTP(5)
}
