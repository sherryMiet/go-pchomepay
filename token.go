package go_pchomepay_sdk

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	TokenUrl     = "https://api.pchomepay.com.tw/v1/token"
	TestTokenUrl = "https://sandbox-api.pchomepay.com.tw/v1/token"
)

type Client struct {
	APPID  string
	Secret string
}

type Token struct {
	Token           string `json:"token"`
	ExpireIn        int    `json:"expired_in"`
	ExpireTimestamp int    `json:"expired_timestamp"`
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Set(APPID, Secret string) *Client {
	c.Secret = Secret
	c.APPID = APPID
	return c
}

func (c *Client) GetTokenTest() (*Token, error) {
	res := Token{}
	client := &http.Client{}

	sEnc := b64.StdEncoding.EncodeToString([]byte(c.APPID + ":" + c.Secret))
	r, _ := http.NewRequest(http.MethodPost, TestTokenUrl, nil) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Basic "+sEnc)
	resp, err := client.Do(r)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	fmt.Println(resp.Status)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	print(string(body))
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetToken() (*Token, error) {
	res := Token{}
	client := &http.Client{}

	sEnc := b64.StdEncoding.EncodeToString([]byte(c.APPID + ":" + c.Secret))
	r, _ := http.NewRequest(http.MethodPost, TokenUrl, nil) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Basic "+sEnc)
	resp, err := client.Do(r)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	fmt.Println(resp.Status)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	print(string(body))
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
