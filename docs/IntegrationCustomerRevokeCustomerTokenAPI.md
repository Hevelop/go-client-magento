# \IntegrationCustomerRevokeCustomerTokenAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1IntegrationCustomerRevokecustomertoken**](IntegrationCustomerRevokeCustomerTokenAPI.md#PostV1IntegrationCustomerRevokecustomertoken) | **Post** /V1/integration/customer/revoke-customer-token | integration/customer/revoke-customer-token



## PostV1IntegrationCustomerRevokecustomertoken

> bool PostV1IntegrationCustomerRevokecustomertoken(ctx).Execute()

integration/customer/revoke-customer-token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-magento"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationCustomerRevokeCustomerTokenAPI.PostV1IntegrationCustomerRevokecustomertoken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCustomerRevokeCustomerTokenAPI.PostV1IntegrationCustomerRevokecustomertoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1IntegrationCustomerRevokecustomertoken`: bool
	fmt.Fprintf(os.Stdout, "Response from `IntegrationCustomerRevokeCustomerTokenAPI.PostV1IntegrationCustomerRevokecustomertoken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1IntegrationCustomerRevokecustomertokenRequest struct via the builder pattern


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

