package admitad

import (
	"github.com/dghubble/sling"
	"net/http"
)

type PaymentSettingsService struct {
	sling *sling.Sling
}

func NewPaymentSettingsService(sling *sling.Sling) *PaymentSettingsService {
	return &PaymentSettingsService{
		sling: sling.Patch("me/payment/settings/"),
	}
}

type PaymentSettings struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Currency []string `json:"currency"`
}

type PaymentSettingsShowParams struct {
}

func (s *PaymentSettingsService) Show(params *PaymentSettingsShowParams) ([]PaymentSettings, *http.Response, error) {
	paymentSettingsList := []PaymentSettings{}
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(&paymentSettingsList, apiError)
	return paymentSettingsList, resp, err
}
