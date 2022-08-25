package controllers

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// PostSoap Envia una solicitud a la AFIP y devuelve el resultado del body de la respuesta.
func PostSoap(payload []byte, endPoint, soapAction string) ([]byte, error) {

	// wsdl service url
	// prepare the request
	req, err := http.NewRequest("POST", endPoint, bytes.NewReader(payload))
	if err != nil {
		fmt.Println("Error on creating request object.", err.Error())
		return nil, err
	}

	// set the content type header, as well as the oter required headers
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Set("SOAPAction", soapAction)

	// prepare the client request
	client := &http.Client{
		Timeout: GetTimeOutHttp(),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on dispatching request. ", err.Error())
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return body, nil
}

func GetTimeOutHttp() time.Duration {

	var timeOut time.Duration
	timeStr := os.Getenv("HTTP_TIME_OUT")
	if timeStr != "" {
		timeOut, _ = time.ParseDuration(os.Getenv("HTTP_TIME_OUT"))
	} else {
		timeOut, _ = time.ParseDuration("30s")
	}

	return timeOut
}
