package main

import (
	"github.com/bankole7782/zazabul"
	"fmt"
)


var template = `
// email is used for communication
// email must follow the format for email ie username@host.ext
// email is compulsory
email: banker@banban.com

// region should be gotten from google cloud documentation
region: us-central1

// zone should be gotten from google cloud documentation
// zone usually derived from the regions and ending with -a or -b or -c
zone: us-central1-a

// test of colon
addr: https://example.com


`
func main() {
	conf, err := zazabul.ParseConfig(template)
	if err != nil {
		panic(err)
	}
	fmt.Println(conf)

	fmt.Println("email: ", conf.Get("email"))

	conf.Update(map[string]string {
		"email": "banban@banban.com",
		"zone": "us-central1-c",
	})

	fmt.Println("email: ", conf.Get("email"))
	fmt.Println("jaker: ", conf.Get("jaker"))

	err = conf.Write("/tmp/test.zaz")
	if err != nil {
		panic(err)
	}
}
