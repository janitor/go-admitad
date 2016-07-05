package main

import (
	"flag"
	"fmt"

	"github.com/janitor/go-admitad/admitad"
	"golang.org/x/oauth2"
)

func main() {
	accessToken := flag.String("access-token", "", "Admitad Access Token")
	flag.Parse()

	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: *accessToken}
	httpClient := config.Client(oauth2.NoContext, token)

	client := admitad.NewClient(httpClient)
	params := &admitad.MeShowParams{}
	me, _, _ := client.Me.Show(params)
	fmt.Println(me)
}
