# \CartsCartIdPaymentMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartIdPaymentmethods**](CartsCartIdPaymentMethodsAPI.md#GetV1CartsCartIdPaymentmethods) | **Get** /V1/carts/{cartId}/payment-methods | carts/{cartId}/payment-methods



## GetV1CartsCartIdPaymentmethods

> []QuoteDataPaymentMethodInterface GetV1CartsCartIdPaymentmethods(ctx, cartId).Execute()

carts/{cartId}/payment-methods



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
	cartId := int32(56) // int32 | The cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdPaymentMethodsAPI.GetV1CartsCartIdPaymentmethods(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdPaymentMethodsAPI.GetV1CartsCartIdPaymentmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdPaymentmethods`: []QuoteDataPaymentMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdPaymentMethodsAPI.GetV1CartsCartIdPaymentmethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdPaymentmethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]QuoteDataPaymentMethodInterface**](QuoteDataPaymentMethodInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
