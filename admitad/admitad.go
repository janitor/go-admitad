package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

const admitadApi = "https://api.admitad.com/"

type APIError struct {
}

type Client struct {
	sling   *sling.Sling
	Me      *MeService
	Tickets *TicketService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(admitadApi)
	return &Client{
		sling: base,

		Me:      NewMeService(base.New()),
		Tickets: NewTicketService(base.New()),
	}
}
