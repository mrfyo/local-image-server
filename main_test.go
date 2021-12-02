package main

import "testing"

func TestGetFileType(t *testing.T) {

	filename := "abc.png"

	if GetFileType(filename) != "png" {
		t.FailNow()
	}
}
