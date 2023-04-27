BestCaptchaSolver.com go API wrapper
=========================================

bestcaptchasolver-go is a super easy to use bypass captcha go API wrapper for bestcaptchasolver.com captcha service

## Installation    
```bash
go get github.com/bestcaptchasolver/bestcaptchasolver-go
```

## Usage
```go
import (
	"github.com/bestcaptchasolver/bestcaptchasolver-go"
)
```

## How to use?

Simply import the module, initialize a new API variable using your access token then start using it

``` go
api := bestcaptchasolverapi.New("ACCESS_TOKEN_HERE")
```
**Get balance**

``` go
balance, err := api.GetBalance()            
if err != nil {
    fmt.Printf("ERROR balance: %s\n", err.Error())
    return
}
fmt.Printf("Balance: $%f\n", balance)
```

**Submit image captcha**

``` go
// read image as bas64 encoded string
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
```
Read your image as a base64 encoded string and use that to submit it to our API
For setting the affiliate_id, set the `affiliate_id` parameter

**Submit reCAPTCHA details**

For recaptcha submission there are two things that are required.
- page_url
- site_key
- type (optional, defaults to 1 if not given)
  - `1` - v2
  - `2` - invisible
  - `3` - v3
  - `4` - enterprise v2
  - `5` - enterprise v3
- v3_action (optional)
- v3_min_score (optional)
- domain (optional) - i.e `www.google.com` or `recaptcha.net`
- data_s (optional)
- cookie_input (optional)
- user_agent (optional)
- affiliate_id (optional)
- proxy (optional)

``` go
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
```

This method returns a captchaID. This ID will be used next, to retrieve the g-response, once workers have 
completed the captcha. This takes somewhere between 10-80 seconds.

**Geetest**
- domain
- gt
- challenge
- api_server (optional)
- user_agent (optional)
- proxy (optional)

```go
parameters := map[string]string{
    "domain":    "www.example.com",
    "gt":        "a84b2f014c197bafc401985ab3459c14",
    "challenge": "61b6eb54f3841eb91ae806b6ead3337k",
}
//parameters["api_server"] = "GT domain"
//parameters["user_agent"] = "user agent for solving captcha"
//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
//parameters["affiliate_id"] = "affiliate_id from /account webpage"
captchaId, err := api.Submit("geetest", parameters)
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
fmt.Printf("Solution: %s\n", result["solution"])
```

**GeetestV4**

- domain
- captchaid
- user_agent (optional)
- proxy (optional)

**Important:** This is not the captchaid that's in our system that you receive while submitting a captcha. Gather this from HTML source of page with geetestv4 captcha, inside the `<script>` tag you'll find a link that looks like this: https://i.imgur.com/XcZd47y.png

```go
parameters := map[string]string{
    "domain":    "www.example.com",
    "captchaid": "747b5ae3ed8aab4be64784e01556bb72",
}
//parameters["user_agent"] = "user agent for solving captcha"
//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
//parameters["affiliate_id"] = "affiliate_id from /account webpage"
captchaId, err := api.Submit("geetestv4", parameters)
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
fmt.Printf("Solution: %s\n", result["solution"])
```

**Capy**
- page_url
- site_key
- user_agent (optional)
- proxy (optional)

```go
// submit captcha
parameters := map[string]string{
    "page_url": "www.example.com",
    "site_key": "Cde6hPLkiZRMYC3uh416VD2U3mNs6v",
}
//parameters["user_agent"] = "user agent for solving captcha"
//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
//parameters["affiliate_id"] = "affiliate_id from /account webpage"
captchaId, err := api.Submit("capy", parameters)
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
fmt.Printf("Solution: %s\n", result["solution"])
```

**hCaptcha**
- page_url
- site_key
- invisible (optional)
- payload (optional)
- domain (optional)
- user_agent (optional)
- proxy (optional)

```go
parameters := map[string]string{
    "page_url": "www.example.com",
    "site_key": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
}
//parameters["invisible"] = "1"
//parameters["payload"] = "{\"rqdata\": \"get from page source / DOM\"}"
//parameters["domain"] = "hcaptcha.com"
//parameters["user_agent"] = "user agent for solving captcha"
//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
//parameters["affiliate_id"] = "affiliate_id from /account webpage"
captchaId, err := api.Submit("hcaptcha", parameters)
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
fmt.Printf("Solution: %s\n", result["solution"])
```

**FunCaptcha (Arkose Labs)**
- page_url
- s_url
- site_key
- data (optional)
- user_agent (optional)
- proxy (optional)

```go
parameters := map[string]string{
    "page_url": "www.example.com",
    "site_key": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "s_url":    "https://api.arkoselabs.com",
}
//parameters["data"] = "{\"x\":\"y\"}"
//parameters["user_agent"] = "user agent for solving captcha"
//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
//parameters["affiliate_id"] = "affiliate_id from /account webpage"
captchaId, err := api.Submit("funcaptcha", parameters)
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
fmt.Printf("Solution: %s\n", result["solution"])
```

**Turnstile (Cloudflare)**
- page_url
- site_key
- action (optional)
- cdata (optional)
- domain (optional)
- user_agent (optional)
- proxy (optional)

```go
parameters := map[string]string{
    "page_url": "www.example.com",
    "site_key": "0x4AAAAAAABfrvQ6Lbx22GS",
}
//parameters["action"] = "taken from page source / DOM"
//parameters["cdata"] = "taken from page source / DOM"
//parameters["domain"] = "challanges.cloudflare.com"
//parameters["user_agent"] = "user agent for solving captcha"
//parameters["proxy"] = "123.45.67.89:3031 or user:pass@123.45.67.89:3031"
//parameters["affiliate_id"] = "affiliate_id from /account webpage"
captchaId, err := api.Submit("turnstile", parameters)
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
fmt.Printf("Solution: %s\n", result["solution"])
```

**Task**
- template_name
- page_url
- variables

```go
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
```

#### Task pushVariables
Update task variables while it is being solved by the worker. Useful when dealing with data / variables, of which
value you don't know, only after a certain step or action of the task. For example, in websites that require 2 factor
authentication code.

When the task (while running on workers machine) is getting to an action defined in the template, that requires a variable, but variable was not
set with the task submission, it will wait until the variable is updated through push.

The `TaskPushVariables(captchaId, pushVariables)` method can be used as many times as it is needed.

```go
push variables while task is being solved - i.e 2fa code
pushVariables := "{\"2fa_code\": \"35503\"}"
err = api.TaskPushVariables(captchaId, pushVariables)
```

**Set captcha bad**

When a captcha was solved wrong by our workers, you can notify the server with it's ID,
so we know something went wrong.

``` go
err = api.SetCaptchaBad(captchaId)
```

## Examples
Check `examples` folder

## License
API library is licensed under the MIT License

## More information
More details about the server-side API can be found [here](https://bestcaptchasolver.com/api )


<sup><sub>captcha, bypasscaptcha, decaptcher, decaptcha, 2captcha, deathbycaptcha, anticaptcha, 
bypassrecaptchav2, bypassnocaptcharecaptcha, bypassinvisiblerecaptcha, captchaservicesforrecaptchav2, 
recaptchav2captchasolver, googlerecaptchasolver, recaptchasolverpython, recaptchabypassscript, bestcaptchasolver</sup></sub>

