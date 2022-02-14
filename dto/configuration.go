package dto

//Configuration represents struct that serves to initialize receiver
type Configuration struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	TokenURL     string `json:"tokenUrl"`
	ApiHostName  string `json:"apiHostName"`
	PoweredBy    string `json:"poweredBy"`
	UserAgent    string `json:"userAgent"`
	BearerToken  string `json:"bearerToken"`
}
