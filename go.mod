module github.com/PNAP/go-sdk-helper-bmc

//replace github.com/PNAP/go-sdk-helper-bmc => D:/repos/go-sdk-helper-bmc

go 1.13

require (
	//github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/phoenixnap/go-sdk-bmc/auditapi/v2 v2.0.6
	github.com/phoenixnap/go-sdk-bmc/billingapi v1.3.2
	github.com/phoenixnap/go-sdk-bmc/bmcapi/v2 v2.3.1-0.20231010082231-b3a9e39e365a
	github.com/phoenixnap/go-sdk-bmc/ipapi/v2 v2.0.5
	github.com/phoenixnap/go-sdk-bmc/locationapi v1.0.0
	github.com/phoenixnap/go-sdk-bmc/networkapi/v2 v2.3.1
	github.com/phoenixnap/go-sdk-bmc/networkstorageapi v1.4.0
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2 v2.0.5
	github.com/phoenixnap/go-sdk-bmc/tagapi/v2 v2.0.5
	github.com/spf13/viper v1.6.2
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
