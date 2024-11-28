# \TfaProviderGoogleConfigureAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderGoogleConfigure**](TfaProviderGoogleConfigureAPI.md#PostV1TfaProviderGoogleConfigure) | **Post** /V1/tfa/provider/google/configure | tfa/provider/google/configure



## PostV1TfaProviderGoogleConfigure

> TwoFactorAuthDataGoogleConfigureInterface PostV1TfaProviderGoogleConfigure(ctx).PostV1TfaProviderDuoSecurityConfigureRequest(postV1TfaProviderDuoSecurityConfigureRequest).Execute()

tfa/provider/google/configure



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
	postV1TfaProviderDuoSecurityConfigureRequest := *openapiclient.NewPostV1TfaProviderDuoSecurityConfigureRequest("TfaToken_example") // PostV1TfaProviderDuoSecurityConfigureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderGoogleConfigureAPI.PostV1TfaProviderGoogleConfigure(context.Background()).PostV1TfaProviderDuoSecurityConfigureRequest(postV1TfaProviderDuoSecurityConfigureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderGoogleConfigureAPI.PostV1TfaProviderGoogleConfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderGoogleConfigure`: TwoFactorAuthDataGoogleConfigureInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderGoogleConfigureAPI.PostV1TfaProviderGoogleConfigure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderGoogleConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderDuoSecurityConfigureRequest** | [**PostV1TfaProviderDuoSecurityConfigureRequest**](PostV1TfaProviderDuoSecurityConfigureRequest.md) |  | 

### Return type

[**TwoFactorAuthDataGoogleConfigureInterface**](TwoFactorAuthDataGoogleConfigureInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

