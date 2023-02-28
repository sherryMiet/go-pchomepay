package go_pchomepay_sdk

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func SendPCHOMEPayRequest(data []byte, token string, Url string) (res []byte, err error) {
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, Url, bytes.NewReader(data)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("pcpay-token", token)
	resp, err := client.Do(r)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	fmt.Println(resp.Status)

	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	print(string(res))

	return res, nil
}
