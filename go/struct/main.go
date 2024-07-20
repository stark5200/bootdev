package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func getBasicAuth(info authenticationInfo) formatted string {
	formatted := "Authorization: Basic " + authenticationInfo.username + ":" + authenticationInfo.password
	return
}
