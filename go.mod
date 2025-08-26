module github.com/PNAP/go-sdk-helper-bmc

//replace github.com/PNAP/go-sdk-helper-bmc => D:/repos/go-sdk-helper-bmc

go 1.23.0

require (
	//github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/phoenixnap/go-sdk-bmc/auditapi/v3 v3.0.4
	github.com/phoenixnap/go-sdk-bmc/billingapi/v3 v3.0.0-20250805162518-0573cf664440
	github.com/phoenixnap/go-sdk-bmc/bmcapi/v3 v3.0.0-20250805162518-0573cf664440
	github.com/phoenixnap/go-sdk-bmc/invoicingapi v1.0.4
	github.com/phoenixnap/go-sdk-bmc/ipapi/v3 v3.0.0-20250805162518-0573cf664440
	github.com/phoenixnap/go-sdk-bmc/locationapi/v3 v3.0.2
	github.com/phoenixnap/go-sdk-bmc/networkapi/v4 v4.0.0-20250805162518-0573cf664440
	github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3 v3.0.2
	github.com/phoenixnap/go-sdk-bmc/paymentsapi v1.0.4
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v3 v3.1.1
	github.com/phoenixnap/go-sdk-bmc/tagapi/v3 v3.0.4
	github.com/spf13/viper v1.20.1
	golang.org/x/oauth2 v0.27.0
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/sagikazarmark/locafero v0.7.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.12.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
