package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/persisais/go-freeipa/freeipa"
)

func main() {
	tspt := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // WARNING DO NOT USE THIS OPTION IN PRODUCTION
		},
	}

	c, err := freeipa.ConnectWithKerberosTicket("freeipa.example.test", tspt)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.UserShow(&freeipa.UserShowArgs{}, &freeipa.UserShowOptionalArgs{UID: freeipa.String("admin")})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Result.String())
}
