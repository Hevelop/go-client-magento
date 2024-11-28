# \TfaProviderAuthyActivateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderAuthyActivate**](TfaProviderAuthyActivateAPI.md#PostV1TfaProviderAuthyActivate) | **Post** /V1/tfa/provider/authy/activate | tfa/provider/authy/activate



## PostV1TfaProviderAuthyActivate

> ErrorResponse PostV1TfaProviderAuthyActivate(ctx).PostV1TfaProviderAuthyActivateRequest(postV1TfaProviderAuthyActivateRequest).Execute()

tfa/provider/authy/activate



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
	postV1TfaProviderAuthyActivateRequest := *openapiclient.NewPostV1TfaProviderAuthyActivateRequest("TfaToken_example", "Otp_example") // PostV1TfaProviderAuthyActivateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderAuthyActivateAPI.PostV1TfaProviderAuthyActivate(context.Background()).PostV1TfaProviderAuthyActivateRequest(postV1TfaProviderAuthyActivateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderAuthyActivateAPI.PostV1TfaProviderAuthyActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderAuthyActivate`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderAuthyActivateAPI.PostV1TfaProviderAuthyActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderAuthyActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderAuthyActivateRequest** | [**PostV1TfaProviderAuthyActivateRequest**](PostV1TfaProviderAuthyActivateRequest.md) |  | 

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

