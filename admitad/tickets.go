package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	TICKET_STATUS_NEW      = 1
	TICKET_STATUS_UNDERWAY = 2
	TICKET_STATUS_WAIT     = 3
	TICKET_STATUS_CLOSED   = 4
)

const (
	TICKET_PRIORITY_NORMAL    = 1
	TICKET_PRIORITY_IMPORTANT = 2
	TICKET_PRIORITY_PROBLEM   = 3
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
	Category     struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	AdvCampaign struct {
		Id   int    `json:"id"`
		Name string `json"name"`
	}
}

type TicketList struct {
	Results []Ticket `json:"results`
}

type TicketShowParams struct {
	Status    int    `url:"status,omitempty"`
	DateStart string `url:"date_start,omitempty"`
	DateEnd   string `url:"date_end,omitempty"`
}

func (s *TicketService) Show(params *TicketShowParams) ([]Ticket, *http.Response, error) {
	ticketList := new(TicketList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(ticketList, apiError)
	return ticketList.Results, resp, err
}

//type TicketCreateParams struct {
//	Subject     string `json:"subject"`
//	Text        string `json:"text"`
//	AdvCampaign int    `json:"advcampaign"`
//	Category    int    `json:"category"`
//	Priority    int    `json:"priority"`
//}
//
//func (s *TicketService) Create(params *TicketCreateParams) (Ticket, *http.Response, error) {
//	ticket := new(Ticket)
//	apiError := new(APIError)
//	resp, err := s.sling.New().Post("create/").BodyJSON(params).Receive(ticket, apiError)
//	return *ticket, resp, err
//}
