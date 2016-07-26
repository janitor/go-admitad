package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

var BrokenLinkErrorReasonDescription = map[int]string{
	0: "Program suspended",
	1: "You are not cooperating with the program",
	2: "Banner deleted",
	3: "Invalid deeplink domain",
}

type BrokenLinkService struct {
	sling *sling.Sling
}

func NewBrokenLinkService(sling *sling.Sling) *BrokenLinkService {
	return &BrokenLinkService{
		sling: sling.Patch("broken_links/"),
	}
}

type BrokenLink struct {
	ClickLink string `json:"click_link"`
	Clicks    int    `json:"clicks"`
	Datetime  string `json:"datetime"`
	Id        int    `json:"id"`
	ErrReason int    `json:"err_reason"`
	RefLink   string `json:"ref_link"`
	Campaign  struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"campaign"`
	Website struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"website"`
}

type BrokenLinkList struct {
	Results []BrokenLink `json:"results"`
}

type BrokenLinkShowParams struct {
	Limit     int    `url:"limit,omitempty"`
	Offset    int    `url:"offset,omitempty"`
	Reason    int    `url:"reason,omitempty"`
	Website   int    `url:"website,omitempty"`
	Campaign  int    `url:"campaign,omitempty"`
	DateStart string `url:"date_start,omitempty"`
	DateEnd   string `url:"date_end,omitempty"`
}

func (s *BrokenLinkService) Show(params *BrokenLinkShowParams) ([]BrokenLink, *http.Response, error) {
	brokenLinkList := new(BrokenLinkList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(brokenLinkList, apiError)
	return brokenLinkList.Results, resp, err
}
