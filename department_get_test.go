package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDepartmentGet(t *testing.T) {
	req := client.NewDepartmentGetRequest()
	// req.QueryParams().Count = 1
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
