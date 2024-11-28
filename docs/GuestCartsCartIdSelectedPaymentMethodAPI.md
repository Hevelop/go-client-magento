# \GuestCartsCartIdSelectedPaymentMethodAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdSelectedpaymentmethod**](GuestCartsCartIdSelectedPaymentMethodAPI.md#GetV1GuestcartsCartIdSelectedpaymentmethod) | **Get** /V1/guest-carts/{cartId}/selected-payment-method | guest-carts/{cartId}/selected-payment-method
[**PutV1GuestcartsCartIdSelectedpaymentmethod**](GuestCartsCartIdSelectedPaymentMethodAPI.md#PutV1GuestcartsCartIdSelectedpaymentmethod) | **Put** /V1/guest-carts/{cartId}/selected-payment-method | guest-carts/{cartId}/selected-payment-method



## GetV1GuestcartsCartIdSelectedpaymentmethod

> QuoteDataPaymentInterface GetV1GuestcartsCartIdSelectedpaymentmethod(ctx, cartId).Execute()

guest-carts/{cartId}/selected-payment-method



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
	cartId := "cartId_example" // string | The cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdSelectedPaymentMethodAPI.GetV1GuestcartsCartIdSelectedpaymentmethod(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdSelectedPaymentMethodAPI.GetV1GuestcartsCartIdSelectedpaymentmethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdSelectedpaymentmethod`: QuoteDataPaymentInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdSelectedPaymentMethodAPI.GetV1GuestcartsCartIdSelectedpaymentmethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdSelectedpaymentmethodRequest struct via the builder pattern


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


## PutV1GuestcartsCartIdSelectedpaymentmethod

> int32 PutV1GuestcartsCartIdSelectedpaymentmethod(ctx, cartId).PutV1CartsMineSelectedpaymentmethodRequest(putV1CartsMineSelectedpaymentmethodRequest).Execute()

guest-carts/{cartId}/selected-payment-method



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
	cartId := "cartId_example" // string | The cart ID.
	putV1CartsMineSelectedpaymentmethodRequest := *openapiclient.NewPutV1CartsMineSelectedpaymentmethodRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PutV1CartsMineSelectedpaymentmethodRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdSelectedPaymentMethodAPI.PutV1GuestcartsCartIdSelectedpaymentmethod(context.Background(), cartId).PutV1CartsMineSelectedpaymentmethodRequest(putV1CartsMineSelectedpaymentmethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdSelectedPaymentMethodAPI.PutV1GuestcartsCartIdSelectedpaymentmethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GuestcartsCartIdSelectedpaymentmethod`: int32
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdSelectedPaymentMethodAPI.PutV1GuestcartsCartIdSelectedpaymentmethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GuestcartsCartIdSelectedpaymentmethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsMineSelectedpaymentmethodRequest** | [**PutV1CartsMineSelectedpaymentmethodRequest**](PutV1CartsMineSelectedpaymentmethodRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

