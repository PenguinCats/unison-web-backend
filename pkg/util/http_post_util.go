package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: 5 * time.Second,
	}
)

type HttpPostResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpPost(url string, payload interface{}, d interface{}) (err error) {
	marshal, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("http error")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := HttpPostResponse{
		Data: d,
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return errors.New(res.Msg)
	}

	return nil
}
