package admitad

import (
	"net/http"

	"github.com/dghubble/sling"
)

type NewsService struct {
	sling *sling.Sling
}

func NewNewsService(sling *sling.Sling) *NewsService {
	return &NewsService{
		sling: sling.Patch("news/"),
	}
}

type News struct {
	Id           int    `json:"id"`
	Language     string `json:"language"`
	Url          string `json:"url"`
	Content      string `json:"content"`
	Datetime     string `json:"datetime"`
	ShortContent string `json:"short_content"`
}

type NewsList struct {
	Results []News `json:"results"`
}

type NewsShowParams struct {
	Language string `url:"language"`
}

func (s *NewsService) Show(params *NewsShowParams) ([]News, *http.Response, error) {
	newsList := new(NewsList)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(newsList, apiError)
	return newsList.Results, resp, err
}
