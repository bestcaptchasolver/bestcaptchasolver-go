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

	// submit task
	parameters := map[string]string{
		"template_name": "Login test page",
		"page_url":      "https://bestcaptchasolver.com/automation/login",
		"variables":     "{\"username\": \"xyz\", \"password\": \"0000\"}",
	}
	//parameters["user_agent"] = "user agent for solving captcha"
	//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
	//parameters["affiliate_id"] = "affiliate_id from /account webpage"
	captchaId, err := api.Submit("task", parameters)
	if err != nil {
		fmt.Printf("ERROR submit: %s\n", err.Error())
		return
	}

	// push variables while task is being solved - i.e 2fa code
	//pushVariables := "{\"2fa_code\": \"35503\"}"
	//err = api.TaskPushVariables(captchaId, pushVariables)

	// wait for captcha to be solved
	fmt.Printf("Waiting for captcha #%d to be solved ...\n", captchaId)
	result, err := api.Solve(captchaId, 10)
	if err != nil {
		fmt.Printf("ERROR solve: %s\n", err.Error())
		return
	}
	fmt.Printf("Solution: %s\n", result["solution"])

	// if task was solved incorrectly, set it as bad
	//err = api.SetCaptchaBad(captchaId)
}
