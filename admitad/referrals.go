package admitad

import (
	"github.com/dghubble/sling"
	"net/http"
)

type ReferralService struct {
	sling *sling.Sling
}

func NewReferralService (sling *sling.Sling) *ReferralService {
	return &ReferralService{
		sling: sling.Patch("referrals/"),
	}
}

type Referral struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Payment int `json:"payment"`
}

type ReferralList struct {
	Results []Referral `json:"results"`
}

type ReferralShowParams struct {
	DateStart string `url:"date_start"`
	DateEnd string `url:"date_end"`
}

func (s *ReferralService) Show(params *ReferralShowParams) ([]Referral, *http.Response, error) {
	referralList := new(ReferralList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(referralList, apiError)
	return referralList.Results, resp, err
}