package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func Download(w io.Writer, uri string, session string) error {
	url, err := url.Parse(uri)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = io.Copy(w, res.Body)
	if err != nil {
		return err
	}
	return nil
}

const sessionTokenEnvKey = "AOC_SESSION"

var (
	errSessionTokenNotFoundInEnv = fmt.Errorf("session token not found (please set %s=token)", sessionTokenEnvKey)
)

func GetSessionToken() (token string, err error) {
	token = os.Getenv(sessionTokenEnvKey)
	if token == "" {
		err = errSessionTokenNotFoundInEnv
	}
	return token, err
}
