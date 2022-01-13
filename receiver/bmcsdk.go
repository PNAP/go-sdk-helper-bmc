package receiver

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/config"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/mitchellh/go-homedir"
	auditapiclient "github.com/phoenixnap/go-sdk-bmc/auditapi"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/spf13/viper"
)

// BMCSDK is a Client that performs HTTP requests.
type BMCSDK struct {
	APIClient        bmcapiclient.APIClient
	RancherAPIClient rancherapiclient.APIClient
	NetworkAPIClient networkapiclient.APIClient
	TagAPIClient     tagapiclient.APIClient
	AuditAPIClient   auditapiclient.APIClient
}

//NewBMCSDKWithDefaultConfig creates a new BMCSDK receiver with credentials from config file on default path. Verification of configuration file will be done prior to creation
//and error will be returned in case credentials or whole configuration file is missing
func NewBMCSDKWithDefaultConfig(auth dto.Configuration) (BMCSDK, error) {
	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		return BMCSDK{}, err
	}

	configPath := home + config.DefaultConfigPath
	confPathErr := Verify(configPath)
	if confPathErr != nil {
		return BMCSDK{}, confPathErr
	}

	config := load(configPath)
	bmcApiConfiguration := bmcapiclient.NewConfiguration()
	rancherApiConfiguration := rancherapiclient.NewConfiguration()
	networkApiConfiguration := networkapiclient.NewConfiguration()
	tagApiConfiguration := tagapiclient.NewConfiguration()
	auditApiConfiguration := auditapiclient.NewConfiguration()

	bmcApiConfiguration.HTTPClient = config.Client(context.Background())
	rancherApiConfiguration.HTTPClient = config.Client(context.Background())
	networkApiConfiguration.HTTPClient = config.Client(context.Background())
	tagApiConfiguration.HTTPClient = config.Client(context.Background())
	auditApiConfiguration.HTTPClient = config.Client(context.Background())

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)

	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient}
	return sdkClient, err
}

// NewBMCSDK creates a new BMCSDK receiver with credentials from auth object
func NewBMCSDK(auth dto.Configuration) BMCSDK {

	tokenUrl := config.TokenURL
	if auth.TokenURL != "" {
		tokenUrl = auth.TokenURL
	}
	config := clientcredentials.Config{
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		TokenURL:     tokenUrl,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	bmcApiConfiguration := bmcapiclient.NewConfiguration()
	rancherApiConfiguration := rancherapiclient.NewConfiguration()
	networkApiConfiguration := networkapiclient.NewConfiguration()
	tagApiConfiguration := tagapiclient.NewConfiguration()
	auditApiConfiguration := auditapiclient.NewConfiguration()

	if auth.ApiHostName != "" {
		bmcApiConfiguration.Servers = bmcapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "bmc/v0",
			},
		}
		rancherApiConfiguration.Servers = rancherapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "solutions/rancher/v0",
			},
		}
		networkApiConfiguration.Servers = networkapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "networks/v1",
			},
		}
		tagApiConfiguration.Servers = tagapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "tag-manager/v1",
			},
		}
		auditApiConfiguration.Servers = auditapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "audit/v1",
			},
		}
	}

	bmcApiConfiguration.HTTPClient = config.Client(context.Background())
	rancherApiConfiguration.HTTPClient = config.Client(context.Background())
	networkApiConfiguration.HTTPClient = config.Client(context.Background())
	tagApiConfiguration.HTTPClient = config.Client(context.Background())
	auditApiConfiguration.HTTPClient = config.Client(context.Background())

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)

	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient}
	return sdkClient
}

//NewBMCSDKWithCustomConfig creates a new BMCSDK receiver with credentials from configuration file on custom path. Verification of configuration file will be done prior to creation
//and error will be returned in case credentials or whole configuration file is missing
func NewBMCSDKWithCustomConfig(path string, auth dto.Configuration) (BMCSDK, error) {
	err := Verify(path)
	if err != nil {
		return BMCSDK{}, err
	}
	config := load(path)
	bmcApiConfiguration := bmcapiclient.NewConfiguration()
	rancherApiConfiguration := rancherapiclient.NewConfiguration()
	networkApiConfiguration := networkapiclient.NewConfiguration()
	tagApiConfiguration := tagapiclient.NewConfiguration()
	auditApiConfiguration := auditapiclient.NewConfiguration()

	bmcApiConfiguration.HTTPClient = config.Client(context.Background())

	rancherApiConfiguration.HTTPClient = config.Client(context.Background())
	networkApiConfiguration.HTTPClient = config.Client(context.Background())
	tagApiConfiguration.HTTPClient = config.Client(context.Background())
	auditApiConfiguration.HTTPClient = config.Client(context.Background())

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)

	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient}
	return sdkClient, err
}

func load(configPath string) clientcredentials.Config {

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.ReadInConfig()
	clientID := viper.GetString("clientId")
	clientSecret := viper.GetString("clientSecret")
	tokenURL := config.TokenURL

	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}
	return config
}

//Verify verifies existence of configuration file and credentials
func Verify(configPath string) error {

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Checks whether the config file exists, by attempting to cast the error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("A config file is required in order to proceed.\n" +
				"Config file path is (" + configPath + "config.yaml)\n\n" +
				"The following shows a sample config file:\n\n" +
				"# =====================================================\n" +
				"# Sample yaml config file\n" +
				"# =====================================================\n\n" +
				"# Authentication\n" +
				"clientId: <enter your client id>\n" +
				"clientSecret: <enter your client secret>")
		}
		return err
	}
	clientID := viper.GetString("clientId")
	if clientID == "" {
		return fmt.Errorf("API client ID not found in configuration")
	}
	clientSecret := viper.GetString("clientSecret")
	if clientSecret == "" {
		return fmt.Errorf("API clientSecret not found in configuration")
	}
	return nil
}
