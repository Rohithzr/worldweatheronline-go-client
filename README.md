# World Weather Online's API Client [wwogc]
--
[![Build Status](https://travis-ci.org/Rohithzr/worldweatheronline-api.svg?branch=master)](https://travis-ci.org/Rohithzr/worldweatheronline-api)

    import "github.com/Rohithzr/worldweatheronline_go_client"


## Usage

#### create Client

```go
credentials := wwogc.Credentials{
			"your_key",
			"json",
			"premium",
			"EN"
			}
```

#### create Query
```go
param1 := wwogc.Params{"q","Delhi"}
param2 := wwogc.Params("num_of_days", "1"}
query := []wwogc.Params{param1, param2}	
```

#### Methods
Methods available are: 

1. Search API (searchApi)
2. Local Weather API (localWeatherApi)
3. Time Zone API (tizeZoneApi)
4. Ski Weather API (skiWeatherApi)
5. Marine Weather API (marineWeatherApi)
6. Historical Weather API (historicalWeatherApi)

#### func  GetData

Pass the Above created query and credentials variable and the method you want to call

```go
 code, data := GetData(query, credentials, "searchApi")
```
Gets the data from official world weather online servers and returns
 - code: HTTP Status Code (0 if error)
 - data: HTTP Body (blank if error)
