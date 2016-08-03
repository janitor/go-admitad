package admitad

import (
	"github.com/dghubble/sling"
	"net/http"
)

type PayoutService struct {
	sling *sling.Sling
}

func NewPayoutService(sling *sling.Sling) *PayoutService {
	return &PayoutService{
		sling: sling.Patch("payments/"),
	}
}

type Payout struct {
	Id           int    `json:"id"`
	Status       string `json:"status"`
	Comment      string `json:"comment"`
	Datetime     string `json:"datetime"`
	HasStatement bool   `json:"has_statement"`
}

type PayoutList struct {
	Results []Payout `json:"results"`
}

type PayoutShowParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (s *PayoutService) Show(params *PayoutShowParams) ([]Payout, *http.Response, error) {
	payoutList := new(PayoutList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(payoutList, apiError)
	return payoutList.Results, resp, err
}

//type PayoutCreateParams struct {
//	Currency string
//}
//
//func (s *PayoutService) Create(params *PayoutCreateParams) (Payout, *http.Response, error) {
//	payout := new(Payout)
//	apiError := new(APIError)
//	pathURL := fmt.Sprintf("request/%s/", params.Currency)
//	resp, err := s.sling.New().Post(pathURL).BodyJSON(params).Receive(payout, apiError)
//	return *payout, resp, err
//}
