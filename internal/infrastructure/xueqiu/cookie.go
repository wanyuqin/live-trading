package xueqiu

import (
	"net/http"
	"net/http/cookiejar"
)

var u = "https://xueqiu.com/hq"

var cookies []*http.Cookie

func GetCookie() error {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: cookieJar,
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	cookies = cookieJar.Cookies(req.URL)

	return nil
}
