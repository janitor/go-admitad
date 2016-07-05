# go-admitad [![Build Status](https://travis-ci.org/janitor/go-admitad.svg?branch=master)](https://travis-ci.org/janitor/go-admitad)

go-admitad is a Go library for the [Admitad API](https://developers.admitad.com/en/).


## Install

    go get github.com/janitor/go-admitad/admitad

## Documentation

[GoDoc](https://godoc.org/github.com/janitor/go-admitad/admitad)

## Usage

```go
config := &oauth2.Config{}
token := &oauth2.Token{AccessToken: *accessToken}
httpClient := config.Client(oauth2.NoContext, token)

// api client
client := admitad.NewClient(httpClient)
// user info
me, resp, err := client.Me.Show(&admitad.MeShowParams{})
```

## Testing

    go test

## License

[MIT License](LICENSE)