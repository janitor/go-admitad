package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

type MeService struct {
	sling *sling.Sling
}

func NewMeService(sling *sling.Sling) *MeService {
	return &MeService{
		sling: sling.Path("me/"),
	}
}

type Me struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	DefaultCurrency string `json:"default_currency"`
}

type MeShowParams struct {
}

func (s *MeService) Show(params *MeShowParams) (Me, *http.Response, error) {
	me := new(Me)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(me, apiError)
	return *me, resp, err
}
