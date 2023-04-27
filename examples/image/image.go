package main

import (
	"encoding/base64"
	"fmt"
	"github.com/bestcaptchasolver/bestcaptchasolver-go"
	"io/ioutil"
	"log"
	"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func readImageB64(filePath string) string {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	return base64Encoding
}

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
	b64Image := readImageB64("captcha.jpg")
	parameters := map[string]string{"b64image": b64Image}
	//parameters["is_case"] = "1"
	//parameters["is_phrase"] = "1"
	//parameters["is_math"] = "1"
	//parameters["alphanumeric"] = "1"
	//parameters["minlength"] = "3"
	//parameters["maxlength"] = "6"
	//parameters["affiliate_id"] = "your affiliate ID"
	captchaId, err := api.Submit("image", parameters)
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
	fmt.Printf("Captcha text: %s\n", result["text"])

	// if captcha was solved incorrectly, set it as bad
	//err = api.SetCaptchaBad(captchaId)
}
