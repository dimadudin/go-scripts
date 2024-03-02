package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (info authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", info.username, info.password)
}
