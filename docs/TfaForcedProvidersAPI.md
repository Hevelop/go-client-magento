# \TfaForcedProvidersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaForcedproviders**](TfaForcedProvidersAPI.md#GetV1TfaForcedproviders) | **Get** /V1/tfa/forced-providers | tfa/forced-providers



## GetV1TfaForcedproviders

> []TwoFactorAuthProviderInterface GetV1TfaForcedproviders(ctx).Execute()

tfa/forced-providers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaForcedProvidersAPI.GetV1TfaForcedproviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaForcedProvidersAPI.GetV1TfaForcedproviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaForcedproviders`: []TwoFactorAuthProviderInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaForcedProvidersAPI.GetV1TfaForcedproviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaForcedprovidersRequest struct via the builder pattern


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

