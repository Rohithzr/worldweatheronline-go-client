package wwogc

func genQueryParams(wwogc_params []Params) string{
	queryString := ""
	for i, param := range wwogc_params{
		if param.Name != ""{
			queryString += param.Name + "=" + param.Value
			if i != len(wwogc_params)-1{
				queryString += "&"
			}
		}
	}
	return queryString
}

func genUrlEP(method string) string{
	avlMethods := make(map[string] string)
	avlMethods["searchApi"] = "search"
    avlMethods["localWeatherApi"] = "weather"
    avlMethods["timeZoneApi"] = "tz"
    avlMethods["skiWeatherApi"] = "ski"
    avlMethods["marineWeatherApi"] = "marine"
    avlMethods["historicalWeatherApi"] = "past-weather"

    return avlMethods[method]+".ashx"
}

func genMidUrl(subscription string) string{
	avlSubscription := make(map[string] string)
	avlSubscription["free"] = "/free/v2/"
	avlSubscription["premium"] = "/premium/v1/"

	return avlSubscription[subscription]
}

func genWWOUrl(query []Params, credentials Credentials, method string) string{
    url := ""
    key := credentials.Key
    subscription := credentials.Subscription
    format := credentials.ResponseType
    lang := credentials.Locale
    preUrl := "https://api.worldweatheronline.com" + genMidUrl(subscription) + genUrlEP(method) + "?key=" + key + "&format=" + format + "&lang=" + lang + "&"
    url = preUrl + genQueryParams(query)
    return url
}