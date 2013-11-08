package btsync_api

import (
  "flag"
  //"fmt"
  ioutil "io/ioutil"
  "os"
  "path"
  "testing"
  "time"
)

// Where temporary test dirs are created.
var Dir = "/tmp"

// Reference to the temp dir used for testing.
var TmpDir = ""

// Reference to a temp file in the temp dir.
var TmpFile *os.File

// Prefix used by the temp dir for easier cleanup.
var Prefix = "btsync_api_test_"

// BTSync API credentials.
var login = flag.String("login", "test", "BT Sync API login")
var password = flag.String("password", "test", "BT Sync API password")

// BTSync API port.
var port = flag.Int("port", 8080, "BT Sync API port")

// Create a temp dir to use for tests.
func TestSetup(t *testing.T) {
  dir, err := ioutil.TempDir(Dir, Prefix)
  if err != nil {
    t.Errorf("Unable to create test directory in %s", Dir)
    return
  }

  TmpDir = dir

  file, err := ioutil.TempFile(TmpDir, Prefix)
  if err != nil {
    t.Errorf("Unable to create temp file in %s", TmpDir)
    return
  }

  TmpFile = file
}

// Test creating, removing folders.
func TestFolders(t *testing.T) {
  api := New(*login, *password, *port)
  addFolderResponse, err := api.AddFolder(TmpDir)

  if err != nil {
    t.Errorf("Error making request to add new folder")
    return
  }

  if addFolderResponse.Error != 0 {
    t.Errorf("Error adding new folder")
    return
  }

  getFoldersResponse, err := api.GetFolders()

  if err != nil {
    t.Errorf("Error making request to get all folders")
    return
  }

  if len(*getFoldersResponse) == 0 {
    t.Errorf("Expected at least 1 folder")
    return
  }

  found := false
  var testDir Folder
  for _, v := range *getFoldersResponse {
    if v.Dir == TmpDir {
      found = true
      testDir = v
    }
  }

  if !found {
    t.Errorf("Expected to find %s", TmpDir)
    return
  }

  getFolderResponse, err := api.GetFolder(testDir.Secret)
  if err != nil {
    t.Errorf("Error making request to get a single folder")
    return
  }

  if len(*getFolderResponse) != 1 {
    t.Errorf("Expected a single folder in response")
    return
  }

  if (*getFolderResponse)[0].Dir != TmpDir {
    t.Errorf("Expected %s to be %s", (*getFolderResponse)[0].Dir, TmpDir)
    return
  }

  time.Sleep(15000 * time.Millisecond)

  getFilesResponse, err := api.GetFiles(testDir.Secret)
  if err != nil {
    t.Errorf("Error making request to get files")
    return
  }

  if len(*getFilesResponse) != 1 {
    t.Errorf("Expected 1 file")
    return
  }

  if (*getFilesResponse)[0].Name != path.Base((*TmpFile).Name()) {
    t.Errorf("Expected %s to be %s", (*getFilesResponse)[0].Name, path.Base((*TmpFile).Name()))
    return
  }

  setFilePrefsResponse, err := api.SetFilePrefs(testDir.Secret, path.Base((*TmpFile).Name()), 1)
  if err != nil {
    t.Errorf("Error making request to set file preferences")
    return
  }

  if (*setFilePrefsResponse).State == "created" {
    t.Errorf("Expected file object in response")
    return
  }

  _, err = api.GetFolderPeers(testDir.Secret)
  if err != nil {
    t.Errorf("Error making request to get folder peers")
    return
  }

  getSecretsResponse, err := api.GetSecrets()
  if err != nil {
    t.Errorf("Error requesting secrets")
  }
}

type TestStruct struct {
  Foo string `json:"foo"`
  Bar string `json:"bar"`
}

// Test utility functions.
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

func TestCleanup(t *testing.T) {

}
