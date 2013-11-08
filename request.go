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
  API    *BTSyncAPI
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

  s := fmt.Sprintf(request.API.Endpoint, request.API.Port) + params.Encode()
  return s
}

func (request *Request) Get() (response []byte, ret error) {
  if request.Method == "" {
    return nil, errors.New("Missing method")
  }

  s := request.URL()

  if request.API.Debug {
    request.API.Logger.Printf("\033[32;1mGET:\033[0m %s\n", s)
  }

  client := &http.Client{}

  // BUG(aaron): Currently nothing to handle the case where Basic Auth fails.
  req, err := http.NewRequest("GET", s, nil)
  req.SetBasicAuth(request.API.Username, request.API.Password)

  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  defer res.Body.Close()

  body, _ := ioutil.ReadAll(res.Body)
  return body, nil
}

func (request *Request) GetResponse(response interface{}) (*interface{}, error) {
  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  if request.API.Debug {
    request.API.Logger.Printf("\033[32;1mRES:\033[0m %s\n", rawJson)
  }

  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}
