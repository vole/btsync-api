package btsync_api

import (
  "strconv"
  "strings"
)

const endpoint = "http://127.0.0.1:%d/api/?"

type BTSyncAPI struct {
  Username string
  Password string
  Port     int
  Endpoint string
}

func New(login string, password string, port int) *BTSyncAPI {
  return &BTSyncAPI{login, password, port, endpoint}
}

func (api *BTSyncAPI) Request(method string, args map[string]string) *Request {
  return &Request{
    API:    api,
    Method: method,
    Args:   args,
  }
}

func (api *BTSyncAPI) AddFolderWithSecret(folder string, secret string) (*Response, error) {
  args := map[string]string{
    "dir": folder,
  }

  if secret != "" {
    args["secret"] = secret
  }

  request := api.Request("add_folder", args)

  var response Response
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) AddFolder(folder string) (*Response, error) {
  return api.AddFolderWithSecret(folder, "")
}

func (api *BTSyncAPI) RemoveFolder(secret string) (*Response, error) {
  request := api.Request("remove_folder", map[string]string{
    "secret": secret,
  })

  var response Response
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetFolder(secret string) (*GetFoldersResponse, error) {
  args := map[string]string{}

  if secret != "" {
    args["secret"] = secret
  }

  request := api.Request("get_folders", args)

  var response GetFoldersResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetFolders() (*GetFoldersResponse, error) {
  return api.GetFolder("")
}

func (api *BTSyncAPI) GetFilesForPath(secret string, path string) (*GetFilesResponse, error) {
  args := map[string]string{
    "secret": secret,
  }

  if path != "" {
    args["path"] = path
  }

  request := api.Request("get_files", args)

  var response GetFilesResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetFiles(secret string) (*GetFilesResponse, error) {
  return api.GetFilesForPath(secret, "")
}

func (api *BTSyncAPI) SetFilePrefs(secret string, path string, download int) (*SetFilePrefsResponse, error) {
  request := api.Request("set_file_prefs", map[string]string{
    "secret":   secret,
    "path":     path,
    "download": strconv.Itoa(download),
  })

  var response SetFilePrefsResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetFolderPeers(secret string) (*GetFolderPeersResponse, error) {
  request := api.Request("get_folder_peers", map[string]string{
    "secret": secret,
  })

  var response GetFolderPeersResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetSecretsForSecret(secret string, ecryption bool) (*GetSecretsResponse, error) {
  args := map[string]string{}

  if secret != "" {
    args["secret"] = secret
  }

  if ecryption {
    args["type"] = "encryption"
  }

  request := api.Request("get_secrets", args)

  var response GetSecretsResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetSecrets(encryption bool) (*GetSecretsResponse, error) {
  return api.GetSecretsForSecret("", encryption)
}

func (api *BTSyncAPI) GetFolderPrefs(secret string) (*GetFolderPrefsResponse, error) {
  request := api.Request("get_folder_prefs", map[string]string{
    "secret": secret,
  })

  var response GetFolderPrefsResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) SetFolderPrefs(secret string, prefs map[string]string) (*Response, error) {
  request := api.Request("set_folder_prefs", map[string]string{
    "secret": secret,
  })

  for key, value := range prefs {
    request.Args[key] = string(value)
  }

  var response Response
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetFolderHosts(secret string) (*GetFolderHostsResponse, error) {
  request := api.Request("get_folder_hosts", map[string]string{
    "secret": secret,
  })

  var response GetFolderHostsResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) SetFolderHosts(secret string, hosts []string) (*Response, error) {
  request := api.Request("set_folder_hosts", map[string]string{
    "secret": secret,
    "hosts":  strings.Join(hosts, ","),
  })

  var response Response
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetPreferences() (*GetPreferencesResponse, error) {
  request := api.Request("get_prefs", map[string]string{})

  var response GetPreferencesResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) SetPreferences(prefs Preferences) (*Response, error) {
  request := api.Request("set_prefs", map[string]string{})

  prefsMap := structToMap(prefs)

  for key, value := range prefsMap {
    request.Args[key] = string(value)
  }

  var response Response
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetOS() (*GetOSResponse, error) {
  request := api.Request("get_os", map[string]string{})

  var response GetOSResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetVersion() (*GetVersionResponse, error) {
  request := api.Request("get_version", map[string]string{})

  var response GetVersionResponse
  request.GetResponse(&response)

  return &response, nil
}

func (api *BTSyncAPI) GetSpeed() (*GetSpeedResponse, error) {
  request := api.Request("get_speed", map[string]string{})

  var response GetSpeedResponse
  request.GetResponse(&response)

  return &response, nil
}
