package main

type iOTP interface {
	genRandomOTP(int) string
	saveOTPToCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type otp struct {
	iOTP iOTP
}

func (o *otp) genAndSendOTP(otpLength int) error {
	//generate a random otp
	otpValue := o.iOTP.genRandomOTP(otpLength)
	//save otp to cache
	o.iOTP.saveOTPToCache(otpValue)
	//get the message that needs to be sent to user
	message := o.iOTP.getMessage(otpValue)
	//send notification to user
	err := o.iOTP.sendNotification(message)

	if err != nil {
		return err
	}
	return nil
}
