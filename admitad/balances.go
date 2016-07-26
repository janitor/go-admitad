package admitad

import (
	"github.com/dghubble/sling"
	"net/http"
)

type BalanceService struct {
	sling *sling.Sling
}

func NewBalanceService(sling *sling.Sling) *BalanceService {
	return &BalanceService{
		sling: sling.Patch("me/balance/"),
	}
}

type Balance struct {
	Currency   string `json:"currency"`
	Balance    string `json:"balance"`
	Processing string `json:"processing"`
	Stalled    string `json:"stalled"`
	Today      string `json:"today"`
}

type BalanceList struct {
	Results []Balance `json:"results"`
}

type BalanceShowParams struct {
}

func (s *BalanceService) Show(params *BalanceShowParams) ([]Balance, *http.Response, error) {
	balanceList := []Balance{}
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(&balanceList, apiError)
	return balanceList, resp, err
}
