package tripletex_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDepartmentQueryGet(t *testing.T) {
	req := client.NewDepartmentQueryGetRequest()
	req.QueryParams().Query = "02 Revier"
	req.QueryParams().Fields = "displayName"
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
