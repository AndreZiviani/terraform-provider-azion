package provider

import (
	"github.com/aziontech/azionapi-go-sdk/domains"
	"github.com/aziontech/azionapi-go-sdk/idns"
)

type apiClient struct {
	idnsConfig *idns.Configuration
	idnsApi    *idns.APIClient

	domainsConfig *domains.Configuration
	domainsApi    *domains.APIClient
}

func Client(APIToken string) *apiClient {
	client := &apiClient{
		idnsConfig:    idns.NewConfiguration(),
		domainsConfig: domains.NewConfiguration(),
	}

	client.domainsApi = domains.NewAPIClient(client.domainsConfig)
	client.domainsConfig.AddDefaultHeader("Authorization", "token "+APIToken)

	client.idnsApi = idns.NewAPIClient(client.idnsConfig)
	client.idnsConfig.AddDefaultHeader("Authorization", "token "+APIToken)

	return client
}
