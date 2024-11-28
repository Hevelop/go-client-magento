# \TfaProvidersToActivateUserIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaProviderstoactivateUserId**](TfaProvidersToActivateUserIdAPI.md#GetV1TfaProviderstoactivateUserId) | **Get** /V1/tfa/providers-to-activate/{userId} | tfa/providers-to-activate/{userId}



## GetV1TfaProviderstoactivateUserId

> []TwoFactorAuthProviderInterface GetV1TfaProviderstoactivateUserId(ctx, userId).Execute()

tfa/providers-to-activate/{userId}



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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProvidersToActivateUserIdAPI.GetV1TfaProviderstoactivateUserId(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProvidersToActivateUserIdAPI.GetV1TfaProviderstoactivateUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaProviderstoactivateUserId`: []TwoFactorAuthProviderInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaProvidersToActivateUserIdAPI.GetV1TfaProviderstoactivateUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaProviderstoactivateUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TwoFactorAuthProviderInterface**](TwoFactorAuthProviderInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

