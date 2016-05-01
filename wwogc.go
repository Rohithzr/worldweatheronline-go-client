package wwogc

import (
    "net/http"
    "io/ioutil"
    )
//Gets the data from official world weather online servers
func GetData(query []Params, credentials Credentials, method string) (int, string){
	resp := ""
	status := 0
	uri := genWWOUrl(query, credentials, "searchApi")
    response, err := http.Get(uri)
    if err != nil {
        status = 0
        resp = ""
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            status = response.StatusCode
            resp = ""
        }
        status = response.StatusCode
        resp = string(contents)
    }
    return status, resp
}

