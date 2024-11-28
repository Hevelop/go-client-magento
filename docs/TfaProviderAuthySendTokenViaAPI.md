# \TfaProviderAuthySendTokenViaAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderAuthySendtokenVia**](TfaProviderAuthySendTokenViaAPI.md#PostV1TfaProviderAuthySendtokenVia) | **Post** /V1/tfa/provider/authy/send-token/{via} | tfa/provider/authy/send-token/{via}



## PostV1TfaProviderAuthySendtokenVia

> ErrorResponse PostV1TfaProviderAuthySendtokenVia(ctx, via).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()

tfa/provider/authy/send-token/{via}



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
	via := "via_example" // string | 
	postV1IntegrationAdminTokenRequest := *openapiclient.NewPostV1IntegrationAdminTokenRequest("Username_example", "Password_example") // PostV1IntegrationAdminTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderAuthySendTokenViaAPI.PostV1TfaProviderAuthySendtokenVia(context.Background(), via).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderAuthySendTokenViaAPI.PostV1TfaProviderAuthySendtokenVia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderAuthySendtokenVia`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderAuthySendTokenViaAPI.PostV1TfaProviderAuthySendtokenVia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**via** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderAuthySendtokenViaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1IntegrationAdminTokenRequest** | [**PostV1IntegrationAdminTokenRequest**](PostV1IntegrationAdminTokenRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

