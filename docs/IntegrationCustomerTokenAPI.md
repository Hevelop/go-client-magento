# \IntegrationCustomerTokenAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1IntegrationCustomerToken**](IntegrationCustomerTokenAPI.md#PostV1IntegrationCustomerToken) | **Post** /V1/integration/customer/token | integration/customer/token



## PostV1IntegrationCustomerToken

> string PostV1IntegrationCustomerToken(ctx).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()

integration/customer/token



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
	postV1IntegrationAdminTokenRequest := *openapiclient.NewPostV1IntegrationAdminTokenRequest("Username_example", "Password_example") // PostV1IntegrationAdminTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationCustomerTokenAPI.PostV1IntegrationCustomerToken(context.Background()).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCustomerTokenAPI.PostV1IntegrationCustomerToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1IntegrationCustomerToken`: string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationCustomerTokenAPI.PostV1IntegrationCustomerToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1IntegrationCustomerTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1IntegrationAdminTokenRequest** | [**PostV1IntegrationAdminTokenRequest**](PostV1IntegrationAdminTokenRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

