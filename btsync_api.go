package btsync_api

import (
  "encoding/json"
  "strings"
)

const endpoint = "http://127.0.0.1:%d/api/?"

var Port int = 12345

func AddFolder(folder string, secret string) (*BasicResponse, error) {
  request := &Request{
    Method: "add_folder",
    Args: map[string]string{
      "dir":    folder,
      "secret": secret,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response BasicResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func RemoveFolder(secret string) (*BasicResponse, error) {
  request := &Request{
    Method: "remove_folder",
    Args: map[string]string{
      "secret": secret,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response BasicResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetFolder(secret string) (*GetFoldersResponse, error) {
  request := &Request{
    Method: "get_folders",
    Args: map[string]string{
      "secret": secret,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetFoldersResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetFolders() (*GetFoldersResponse, error) {
  return GetFolder("")
}

func GetFilesForPath(secret string, path string) (*GetFilesResponse, error) {
  request := &Request{
    Method: "get_files",
    Args: map[string]string{
      "secret": secret,
      "path":   path,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetFilesResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetFiles(secret string) (*GetFilesResponse, error) {
  return GetFilesForPath(secret, "")
}

func SetFilePrefs(secret string, path string, download bool) (*BasicResponse, error) {
  request := &Request{
    Method: "set_files_prefs",
    Args: map[string]string{
      "secret":   secret,
      "path":     path,
      "download": boolToString(download),
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response BasicResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetFolderPeers(secret string) (*GetFolderPeersResponse, error) {
  request := &Request{
    Method: "get_folder_peers",
    Args: map[string]string{
      "secret": secret,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetFolderPeersResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetSecrets() (*GetSecretsResponse, error) {
  request := &Request{
    Method: "get_secrets",
    Args:   map[string]string{},
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetSecretsResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetFolderPrefs(secret string) (*GetFolderPrefsResponse, error) {
  request := &Request{
    Method: "get_folder_prefs",
    Args: map[string]string{
      "secret": secret,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetFolderPrefsResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func SetFolderPrefs(secret string, prefs map[string]string) (*BasicResponse, error) {
  request := &Request{
    Method: "set_folder_prefs",
    Args: map[string]string{
      "secret": secret,
    },
  }

  for key, value := range prefs {
    request.Args[key] = string(value)
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response BasicResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetFolderHosts(secret string) (*GetFolderHostsResponse, error) {
  request := &Request{
    Method: "get_folder_hosts",
    Args: map[string]string{
      "secret": secret,
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetFolderHostsResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func SetFolderHosts(secret string, hosts []string) (*BasicResponse, error) {
  request := &Request{
    Method: "set_folder_hosts",
    Args: map[string]string{
      "secret": secret,
      "hosts":  strings.Join(hosts, ","),
    },
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response BasicResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetPreferences() (*GetPreferencesResponse, error) {
  request := &Request{
    Method: "get_prefs",
    Args:   map[string]string{},
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetPreferencesResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func SetPreferences(prefs Preferences) (*BasicResponse, error) {
  request := &Request{
    Method: "set_prefs",
    Args:   map[string]string{},
  }

  prefsMap := structToMap(prefs)

  for key, value := range prefsMap {
    request.Args[key] = string(value)
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response BasicResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetOS() (*GetOSResponse, error) {
  request := &Request{
    Method: "get_os",
    Args:   map[string]string{},
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetOSResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetVersion() (*GetVersionResponse, error) {
  request := &Request{
    Method: "get_version",
    Args:   map[string]string{},
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetVersionResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}

func GetSpeed() (*GetSpeedResponse, error) {
  request := &Request{
    Method: "get_speed",
    Args:   map[string]string{},
  }

  rawJson, err := request.Get()
  if err != nil {
    return nil, err
  }

  var response GetSpeedResponse
  if err := json.Unmarshal(rawJson, &response); err != nil {
    return nil, err
  }

  return &response, nil
}
