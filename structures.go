package wwogc

//Credentials{Key,ResponseType,Subscription,Locale} will create an api client that will send requests to the server where 
//Key is your API Key, 
//ResponseType is response Format (json, xml)
//Subscription is Subscription Type
//Locale is Language code as per ISO Codes speified here http://developer.worldweatheronline.com/api/docs/multilingual.aspx
type Credentials struct {
    Key string
    ResponseType  string
    Subscription string
    Locale string
}

//You will ned to create an Array of Params{name:value} which will be the parameters to be passed to the server to fetch data
type Params struct {
    Name string
    Value  string
}