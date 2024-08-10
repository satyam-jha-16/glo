package main

import (
	"bytes"
	"io"
	"net/http"
)

type MakeRequestCfg struct {
	Method  string
	Url     string
	Body    string
	Headers map[string]string
}

type MakeRequestRes struct {
	Response *http.Response
	Body     string
}

func MakeRequest(cfg *MakeRequestCfg) (*MakeRequestRes, error) {
  client := &http.Client{}

	bodyInpBuff := bytes.NewBufferString(cfg.Body)

	req, err := http.NewRequest(
		cfg.Method,
    cfg.Url,
		bodyInpBuff,
	)
	if err != nil {
    return nil, err
	}

  for key, val := range(cfg.Headers){
    req.Header.Set(key, val,)
  }
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
    return nil, err
	}

  return &MakeRequestRes{
  Response : res,
    Body : string(resBody),
  }, nil
}

