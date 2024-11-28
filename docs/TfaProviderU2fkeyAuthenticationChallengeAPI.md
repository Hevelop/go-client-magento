# \TfaProviderU2fkeyAuthenticationChallengeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderU2fkeyAuthenticationchallenge**](TfaProviderU2fkeyAuthenticationChallengeAPI.md#PostV1TfaProviderU2fkeyAuthenticationchallenge) | **Post** /V1/tfa/provider/u2fkey/authentication-challenge | tfa/provider/u2fkey/authentication-challenge



## PostV1TfaProviderU2fkeyAuthenticationchallenge

> TwoFactorAuthDataU2fWebAuthnRequestInterface PostV1TfaProviderU2fkeyAuthenticationchallenge(ctx).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()

tfa/provider/u2fkey/authentication-challenge



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
	resp, r, err := apiClient.TfaProviderU2fkeyAuthenticationChallengeAPI.PostV1TfaProviderU2fkeyAuthenticationchallenge(context.Background()).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderU2fkeyAuthenticationChallengeAPI.PostV1TfaProviderU2fkeyAuthenticationchallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderU2fkeyAuthenticationchallenge`: TwoFactorAuthDataU2fWebAuthnRequestInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderU2fkeyAuthenticationChallengeAPI.PostV1TfaProviderU2fkeyAuthenticationchallenge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderU2fkeyAuthenticationchallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1IntegrationAdminTokenRequest** | [**PostV1IntegrationAdminTokenRequest**](PostV1IntegrationAdminTokenRequest.md) |  | 

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

