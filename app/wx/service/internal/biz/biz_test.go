package biz

import (
	"encoding/json"
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	"testing"
)

func TestMenuStr(t *testing.T) {
	rst := &v1.Menu{}
	err := json.Unmarshal([]byte(menu), rst)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(rst.Button[0].Name)
	}

}
