package main

import "fmt"

type authenticationInfo struct {
  username string
  password string
}

func (authInfo authenticationInfo) getBasicAuth() string {

  return "Authorication: Basic" + authInfo.username + ":" + authInfo.password
}

func test(authInfo authenticationInfo) {
  fmt.Println(authInfo.getBasicAuth())
  fmt.Println("================================")
}

func main() {
	test(authenticationInfo{
		username: "Sachu",
		password: "GigaChad",
	})
  test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
