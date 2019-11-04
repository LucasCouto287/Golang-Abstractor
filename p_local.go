package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Local struct {
	Path string
}

func (l *Local) Auth() string {
	return "Auth Local"
}

func (l *Local) ListBuckets() string {
	listdir, _ := ioutil.ReadDir(l.Path)
	for i := range listdir {
		fmt.Println(listdir[i])
	}
	return "Local List Buckets"
}
