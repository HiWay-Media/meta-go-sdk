package ig

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func (o *igService) HealthCheck() error {
	resp, err := o.restyPost("/", nil)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("healthcheck error")
	}
	return nil
}

func (o *igService) IsDebug() bool {
	return o.debug
}

// Resty Methods

func (o *igService) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+o.accessToken).
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *igService) restyPostFormUrlEncoded(url string, data map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(data).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *igService) restyPostWithQueryParams(url string, body interface{}, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+o.accessToken).
		SetQueryParams(queryParams).
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *igService) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Authorization", "Bearer "+o.accessToken).
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *igService) debugPrint(data ...interface{}) {
	if o.debug {
		log.Println(data...)
	}
}
