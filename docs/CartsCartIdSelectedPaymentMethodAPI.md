# \CartsCartIdSelectedPaymentMethodAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartIdSelectedpaymentmethod**](CartsCartIdSelectedPaymentMethodAPI.md#GetV1CartsCartIdSelectedpaymentmethod) | **Get** /V1/carts/{cartId}/selected-payment-method | carts/{cartId}/selected-payment-method
[**PutV1CartsCartIdSelectedpaymentmethod**](CartsCartIdSelectedPaymentMethodAPI.md#PutV1CartsCartIdSelectedpaymentmethod) | **Put** /V1/carts/{cartId}/selected-payment-method | carts/{cartId}/selected-payment-method



## GetV1CartsCartIdSelectedpaymentmethod

> QuoteDataPaymentInterface GetV1CartsCartIdSelectedpaymentmethod(ctx, cartId).Execute()

carts/{cartId}/selected-payment-method



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
	resp, r, err := apiClient.CartsCartIdSelectedPaymentMethodAPI.GetV1CartsCartIdSelectedpaymentmethod(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdSelectedPaymentMethodAPI.GetV1CartsCartIdSelectedpaymentmethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdSelectedpaymentmethod`: QuoteDataPaymentInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdSelectedPaymentMethodAPI.GetV1CartsCartIdSelectedpaymentmethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdSelectedpaymentmethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteDataPaymentInterface**](QuoteDataPaymentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CartsCartIdSelectedpaymentmethod

> string PutV1CartsCartIdSelectedpaymentmethod(ctx, cartId).PutV1CartsMineSelectedpaymentmethodRequest(putV1CartsMineSelectedpaymentmethodRequest).Execute()

carts/{cartId}/selected-payment-method



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
	putV1CartsMineSelectedpaymentmethodRequest := *openapiclient.NewPutV1CartsMineSelectedpaymentmethodRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PutV1CartsMineSelectedpaymentmethodRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdSelectedPaymentMethodAPI.PutV1CartsCartIdSelectedpaymentmethod(context.Background(), cartId).PutV1CartsMineSelectedpaymentmethodRequest(putV1CartsMineSelectedpaymentmethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdSelectedPaymentMethodAPI.PutV1CartsCartIdSelectedpaymentmethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsCartIdSelectedpaymentmethod`: string
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdSelectedPaymentMethodAPI.PutV1CartsCartIdSelectedpaymentmethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsCartIdSelectedpaymentmethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsMineSelectedpaymentmethodRequest** | [**PutV1CartsMineSelectedpaymentmethodRequest**](PutV1CartsMineSelectedpaymentmethodRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

