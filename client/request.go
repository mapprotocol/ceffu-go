package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const defaultHTTPTimeout = 20 * time.Second

func (c *client) request(ctx context.Context, path, method string, headers http.Header, body io.Reader) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		request.Header = headers
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("response is nil")
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		var data []byte
		if b, err := io.ReadAll(resp.Body); err == nil {
			data = b
		}

		return nil, NewRequestError(
			path,
			WithCode(strconv.Itoa(resp.StatusCode)),
			WithMessage(string(data)),
			WithBody(data),
		)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *client) Get(ctx context.Context, path string, params interface{}) ([]byte, error) {
	encoded, err := URLEncode(params)
	if err != nil {
		return nil, err
	}
	if encoded != "" {
		path += "?" + encoded
	}
	signature, err := c.sign(encoded)

	headers := http.Header{
		"Content-Type": []string{"application/json"},
		"open-apikey":  []string{c.apiKey},
		"signature":    []string{signature},
	}
	return c.request(ctx, fmt.Sprintf("%s%s", c.domain, path), http.MethodGet, headers, nil)
}

func (c *client) Post(ctx context.Context, path string, body interface{}) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	signature, err := c.sign(string(data))

	headers := http.Header{
		"Content-Type": []string{"application/json"},
		"open-apikey":  []string{c.apiKey},
		"signature":    []string{signature},
	}
	return c.request(ctx, fmt.Sprintf("%s%s", c.domain, path), http.MethodPost, headers, bytes.NewReader(data))
}

func URLEncode(s interface{}) (string, error) {
	if s == nil {
		return "", errors.New("provided value is nil")
	}

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return "", errors.New("provided value is not a struct")
	}

	typ := val.Type()
	urls := url.Values{}
	for i := 0; i < val.NumField(); i++ {
		if !val.Field(i).CanInterface() {
			continue
		}
		name := typ.Field(i).Name
		tag := typ.Field(i).Tag.Get("json")
		if tag != "" {
			index := strings.Index(tag, ",")
			if index == -1 {
				name = tag
			} else {
				name = tag[:index]
			}
		}
		urls.Set(name, fmt.Sprintf("%v", val.Field(i).Interface()))
	}
	return urls.Encode(), nil
}
