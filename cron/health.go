package cron

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func GetHealth(ctx context.Context, url *url.URL) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://"+url.Host+"/health"), nil)
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("server is down")
	}
	return nil
}
