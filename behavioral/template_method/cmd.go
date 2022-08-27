package template

import "fmt"

/*
模版方法是一种行为设计模式，它在基类中定义了一个算法的框架，
允许子类在不修改结构的情况下重写算法的特定步骤。

考虑一个一次性密码功能（OTP）的例子。
将 OTP 传递给用户的方式多种多样 （短信、 邮件等）。
*/

func execute() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}
