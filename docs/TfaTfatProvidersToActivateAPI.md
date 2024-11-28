# \TfaTfatProvidersToActivateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaTfatproviderstoactivate**](TfaTfatProvidersToActivateAPI.md#GetV1TfaTfatproviderstoactivate) | **Get** /V1/tfa/tfat-providers-to-activate | tfa/tfat-providers-to-activate



## GetV1TfaTfatproviderstoactivate

> []TwoFactorAuthProviderInterface GetV1TfaTfatproviderstoactivate(ctx).TfaToken(tfaToken).Execute()

tfa/tfat-providers-to-activate



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
	tfaToken := "tfaToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaTfatProvidersToActivateAPI.GetV1TfaTfatproviderstoactivate(context.Background()).TfaToken(tfaToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaTfatProvidersToActivateAPI.GetV1TfaTfatproviderstoactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaTfatproviderstoactivate`: []TwoFactorAuthProviderInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaTfatProvidersToActivateAPI.GetV1TfaTfatproviderstoactivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaTfatproviderstoactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tfaToken** | **string** |  | 

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

