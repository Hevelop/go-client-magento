# \TfaProviderDuoSecurityConfigureAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderDuoSecurityConfigure**](TfaProviderDuoSecurityConfigureAPI.md#PostV1TfaProviderDuoSecurityConfigure) | **Post** /V1/tfa/provider/duo_security/configure | tfa/provider/duo_security/configure



## PostV1TfaProviderDuoSecurityConfigure

> TwoFactorAuthDataDuoDataInterface PostV1TfaProviderDuoSecurityConfigure(ctx).PostV1TfaProviderDuoSecurityConfigureRequest(postV1TfaProviderDuoSecurityConfigureRequest).Execute()

tfa/provider/duo_security/configure



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
	resp, r, err := apiClient.TfaProviderDuoSecurityConfigureAPI.PostV1TfaProviderDuoSecurityConfigure(context.Background()).PostV1TfaProviderDuoSecurityConfigureRequest(postV1TfaProviderDuoSecurityConfigureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderDuoSecurityConfigureAPI.PostV1TfaProviderDuoSecurityConfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderDuoSecurityConfigure`: TwoFactorAuthDataDuoDataInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderDuoSecurityConfigureAPI.PostV1TfaProviderDuoSecurityConfigure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderDuoSecurityConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderDuoSecurityConfigureRequest** | [**PostV1TfaProviderDuoSecurityConfigureRequest**](PostV1TfaProviderDuoSecurityConfigureRequest.md) |  | 

### Return type

[**TwoFactorAuthDataDuoDataInterface**](TwoFactorAuthDataDuoDataInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

