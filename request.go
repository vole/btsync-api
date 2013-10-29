package btsync_api

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

type Request struct {
  Method string
  Args   map[string]string
}

func (request *Request) URL() string {
  args := request.Args

  args["method"] = request.Method

  params := url.Values{}
  for key, value := range args {
    params.Add(key, value)
  }

  s := fmt.Sprintf(endpoint, Port) + params.Encode()
  return s
}

func (request *Request) Get() (response []byte, ret error) {
  if request.Method == "" {
    return nil, errors.New("Missing method")
  }

  s := request.URL()

  res, err := http.Get(s)
  defer res.Body.Close()
  if err != nil {
    return nil, err
  }

  body, _ := ioutil.ReadAll(res.Body)
  return body, nil
}

func (request *Request) GetResponse(response interface{}) (*interface{}, error) {
  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}
