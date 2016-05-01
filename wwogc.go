package wwogc

import (
    "fmt"
    "net/http"
    "io/ioutil"
    )

func GetData(query []Params, credentials Credentials, method string) string{
	resp := ""
	uri := genWWOUrl(query, credentials, "searchApi")
    response, err := http.Get(uri)
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
        }
        resp = string(contents)
    }
    return resp
}