package btsync_api

import (
  "testing"
)

func TestAddFolder(t *testing.T) {
  response, err := AddFolder("new folder", "secret")

  if err != nil {
    t.Errorf("Error making request to add new folder")
    return
  }

  if response.Error != 0 {
    t.Errorf("Error adding new folder")
    return
  }
}

func TestRemoveFolder(t *testing.T) {
  response, err := RemoveFolder("secret")

  if err != nil {
    t.Errorf("Error making request to remove a folder")
    return
  }

  if response.Error != 0 {
    t.Errorf("Error removing folder")
    return
  }
}

func TestGetFolder(t *testing.T) {
  response, err := GetFolder("secret")

  if err != nil {
    t.Errorf("Error making request to get a folder")
    return
  }

  if response.Folders == nil {
    t.Errorf("Error getting folder")
    return
  }
}

func TestGetFolders(t *testing.T) {
  response, err := GetFolders()

  if err != nil {
    t.Errorf("Error making request to get all folder")
    return
  }

  if response.Folders == nil {
    t.Errorf("Error getting folders")
    return
  }
}

type TestStruct struct {
  Foo string `json:"foo"`
  Bar string `json:"bar"`
}

func TestUtils(t *testing.T) {
  s := &TestStruct{
    Foo: "foo",
    Bar: "bar",
  }

  m := structToMap(s)

  if m["foo"] != "foo" {
    t.Errorf("Error converting struct to map")
  }
}
