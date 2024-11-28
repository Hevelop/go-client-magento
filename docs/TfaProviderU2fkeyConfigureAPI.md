# \TfaProviderU2fkeyConfigureAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderU2fkeyConfigure**](TfaProviderU2fkeyConfigureAPI.md#PostV1TfaProviderU2fkeyConfigure) | **Post** /V1/tfa/provider/u2fkey/configure | tfa/provider/u2fkey/configure



## PostV1TfaProviderU2fkeyConfigure

> TwoFactorAuthDataU2fWebAuthnRequestInterface PostV1TfaProviderU2fkeyConfigure(ctx).PostV1TfaProviderDuoSecurityConfigureRequest(postV1TfaProviderDuoSecurityConfigureRequest).Execute()

tfa/provider/u2fkey/configure



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
	resp, r, err := apiClient.TfaProviderU2fkeyConfigureAPI.PostV1TfaProviderU2fkeyConfigure(context.Background()).PostV1TfaProviderDuoSecurityConfigureRequest(postV1TfaProviderDuoSecurityConfigureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderU2fkeyConfigureAPI.PostV1TfaProviderU2fkeyConfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderU2fkeyConfigure`: TwoFactorAuthDataU2fWebAuthnRequestInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderU2fkeyConfigureAPI.PostV1TfaProviderU2fkeyConfigure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderU2fkeyConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderDuoSecurityConfigureRequest** | [**PostV1TfaProviderDuoSecurityConfigureRequest**](PostV1TfaProviderDuoSecurityConfigureRequest.md) |  | 

### Return type

[**TwoFactorAuthDataU2fWebAuthnRequestInterface**](TwoFactorAuthDataU2fWebAuthnRequestInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

