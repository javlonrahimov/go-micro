package main

import (
	"authentication/data"
	"os"
	"testing"
)

var testApp Config

func TesTestMain(m *testing.M) {
	repo := data.NewPostgressTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}