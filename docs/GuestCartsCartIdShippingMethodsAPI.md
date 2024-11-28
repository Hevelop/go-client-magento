# \GuestCartsCartIdShippingMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdShippingmethods**](GuestCartsCartIdShippingMethodsAPI.md#GetV1GuestcartsCartIdShippingmethods) | **Get** /V1/guest-carts/{cartId}/shipping-methods | guest-carts/{cartId}/shipping-methods



## GetV1GuestcartsCartIdShippingmethods

> []QuoteDataShippingMethodInterface GetV1GuestcartsCartIdShippingmethods(ctx, cartId).Execute()

guest-carts/{cartId}/shipping-methods



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
	cartId := "cartId_example" // string | The shopping cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdShippingMethodsAPI.GetV1GuestcartsCartIdShippingmethods(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdShippingMethodsAPI.GetV1GuestcartsCartIdShippingmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdShippingmethods`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdShippingMethodsAPI.GetV1GuestcartsCartIdShippingmethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The shopping cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdShippingmethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]QuoteDataShippingMethodInterface**](QuoteDataShippingMethodInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

