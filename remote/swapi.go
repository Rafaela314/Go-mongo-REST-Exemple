package swapi

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/settings"
)

type controller struct {
	SwapiURL string
}

type Swapi interface {
	Get(id, url string) ([]byte, error)
}

func New() Swapi {
	return &controller{SwapiURL: settings.Setting.SwapiURL}
}

func (c *controller) Get(key, url string) ([]byte, error) {
	head := map[string]string{
		"Content-Type": "application/json",
	}

	getPlanetsURL := fmt.Sprintf("%s/%s/%s", c.SwapiURL, url, key)

	fmt.Printf("/n URL SWA %v /n", getPlanetsURL)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", getPlanetsURL, nil)

	//Setting Headers
	for k, v := range head {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[GetInfos] - Error on make GET request, URL: %s , ERROR: %s", url, err.Error()))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[GetInfos] - Error on Read Body result, URL: %s, ERROR: %s", url, err.Error()))
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("[GetInfos] - Got Message error %d", resp.StatusCode))
	}

	return body, nil
}
