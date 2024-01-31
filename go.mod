module github.com/PNAP/go-sdk-helper-bmc

//replace github.com/PNAP/go-sdk-helper-bmc => D:/repos/go-sdk-helper-bmc

go 1.13

require (
	//github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/phoenixnap/go-sdk-bmc/auditapi/v2 v2.0.6
	github.com/phoenixnap/go-sdk-bmc/billingapi/v2 v2.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/bmcapi/v3 v3.0.0-20240125122650-80a2524401f4
	github.com/phoenixnap/go-sdk-bmc/invoicingapi v0.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/ipapi/v3 v3.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/locationapi/v2 v2.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/networkapi/v3 v3.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v2 v2.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v3 v3.0.0-20240123132820-9c39f0808161
	github.com/phoenixnap/go-sdk-bmc/tagapi/v3 v3.0.0-20240123132820-9c39f0808161
	github.com/spf13/viper v1.6.2
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
