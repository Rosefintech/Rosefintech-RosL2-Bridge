/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2020/11/6 16:02
*/

package tools

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpMethod int

const (
	Get HttpMethod = iota
	Post
	Put
	Delete
)



func HttpRequest(url, ContentType string, method HttpMethod, v io.Reader, header map[string]string) ([]byte, error) {
	var (
		trans = http.Transport{
			DisableKeepAlives: true,
			//Proxy: f,
			MaxIdleConnsPerHost: -1,
		}
		cli = &http.Client{
			Transport: &trans,
			//Timeout: time.Second*5,
		}
		err  error
		resp *http.Response
		req  *http.Request
		body []byte
	)
	defer cli.CloseIdleConnections()

	headerSet := func(req *http.Request) {
		if len(header) != 0 {
			for k, v := range header {
				req.Header.Set(k, v)
			}
		}
		return
	}
	err = nil
	switch method {
	case Get:
		req, err = http.NewRequest("GET", url, v)
		if err != nil {
			break
		}
		headerSet(req)

		break
	case Post:
		req, err = http.NewRequest("POST", url, v)
		if err != nil {
			break
		}
		headerSet(req)
		req.Header.Set("Content-Type", ContentType)

		break
	case Put:
		req, err = http.NewRequest("PUT", url, v)
		if err != nil {
			break
		}
		headerSet(req)
		req.Header.Set("Content-Type", ContentType)
		break
	case Delete:
		req, err = http.NewRequest("PUT", url, v)
		if err != nil {
			break
		}
		headerSet(req)
		req.Header.Set("Content-Type", ContentType)
		break
	default:
		return nil, errors.New("no http type")
	}
	//
	if err != nil {
		return nil, err
	}

	//
	resp, err = cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	Logger.Debug(string(body))
	return body, err
}
