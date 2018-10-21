package common

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func Get(url string, token string, failedUrls map[string]string) {
	log.Infof("request %s", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Error(err)
	}

	req.Header.Add("Authorization", token)

	response, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var resp Response

	e := json.Unmarshal(respBody, &resp)
	if e != nil {
		log.Error(e)
	}

	if !resp.Ok {
		failedUrls[url] = url
		log.Info(string(respBody), url)
	}
}