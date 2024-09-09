module github.com/PNAP/go-sdk-helper-bmc

//replace github.com/PNAP/go-sdk-helper-bmc => D:/repos/go-sdk-helper-bmc

go 1.13

require (
	//github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/phoenixnap/go-sdk-bmc/auditapi/v3 v3.0.2
	github.com/phoenixnap/go-sdk-bmc/billingapi/v3 v3.0.0
	github.com/phoenixnap/go-sdk-bmc/bmcapi/v3 v3.1.0
	github.com/phoenixnap/go-sdk-bmc/invoicingapi v1.0.2
	github.com/phoenixnap/go-sdk-bmc/ipapi/v3 v3.0.2
	github.com/phoenixnap/go-sdk-bmc/locationapi/v3 v3.0.0
	github.com/phoenixnap/go-sdk-bmc/networkapi/v3 v3.0.0-20240904202613-fe33e8c87c5f
	github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3 v3.0.0
	github.com/phoenixnap/go-sdk-bmc/paymentsapi v1.0.2
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v3 v3.0.2
	github.com/phoenixnap/go-sdk-bmc/tagapi/v3 v3.0.2
	github.com/spf13/viper v1.6.2
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
