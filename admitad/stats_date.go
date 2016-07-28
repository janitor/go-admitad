package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

type StatsDateService struct {
	sling *sling.Sling
}

func NewStatsDateService(sling *sling.Sling) *StatsDateService {
	return &StatsDateService{
		sling: sling.Patch("statistics/dates/"),
	}
}

type StatsDate struct {
	Date               string  `json:"date"`
	Currency           string  `json:"currency"`
	LeadsSum           int     `json:"leads_sum"`
	SalesSum           int     `json:"sales_sum"`
	Views              int     `json:"views"`
	Clicks             int     `json:"clicks"`
	PaymentSumDeclined float32 `json:"payment_sum_declined"`
	PaymentSumApproved float32 `json:"payment_sum_approved"`
	PaymentSumOpen     float32 `json:"payment_sum_open"`
}

type StatsDateList struct {
	Results []StatsDate `json:"results"`
}

type StatsDateShowParams struct {
	DateStart string `url:"date_start,omitempty"`
	DateEnd   string `url:"date_end,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Offset    int    `url:"offset,omitempty"`
	Website   int    `url:"website,omitempty"`
	Campaign  int    `url:"campaign,omitempty"`
}

func (s *StatsDateService) Show(params *StatsDateShowParams) ([]StatsDate, *http.Response, error) {
	statsDateList := new(StatsDateList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(statsDateList, apiError)
	return statsDateList.Results, resp, err
}
