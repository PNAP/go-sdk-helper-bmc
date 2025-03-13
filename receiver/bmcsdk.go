package receiver

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/config"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"golang.org/x/oauth2/clientcredentials"

	"net/http"

	"github.com/mitchellh/go-homedir"
	auditapiclient "github.com/phoenixnap/go-sdk-bmc/auditapi/v3"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	invoicingapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	locationapiclient "github.com/phoenixnap/go-sdk-bmc/locationapi/v3"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	paymentsapiclient "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v3"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi/v3"
	"github.com/spf13/viper"
)

// BMCSDK is a Client that performs HTTP requests.
type BMCSDK struct {
	APIClient               bmcapiclient.APIClient
	RancherAPIClient        rancherapiclient.APIClient
	NetworkAPIClient        networkapiclient.APIClient
	TagAPIClient            tagapiclient.APIClient
	AuditAPIClient          auditapiclient.APIClient
	IpBlockAPIClient        ipapiclient.APIClient
	NetworkStorageAPIClient networkstorageapiclient.APIClient
	BillingAPIClient        billingapiclient.APIClient
	LocationAPIClient       locationapiclient.APIClient
	InvoicingAPIClient      invoicingapiclient.APIClient
	PaymentsAPIClient       paymentsapiclient.APIClient
	PNAPClient              PNAPClient
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
	ipApiConfiguration := ipapiclient.NewConfiguration()
	networkStorageApiConfiguration := networkstorageapiclient.NewConfiguration()
	billingApiConfiguration := billingapiclient.NewConfiguration()
	locationApiConfiguration := locationapiclient.NewConfiguration()
	invoicingApiConfiguration := invoicingapiclient.NewConfiguration()
	paymentsApiConfiguration := paymentsapiclient.NewConfiguration()

	bmcApiConfiguration.HTTPClient = config.Client(context.Background())
	rancherApiConfiguration.HTTPClient = config.Client(context.Background())
	networkApiConfiguration.HTTPClient = config.Client(context.Background())
	tagApiConfiguration.HTTPClient = config.Client(context.Background())
	auditApiConfiguration.HTTPClient = config.Client(context.Background())
	ipApiConfiguration.HTTPClient = config.Client(context.Background())
	networkStorageApiConfiguration.HTTPClient = config.Client(context.Background())
	billingApiConfiguration.HTTPClient = config.Client(context.Background())
	locationApiConfiguration.HTTPClient = config.Client(context.Background())
	invoicingApiConfiguration.HTTPClient = config.Client(context.Background())
	paymentsApiConfiguration.HTTPClient = config.Client(context.Background())

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
		ipApiConfiguration.UserAgent = auth.UserAgent
		networkStorageApiConfiguration.UserAgent = auth.UserAgent
		billingApiConfiguration.UserAgent = auth.UserAgent
		locationApiConfiguration.UserAgent = auth.UserAgent
		invoicingApiConfiguration.UserAgent = auth.UserAgent
		paymentsApiConfiguration.UserAgent = auth.UserAgent
	}

	if auth.PoweredBy != "" {
		bmcApiConfiguration.XPoweredBy = auth.PoweredBy
		rancherApiConfiguration.XPoweredBy = auth.PoweredBy
		networkApiConfiguration.XPoweredBy = auth.PoweredBy
		tagApiConfiguration.XPoweredBy = auth.PoweredBy
		auditApiConfiguration.XPoweredBy = auth.PoweredBy
		ipApiConfiguration.XPoweredBy = auth.PoweredBy
		networkStorageApiConfiguration.XPoweredBy = auth.PoweredBy
		billingApiConfiguration.XPoweredBy = auth.PoweredBy
		locationApiConfiguration.XPoweredBy = auth.PoweredBy
		invoicingApiConfiguration.XPoweredBy = auth.PoweredBy
		paymentsApiConfiguration.XPoweredBy = auth.PoweredBy
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)
	ipApiClient := ipapiclient.NewAPIClient(ipApiConfiguration)
	networkStorageApiClient := networkstorageapiclient.NewAPIClient(networkStorageApiConfiguration)
	billingApiClient := billingapiclient.NewAPIClient(billingApiConfiguration)
	locationApiClient := locationapiclient.NewAPIClient(locationApiConfiguration)
	invoicingApiClient := invoicingapiclient.NewAPIClient(invoicingApiConfiguration)
	paymentsApiClient := paymentsapiclient.NewAPIClient(paymentsApiConfiguration)
	pnapClient, err := NewPNAPClientWithDefaultConfig()
	pnapClient.SetAuthentication(auth)

	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient, *ipApiClient, *networkStorageApiClient, *billingApiClient,
		*locationApiClient, *invoicingApiClient, *paymentsApiClient, pnapClient}
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
	ipApiConfiguration := ipapiclient.NewConfiguration()
	networkStorageApiConfiguration := networkstorageapiclient.NewConfiguration()
	billingApiConfiguration := billingapiclient.NewConfiguration()
	locationApiConfiguration := locationapiclient.NewConfiguration()
	invoicingApiConfiguration := invoicingapiclient.NewConfiguration()
	paymentsApiConfiguration := paymentsapiclient.NewConfiguration()

	if auth.ApiHostName != "" {
		bmcApiConfiguration.Servers = bmcapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "bmc/v1",
			},
		}
		rancherApiConfiguration.Servers = rancherapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "solutions/rancher/v1beta",
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
		ipApiConfiguration.Servers = ipapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "ips/v1",
			},
		}
		networkStorageApiConfiguration.Servers = networkstorageapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "network-storage/v1",
			},
		}
		billingApiConfiguration.Servers = billingapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "billing/v1",
			},
		}
		locationApiConfiguration.Servers = locationapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "location-api/v1",
			},
		}
		invoicingApiConfiguration.Servers = invoicingapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "invoicing/v1",
			},
		}
		paymentsApiConfiguration.Servers = paymentsapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "payments/v1",
			},
		}
	}

	bmcApiConfiguration.HTTPClient = config.Client(context.Background())
	rancherApiConfiguration.HTTPClient = config.Client(context.Background())
	networkApiConfiguration.HTTPClient = config.Client(context.Background())
	tagApiConfiguration.HTTPClient = config.Client(context.Background())
	auditApiConfiguration.HTTPClient = config.Client(context.Background())
	ipApiConfiguration.HTTPClient = config.Client(context.Background())
	networkStorageApiConfiguration.HTTPClient = config.Client(context.Background())
	billingApiConfiguration.HTTPClient = config.Client(context.Background())
	locationApiConfiguration.HTTPClient = config.Client(context.Background())
	invoicingApiConfiguration.HTTPClient = config.Client(context.Background())
	paymentsApiConfiguration.HTTPClient = config.Client(context.Background())

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
		ipApiConfiguration.UserAgent = auth.UserAgent
		networkStorageApiConfiguration.UserAgent = auth.UserAgent
		billingApiConfiguration.UserAgent = auth.UserAgent
		locationApiConfiguration.UserAgent = auth.UserAgent
		invoicingApiConfiguration.UserAgent = auth.UserAgent
		paymentsApiConfiguration.UserAgent = auth.UserAgent
	}

	if auth.PoweredBy != "" {
		bmcApiConfiguration.XPoweredBy = auth.PoweredBy
		rancherApiConfiguration.XPoweredBy = auth.PoweredBy
		networkApiConfiguration.XPoweredBy = auth.PoweredBy
		tagApiConfiguration.XPoweredBy = auth.PoweredBy
		auditApiConfiguration.XPoweredBy = auth.PoweredBy
		ipApiConfiguration.XPoweredBy = auth.PoweredBy
		networkStorageApiConfiguration.XPoweredBy = auth.PoweredBy
		billingApiConfiguration.XPoweredBy = auth.PoweredBy
		locationApiConfiguration.XPoweredBy = auth.PoweredBy
		invoicingApiConfiguration.XPoweredBy = auth.PoweredBy
		paymentsApiConfiguration.XPoweredBy = auth.PoweredBy
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)
	ipApiClient := ipapiclient.NewAPIClient(ipApiConfiguration)
	networkStorageApiClient := networkstorageapiclient.NewAPIClient(networkStorageApiConfiguration)
	billingApiClient := billingapiclient.NewAPIClient(billingApiConfiguration)
	locationApiClient := locationapiclient.NewAPIClient(locationApiConfiguration)
	invoicingApiClient := invoicingapiclient.NewAPIClient(invoicingApiConfiguration)
	paymentsApiClient := paymentsapiclient.NewAPIClient(paymentsApiConfiguration)

	pnapClient := NewPNAPClient(auth)

	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient, *ipApiClient, *networkStorageApiClient, *billingApiClient,
		*locationApiClient, *invoicingApiClient, *paymentsApiClient, pnapClient}
	return sdkClient
}

