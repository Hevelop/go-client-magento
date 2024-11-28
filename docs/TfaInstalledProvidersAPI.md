# \TfaInstalledProvidersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaInstalledproviders**](TfaInstalledProvidersAPI.md#GetV1TfaInstalledproviders) | **Get** /V1/tfa/installed-providers | tfa/installed-providers



## GetV1TfaInstalledproviders

> []TwoFactorAuthProviderInterface GetV1TfaInstalledproviders(ctx).Execute()

tfa/installed-providers



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
	resp, r, err := apiClient.TfaInstalledProvidersAPI.GetV1TfaInstalledproviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaInstalledProvidersAPI.GetV1TfaInstalledproviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaInstalledproviders`: []TwoFactorAuthProviderInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaInstalledProvidersAPI.GetV1TfaInstalledproviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaInstalledprovidersRequest struct via the builder pattern


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

