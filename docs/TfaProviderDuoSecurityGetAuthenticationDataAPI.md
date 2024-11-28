# \TfaProviderDuoSecurityGetAuthenticationDataAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderDuoSecurityGetauthenticationdata**](TfaProviderDuoSecurityGetAuthenticationDataAPI.md#PostV1TfaProviderDuoSecurityGetauthenticationdata) | **Post** /V1/tfa/provider/duo_security/get-authentication-data | tfa/provider/duo_security/get-authentication-data



## PostV1TfaProviderDuoSecurityGetauthenticationdata

> TwoFactorAuthDataDuoDataInterface PostV1TfaProviderDuoSecurityGetauthenticationdata(ctx).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()

tfa/provider/duo_security/get-authentication-data



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
	resp, r, err := apiClient.TfaProviderDuoSecurityGetAuthenticationDataAPI.PostV1TfaProviderDuoSecurityGetauthenticationdata(context.Background()).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderDuoSecurityGetAuthenticationDataAPI.PostV1TfaProviderDuoSecurityGetauthenticationdata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderDuoSecurityGetauthenticationdata`: TwoFactorAuthDataDuoDataInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderDuoSecurityGetAuthenticationDataAPI.PostV1TfaProviderDuoSecurityGetauthenticationdata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderDuoSecurityGetauthenticationdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1IntegrationAdminTokenRequest** | [**PostV1IntegrationAdminTokenRequest**](PostV1IntegrationAdminTokenRequest.md) |  | 

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

