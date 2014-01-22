package GoOpenSdk 

type Oauth2 struct {

    AppKey string
    AppSecret string
    ResponseType string
    GrantType string
    Callback string
    Authorize string
    RequestCodeURL string
    AccessTokenURL string
}


(this *Oauth2) func GetRequestCode() string {
    
}

(this *Oauth2) func GetAccessToken() string {
    
}