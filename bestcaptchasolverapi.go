package bestcaptchasolverapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	BASE_URL   = "https://bcsapi.xyz/api"
	USER_AGENT = "goClient"
	TIMEOUT    = 120
)

type Request struct {
	Params map[string]string
	URL    string
}

type Response struct {
	Body map[string]interface{}
}

type Client struct {
	BaseURL     *url.URL
	AccessToken string
	Timeout     int
	httpClient  *http.Client
}

// New - create a new API client
func New(accessToken string) *Client {
	base, _ := url.Parse(BASE_URL)
	return &Client{
		BaseURL:     base,
		AccessToken: accessToken,
		httpClient: &http.Client{
			Timeout: time.Duration(TIMEOUT) * time.Second,
		},
	}
}

// GetRequest - make GET request to API server
func (c *Client) GetRequest(url string) (Response, error) {
	response := Response{}
	// make HTTP request
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", USER_AGENT)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, err
	}

	// read response body
	defer func(Body io.ReadCloser) {
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	bodyString := string(bodyBytes)

	// convert string to map (JSON)
	err = json.Unmarshal([]byte(bodyString), &response.Body)
	status := fmt.Sprint(response.Body["status"])
	// check if response contains error
	if status == "error" {
		return response, errors.New(fmt.Sprint(response.Body["error"]))
	}

	return response, err
}

// PostRequest - make POST request to API server
func (c *Client) PostRequest(url string, parameters map[string]string) (Response, error) {
	response := Response{}
	// make HTTP request
	jsonValue, _ := json.Marshal(parameters)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, err
	}

	// read response body
	defer func(Body io.ReadCloser) {
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	bodyString := string(bodyBytes)

	// convert string to map (JSON)
	err = json.Unmarshal([]byte(bodyString), &response.Body)
	status := fmt.Sprint(response.Body["status"])
	// check if response contains error
	if status == "error" {
		return response, errors.New(fmt.Sprint(response.Body["error"]))
	}

	return response, err
}

// GetBalance - get the user balance
func (c *Client) GetBalance() (float64, error) {
	reqUrl := fmt.Sprintf("%s/user/balance?access_token=%s", c.BaseURL, c.AccessToken)
	response, err := c.GetRequest(reqUrl)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(fmt.Sprint(response.Body["balance"]), 64)
}

// Submit - submit a captcha and get back the id
func (c *Client) Submit(captchaType string, parameters map[string]string) (int64, error) {
	reqUrl := fmt.Sprintf("%s/captcha/%s", c.BaseURL, captchaType)
	// append access token
	parameters["access_token"] = c.AccessToken

	// make request to API server
	response, err := c.PostRequest(reqUrl, parameters)
	if err != nil {
		return 0, err
	}
	// convert the JSON string id to int
	cid := fmt.Sprintf("%f", response.Body["id"])
	x := strings.Split(cid, ".")
	return strconv.ParseInt(x[0], 10, 64)
}

// GetResult - get the captcha solution using the captcha ID returned on submission
func (c *Client) GetResult(captchaId int64) (Response, error) {
	reqUrl := fmt.Sprintf("%s/captcha/%d?access_token=%s", c.BaseURL, captchaId, c.AccessToken)
	response, err := c.GetRequest(reqUrl)
	if err != nil {
		return response, err
	}
	return response, nil
}

// Solve - waits for captcha to be solved and returns solution
func (c *Client) Solve(captchaId int64, pollingInterval int) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	counter := 360 / pollingInterval
	for counter > 0 {
		time.Sleep(time.Duration(pollingInterval) * time.Second)

		// get the result
		res, err := c.GetResult(captchaId)

		if err != nil {
			return result, err
		}
		if res.Body["status"] == "completed" {
			return res.Body, nil
		}
		counter -= 1
	}

	return result, errors.New("captcha could not be solved in time")
}

// TaskPushVariables - push
func (c *Client) TaskPushVariables(captchaId int64, pushVariables string) error {
	reqUrl := fmt.Sprintf("%s/captcha/task/pushVariables/%d", c.BaseURL, captchaId)
	parameters := make(map[string]string)
	// append access token
	parameters["access_token"] = c.AccessToken
	parameters["pushVariables"] = pushVariables

	// make request to API server
	response, err := c.PostRequest(reqUrl, parameters)
	if err != nil {
		return err
	}
	if response.Body["status"] == "error" {
		return errors.New(fmt.Sprint(response.Body["error"]))
	}

	return nil
}

// SetCaptchaBad - sets the captcha as bad captcha
func (c *Client) SetCaptchaBad(captchaId int64) error {
	reqUrl := fmt.Sprintf("%s/captcha/bad/%d", c.BaseURL, captchaId)
	parameters := make(map[string]string)
	// append access token
	parameters["access_token"] = c.AccessToken

	// make request to API server
	response, err := c.PostRequest(reqUrl, parameters)
	if err != nil {
		return err
	}
	if response.Body["status"] == "error" {
		return errors.New(fmt.Sprint(response.Body["error"]))
	}

	return nil
}
