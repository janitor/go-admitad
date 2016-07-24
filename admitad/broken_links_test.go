package admitad

import (
	"net/http"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestBrokenLinkService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/broken_links/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
    "results": [
        {
            "campaign": {
                "id": 6,
                "name": "AdvCamp 1"
            },
            "click_link": "http://ad.admitad.com/g/395b832b8259505879f5234642e5a7/",
            "clicks": 1,
            "datetime": "2015-10-26T12:34:55",
            "err_reason": 2,
            "id": 20,
            "ref_link": "https://www.google.by/?gfe_rd=cr&ei=GfQtVoaSAqKF8QeN_rv4Ag",
            "website": {
                "id": 22,
                "name": "site1_of_webmaster1"
            }
        }]}`)
	})

	client := NewClient(httpClient)
	links, _, err := client.BrokenLinks.Show(&BrokenLinkShowParams{})
	expected := BrokenLink{
		Clicks: 1,
		Campaign: struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{
			Id:   6,
			Name: "AdvCamp 1",
		},
		Website: struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{
			Id:   22,
			Name: "site1_of_webmaster1",
		},
		ClickLink: "http://ad.admitad.com/g/395b832b8259505879f5234642e5a7/",
		Datetime:  "2015-10-26T12:34:55",
		ErrReason: 2,
		Id:        20,
		RefLink:   "https://www.google.by/?gfe_rd=cr&ei=GfQtVoaSAqKF8QeN_rv4Ag",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, links[0])
}