// NewBMCSDKWithTokenAuthentication creates a new BMCSDK receiver with Bearer token directly set to the header
func NewBMCSDKWithTokenAuthentication(auth dto.Configuration) BMCSDK {

	//tokenUrl := config.TokenURL
	/* if auth.TokenURL != "" {
		tokenUrl = auth.TokenURL
	} */
	/* config := clientcredentials.Config{
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		TokenURL:     tokenUrl,
		Scopes:       []string{"bmc", "bmc.read"},
	} */

	bmcApiConfiguration := bmcapiclient.NewConfiguration()
	rancherApiConfiguration := rancherapiclient.NewConfiguration()
	networkApiConfiguration := networkapiclient.NewConfiguration()
	tagApiConfiguration := tagapiclient.NewConfiguration()
	auditApiConfiguration := auditapiclient.NewConfiguration()
	ipApiConfiguration := ipapiclient.NewConfiguration()
	networkStorageApiConfiguration := networkstorageapiclient.NewConfiguration()
	billingApiConfiguration := billingapiclient.NewConfiguration()
	locationApiConfiguration := locationapiclient.NewConfiguration()
	invoicingApiConfiguration := invoicingapiclient.NewConfiguration()
	paymentsApiConfiguration := paymentsapiclient.NewConfiguration()

	if auth.ApiHostName != "" {
		bmcApiConfiguration.Servers = bmcapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "bmc/v1",
			},
		}
		rancherApiConfiguration.Servers = rancherapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "solutions/rancher/v1beta",
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
		ipApiConfiguration.Servers = ipapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "ips/v1",
			},
		}
		networkStorageApiConfiguration.Servers = networkstorageapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "network-storage/v1",
			},
		}
		billingApiConfiguration.Servers = billingapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "billing/v1",
			},
		}
		locationApiConfiguration.Servers = locationapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "location-api/v1",
			},
		}
		invoicingApiConfiguration.Servers = invoicingapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "invoicing/v1",
			},
		}
		paymentsApiConfiguration.Servers = paymentsapiclient.ServerConfigurations{
			{
				URL: auth.ApiHostName + "payments/v1",
			},
		}
	}

	bmcApiConfiguration.HTTPClient = &http.Client{}
	rancherApiConfiguration.HTTPClient = &http.Client{}
	networkApiConfiguration.HTTPClient = &http.Client{}
	tagApiConfiguration.HTTPClient = &http.Client{}
	auditApiConfiguration.HTTPClient = &http.Client{}
	ipApiConfiguration.HTTPClient = &http.Client{}
	networkStorageApiConfiguration.HTTPClient = &http.Client{}
	billingApiConfiguration.HTTPClient = &http.Client{}
	locationApiConfiguration.HTTPClient = &http.Client{}
	invoicingApiConfiguration.HTTPClient = &http.Client{}
	paymentsApiConfiguration.HTTPClient = &http.Client{}

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
		ipApiConfiguration.UserAgent = auth.UserAgent
		networkStorageApiConfiguration.UserAgent = auth.UserAgent
		billingApiConfiguration.UserAgent = auth.UserAgent
		locationApiConfiguration.UserAgent = auth.UserAgent
		invoicingApiConfiguration.UserAgent = auth.UserAgent
		paymentsApiConfiguration.UserAgent = auth.UserAgent
	}

	if auth.BearerToken != "" {
		bmcApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		rancherApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		networkApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		tagApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		auditApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		ipApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		networkStorageApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		billingApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		locationApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		invoicingApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
		paymentsApiConfiguration.AddDefaultHeader("Authorization", "Bearer "+auth.BearerToken)
	}

	if auth.PoweredBy != "" {
		bmcApiConfiguration.XPoweredBy = auth.PoweredBy
		rancherApiConfiguration.XPoweredBy = auth.PoweredBy
		networkApiConfiguration.XPoweredBy = auth.PoweredBy
		tagApiConfiguration.XPoweredBy = auth.PoweredBy
		auditApiConfiguration.XPoweredBy = auth.PoweredBy
		ipApiConfiguration.XPoweredBy = auth.PoweredBy
		networkStorageApiConfiguration.XPoweredBy = auth.PoweredBy
		billingApiConfiguration.XPoweredBy = auth.PoweredBy
		locationApiConfiguration.XPoweredBy = auth.PoweredBy
		invoicingApiConfiguration.XPoweredBy = auth.PoweredBy
		paymentsApiConfiguration.XPoweredBy = auth.PoweredBy
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)
	ipApiClient := ipapiclient.NewAPIClient(ipApiConfiguration)
	networkStorageApiClient := networkstorageapiclient.NewAPIClient(networkStorageApiConfiguration)
	billingApiClient := billingapiclient.NewAPIClient(billingApiConfiguration)
	locationApiClient := locationapiclient.NewAPIClient(locationApiConfiguration)
	invoicingApiClient := invoicingapiclient.NewAPIClient(invoicingApiConfiguration)
	paymentsApiClient := paymentsapiclient.NewAPIClient(paymentsApiConfiguration)

	pnapClient := NewPNAPClientWithTokenAuthentication(auth)

	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient, *ipApiClient, *networkStorageApiClient, *billingApiClient,
		*locationApiClient, *invoicingApiClient, *paymentsApiClient, pnapClient}
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
	ipApiConfiguration := ipapiclient.NewConfiguration()
	networkStorageApiConfiguration := networkstorageapiclient.NewConfiguration()
	billingApiConfiguration := billingapiclient.NewConfiguration()
	locationApiConfiguration := locationapiclient.NewConfiguration()
	invoicingApiConfiguration := invoicingapiclient.NewConfiguration()
	paymentsApiConfiguration := paymentsapiclient.NewConfiguration()

	bmcApiConfiguration.HTTPClient = config.Client(context.Background())

	rancherApiConfiguration.HTTPClient = config.Client(context.Background())
	networkApiConfiguration.HTTPClient = config.Client(context.Background())
	tagApiConfiguration.HTTPClient = config.Client(context.Background())
	auditApiConfiguration.HTTPClient = config.Client(context.Background())
	ipApiConfiguration.HTTPClient = config.Client(context.Background())
	networkStorageApiConfiguration.HTTPClient = config.Client(context.Background())
	billingApiConfiguration.HTTPClient = config.Client(context.Background())
	locationApiConfiguration.HTTPClient = config.Client(context.Background())
	invoicingApiConfiguration.HTTPClient = config.Client(context.Background())
	paymentsApiConfiguration.HTTPClient = config.Client(context.Background())

	if auth.UserAgent != "" {
		bmcApiConfiguration.UserAgent = auth.UserAgent
		rancherApiConfiguration.UserAgent = auth.UserAgent
		networkApiConfiguration.UserAgent = auth.UserAgent
		tagApiConfiguration.UserAgent = auth.UserAgent
		auditApiConfiguration.UserAgent = auth.UserAgent
		ipApiConfiguration.UserAgent = auth.UserAgent
		networkStorageApiConfiguration.UserAgent = auth.UserAgent
		billingApiConfiguration.UserAgent = auth.UserAgent
		locationApiConfiguration.UserAgent = auth.UserAgent
		invoicingApiConfiguration.UserAgent = auth.UserAgent
		paymentsApiConfiguration.UserAgent = auth.UserAgent
	}
	if auth.PoweredBy != "" {
		bmcApiConfiguration.XPoweredBy = auth.PoweredBy
		rancherApiConfiguration.XPoweredBy = auth.PoweredBy
		networkApiConfiguration.XPoweredBy = auth.PoweredBy
		tagApiConfiguration.XPoweredBy = auth.PoweredBy
		auditApiConfiguration.XPoweredBy = auth.PoweredBy
		ipApiConfiguration.XPoweredBy = auth.PoweredBy
		networkStorageApiConfiguration.XPoweredBy = auth.PoweredBy
		billingApiConfiguration.XPoweredBy = auth.PoweredBy
		locationApiConfiguration.XPoweredBy = auth.PoweredBy
		invoicingApiConfiguration.XPoweredBy = auth.PoweredBy
		paymentsApiConfiguration.XPoweredBy = auth.PoweredBy
	}

	bmcApiClient := bmcapiclient.NewAPIClient(bmcApiConfiguration)
	rancherApiClient := rancherapiclient.NewAPIClient(rancherApiConfiguration)
	networkApiClient := networkapiclient.NewAPIClient(networkApiConfiguration)
	tagApiClient := tagapiclient.NewAPIClient(tagApiConfiguration)
	auditApiClient := auditapiclient.NewAPIClient(auditApiConfiguration)
	ipApiClient := ipapiclient.NewAPIClient(ipApiConfiguration)
	networkStorageApiClient := networkstorageapiclient.NewAPIClient(networkStorageApiConfiguration)
	billingApiClient := billingapiclient.NewAPIClient(billingApiConfiguration)
	locationApiClient := locationapiclient.NewAPIClient(locationApiConfiguration)
	invoicingApiClient := invoicingapiclient.NewAPIClient(invoicingApiConfiguration)
	paymentsApiClient := paymentsapiclient.NewAPIClient(paymentsApiConfiguration)

	pnapClient, err := NewPNAPClientWithCustomConfig(path)
	pnapClient.SetAuthentication(auth)
	sdkClient := BMCSDK{*bmcApiClient, *rancherApiClient, *networkApiClient, *tagApiClient, *auditApiClient, *ipApiClient, *networkStorageApiClient, *billingApiClient,
		*locationApiClient, *invoicingApiClient, *paymentsApiClient, pnapClient}
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
