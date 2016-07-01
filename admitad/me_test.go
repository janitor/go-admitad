package admitad

import (
	"net/http"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestMeService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"username": "test_username"}`)
	})

	client := NewClient(httpClient)
	me, _, err := client.Me.Show(&MeShowParams{})
	expected := Me{Username: "test_username"}
	assert.Nil(t, err)
	assert.Equal(t, expected, me)
}
