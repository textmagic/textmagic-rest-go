package textmagic

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"time"
)

const (
	interval    = time.Second
	clientUser  = ""
	clientToken = ""
)

var client = NewClient(clientUser, clientToken)

func toURLValues(m map[string]string) url.Values {
	u := url.Values{}

	for k, v := range m {
		u.Add(k, v)
	}

	return u
}

func debug(v ...interface{}) {
	for _, c := range v {
		fmt.Printf("%#v\n", c)
	}

	os.Exit(0)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())

	return rand.Intn(max-min) + min
}
