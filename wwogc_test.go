package wwogc

import (
	"testing"
    )
func TestGetData(t *testing.T) {
	param1 := Params{"q","Delhi"}
	query := []Params{param1}	
	credentials := Credentials{"9d378416be224717b92172807162004","json","premium","EN"}
	code, data := GetData(query, credentials, "searchApi")
	if code == 0 || len(data) <= 0 {
		t.Errorf("No Content/Internet")
	}
	if code > 0 {
		switch code {
			case 400: t.Errorf("Error: Bad request. Invalid input parameters")
			case 401: t.Errorf("Error: Unauthorized key")
			case 403: t.Errorf("Error: Forbidden. Tampered URL")
			case 404: t.Errorf("Error: Not found")
			case 410: t.Errorf("Error: URL expired")
			case 500: t.Errorf("Error: Internal server error")
			case 503: t.Errorf("Error: Service unavailable")
			case 599: t.Errorf("Error: Connection timed out")
		}
	}
}