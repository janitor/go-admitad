package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

type TicketService struct {
	sling *sling.Sling
}

func NewTicketService(sling *sling.Sling) *TicketService {
	return &TicketService{
		sling: sling.Patch("tickets/"),
	}
}

type Ticket struct {
	Id           int    `json:"id"`
	Text         string `json:"text"`
	PlainText    string `json:"plain_text"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	Priority     int    `json:"priority"`
	Status       int    `json:"status"`
	Subject      string `json:"subject"`
}

type TicketList struct {
	Results []Ticket `json:"results`
}

type TicketShowParams struct {
}

func (s *TicketService) Show(params *TicketShowParams) ([]Ticket, *http.Response, error) {
	ticketList := new(TicketList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(ticketList, apiError)
	return ticketList.Results, resp, err
}
