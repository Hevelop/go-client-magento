# \CustomersMeShippingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersMeShippingAddress**](CustomersMeShippingAddressAPI.md#GetV1CustomersMeShippingAddress) | **Get** /V1/customers/me/shippingAddress | customers/me/shippingAddress



## GetV1CustomersMeShippingAddress

> CustomerDataAddressInterface GetV1CustomersMeShippingAddress(ctx).Execute()

customers/me/shippingAddress



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
	resp, r, err := apiClient.CustomersMeShippingAddressAPI.GetV1CustomersMeShippingAddress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersMeShippingAddressAPI.GetV1CustomersMeShippingAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersMeShippingAddress`: CustomerDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersMeShippingAddressAPI.GetV1CustomersMeShippingAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersMeShippingAddressRequest struct via the builder pattern


### Return type

[**CustomerDataAddressInterface**](CustomerDataAddressInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

