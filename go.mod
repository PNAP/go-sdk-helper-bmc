module github.com/PNAP/go-sdk-helper-bmc

//replace github.com/PNAP/go-sdk-helper-bmc => D:/repos/go-sdk-helper-bmc

go 1.13

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/phoenixnap/go-sdk-bmc/auditapi v1.0.0
	github.com/phoenixnap/go-sdk-bmc/bmcapi v1.1.0
	github.com/phoenixnap/go-sdk-bmc/ipapi v1.0.0
	github.com/phoenixnap/go-sdk-bmc/networkapi v1.0.0
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi v1.0.0
	github.com/phoenixnap/go-sdk-bmc/tagapi v1.0.0

	github.com/spf13/viper v1.6.2
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
