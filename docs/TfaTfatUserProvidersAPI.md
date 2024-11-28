# \TfaTfatUserProvidersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaTfatuserproviders**](TfaTfatUserProvidersAPI.md#GetV1TfaTfatuserproviders) | **Get** /V1/tfa/tfat-user-providers | tfa/tfat-user-providers



## GetV1TfaTfatuserproviders

> []TwoFactorAuthProviderInterface GetV1TfaTfatuserproviders(ctx).TfaToken(tfaToken).Execute()

tfa/tfat-user-providers



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
	resp, r, err := apiClient.TfaTfatUserProvidersAPI.GetV1TfaTfatuserproviders(context.Background()).TfaToken(tfaToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaTfatUserProvidersAPI.GetV1TfaTfatuserproviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaTfatuserproviders`: []TwoFactorAuthProviderInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaTfatUserProvidersAPI.GetV1TfaTfatuserproviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaTfatuserprovidersRequest struct via the builder pattern


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

