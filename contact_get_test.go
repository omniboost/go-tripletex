package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestContactGet(t *testing.T) {
	req := client.NewContactGetRequest()
	// req.QueryParams().Count = 1
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
