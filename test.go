package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

// this is a test file where I have first tried out some of the basics of reqiest handling in golang
func Testing(t *testing.T) {

	resp, err := http.Get("www.example.com/")

	if err != nil {
		t.Fatal(err)
	}

	var respBodyData = make([]byte, 100)

	_, err = resp.Body.Read(respBodyData)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(respBodyData))

}

func TestDo(t *testing.T) {
	client := &http.Client{}

	//some input parsing

	bodyIntString := `{"username":"satyam","password":"satyam123"}`

	bodyInpBuff := bytes.NewBufferString(bodyIntString)

	req, err := http.NewRequest(
		"POST",
		"http://localhost:3000/api/auth",
		bodyInpBuff,
	)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resBody)

}

func TestMakeRequest(t *testing.T) {
	res, err := MakeRequest(&MakeRequestCfg{
		Method: "GET",
		Url:    "https://github.com/satyam-jha-16/",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Body)
}
