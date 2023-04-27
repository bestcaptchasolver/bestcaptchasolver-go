package main

import (
	"fmt"
	"github.com/bestcaptchasolver/bestcaptchasolver-go"
)

func main() {
	api := bestcaptchasolverapi.New("ACCESS_TOKEN_HERE")
	// get the user balance
	balance, err := api.GetBalance()
	if err != nil {
		fmt.Printf("ERROR balance: %s\n", err.Error())
		return
	}
	fmt.Printf("Balance: $%f\n", balance)

	// submit captcha
	parameters := map[string]string{
		"page_url": "www.example.com",
		"site_key": "6LfKLmcUAAAAALGIbt_FxCOLMX_GlwMaKAfbbNCU",
	}
	//reCAPTCHA type(s) - optional, defaults to 1
	//---------------------------------------------
	//1 - v2
	//2 - invisible
	//3 - v3
	//4 - enterprise v2
	//5 - enterprise v3
	//parameters["type"] = "1"
	//parameters["v3_action"] = "v3 recaptcha action"
	//parameters["v3_min_score"] = "0.3"
	//parameters["domain"] = "www.google.com"
	//parameters["data_s"] = "recaptcha data-s parameter used in loading reCAPTCHA"
	//parameters["cookie_input"] = "a=b;c=d"
	//parameters["user_agent"] = "user agent for solving captcha"
	//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
	//parameters["affiliate_id"] = "affiliate_id from /account webpage"
	captchaId, err := api.Submit("recaptcha", parameters)
	if err != nil {
		fmt.Printf("ERROR submit: %s\n", err.Error())
		return
	}

	// wait for captcha to be solved
	fmt.Printf("Waiting for captcha #%d to be solved ...\n", captchaId)
	result, err := api.Solve(captchaId, 10)
	if err != nil {
		fmt.Printf("ERROR solve: %s\n", err.Error())
		return
	}
	fmt.Printf("Solution: %s\n", result["gresponse"])

	// if captcha was solved incorrectly, set it as bad
	//err = api.SetCaptchaBad(captchaId)
}
