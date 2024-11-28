# \ApplepayAuthAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ApplepayAuth**](ApplepayAuthAPI.md#GetV1ApplepayAuth) | **Get** /V1/applepay/auth | applepay/auth



## GetV1ApplepayAuth

> PayPalBraintreeDataAuthDataInterface GetV1ApplepayAuth(ctx).Execute()

applepay/auth



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
	resp, r, err := apiClient.ApplepayAuthAPI.GetV1ApplepayAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplepayAuthAPI.GetV1ApplepayAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ApplepayAuth`: PayPalBraintreeDataAuthDataInterface
	fmt.Fprintf(os.Stdout, "Response from `ApplepayAuthAPI.GetV1ApplepayAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ApplepayAuthRequest struct via the builder pattern


### Return type

[**PayPalBraintreeDataAuthDataInterface**](PayPalBraintreeDataAuthDataInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

