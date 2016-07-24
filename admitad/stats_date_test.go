package admitad

import (
	"net/http"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestStatsDateService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/statistics/dates/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `
		{
			"results": [
			{
					"leads_sum": 0,
					"ctr": 0.0,
					"views": 0,
					"payment_sum_declined": 0.0,
					"clicks": 2,
					"payment_sum_approved": -1254.81,
					"currency": "RUB",
					"ecpm": 0.0,
					"sales_sum": 4,
					"sales_sum_approved": 1,
					"sales_sum_open": 1,
					"sales_sum_declined": 2,
					"leads_sum_approved": 1,
					"leads_sum_open": 2,
					"leads_sum_declined": 0,
					"date": "2011-02-24",
					"cr": 0.0,
					"ecpc": 2509.615,
					"payment_sum_open": 6274.04
				}
			]
		}`)
	})

	client := NewClient(httpClient)
	stats, _, err := client.StatsDate.Show(&StatsDateShowParams{})
	expected := StatsDate{
		Date:     "2011-02-24",
		Views:    0,
		Clicks:   2,
		Currency: "RUB",
		LeadsSum: 0,
		SalesSum: 4,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, stats[0])
}
