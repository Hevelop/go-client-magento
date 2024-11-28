# \CartsCartIdShippingMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartIdShippingmethods**](CartsCartIdShippingMethodsAPI.md#GetV1CartsCartIdShippingmethods) | **Get** /V1/carts/{cartId}/shipping-methods | carts/{cartId}/shipping-methods



## GetV1CartsCartIdShippingmethods

> []QuoteDataShippingMethodInterface GetV1CartsCartIdShippingmethods(ctx, cartId).Execute()

carts/{cartId}/shipping-methods



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
	cartId := int32(56) // int32 | The shopping cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdShippingMethodsAPI.GetV1CartsCartIdShippingmethods(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdShippingMethodsAPI.GetV1CartsCartIdShippingmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdShippingmethods`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdShippingMethodsAPI.GetV1CartsCartIdShippingmethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The shopping cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdShippingmethodsRequest struct via the builder pattern


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

