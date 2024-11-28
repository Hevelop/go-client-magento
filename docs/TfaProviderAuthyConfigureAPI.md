# \TfaProviderAuthyConfigureAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderAuthyConfigure**](TfaProviderAuthyConfigureAPI.md#PostV1TfaProviderAuthyConfigure) | **Post** /V1/tfa/provider/authy/configure | tfa/provider/authy/configure



## PostV1TfaProviderAuthyConfigure

> TwoFactorAuthDataAuthyRegistrationPromptResponseInterface PostV1TfaProviderAuthyConfigure(ctx).PostV1TfaProviderAuthyConfigureRequest(postV1TfaProviderAuthyConfigureRequest).Execute()

tfa/provider/authy/configure



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
	postV1TfaProviderAuthyConfigureRequest := *openapiclient.NewPostV1TfaProviderAuthyConfigureRequest("TfaToken_example", *openapiclient.NewTwoFactorAuthDataAuthyDeviceInterface("Country_example", "PhoneNumber_example", "Method_example")) // PostV1TfaProviderAuthyConfigureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderAuthyConfigureAPI.PostV1TfaProviderAuthyConfigure(context.Background()).PostV1TfaProviderAuthyConfigureRequest(postV1TfaProviderAuthyConfigureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderAuthyConfigureAPI.PostV1TfaProviderAuthyConfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderAuthyConfigure`: TwoFactorAuthDataAuthyRegistrationPromptResponseInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderAuthyConfigureAPI.PostV1TfaProviderAuthyConfigure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderAuthyConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderAuthyConfigureRequest** | [**PostV1TfaProviderAuthyConfigureRequest**](PostV1TfaProviderAuthyConfigureRequest.md) |  | 

### Return type

[**TwoFactorAuthDataAuthyRegistrationPromptResponseInterface**](TwoFactorAuthDataAuthyRegistrationPromptResponseInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

