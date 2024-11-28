# \CartsMineSelectedPaymentMethodAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMineSelectedpaymentmethod**](CartsMineSelectedPaymentMethodAPI.md#GetV1CartsMineSelectedpaymentmethod) | **Get** /V1/carts/mine/selected-payment-method | carts/mine/selected-payment-method
[**PutV1CartsMineSelectedpaymentmethod**](CartsMineSelectedPaymentMethodAPI.md#PutV1CartsMineSelectedpaymentmethod) | **Put** /V1/carts/mine/selected-payment-method | carts/mine/selected-payment-method



## GetV1CartsMineSelectedpaymentmethod

> QuoteDataPaymentInterface GetV1CartsMineSelectedpaymentmethod(ctx).Execute()

carts/mine/selected-payment-method



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
	resp, r, err := apiClient.CartsMineSelectedPaymentMethodAPI.GetV1CartsMineSelectedpaymentmethod(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineSelectedPaymentMethodAPI.GetV1CartsMineSelectedpaymentmethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineSelectedpaymentmethod`: QuoteDataPaymentInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineSelectedPaymentMethodAPI.GetV1CartsMineSelectedpaymentmethod`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineSelectedpaymentmethodRequest struct via the builder pattern


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


## PutV1CartsMineSelectedpaymentmethod

> string PutV1CartsMineSelectedpaymentmethod(ctx).PutV1CartsMineSelectedpaymentmethodRequest(putV1CartsMineSelectedpaymentmethodRequest).Execute()

carts/mine/selected-payment-method



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
	putV1CartsMineSelectedpaymentmethodRequest := *openapiclient.NewPutV1CartsMineSelectedpaymentmethodRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PutV1CartsMineSelectedpaymentmethodRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineSelectedPaymentMethodAPI.PutV1CartsMineSelectedpaymentmethod(context.Background()).PutV1CartsMineSelectedpaymentmethodRequest(putV1CartsMineSelectedpaymentmethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineSelectedPaymentMethodAPI.PutV1CartsMineSelectedpaymentmethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsMineSelectedpaymentmethod`: string
	fmt.Fprintf(os.Stdout, "Response from `CartsMineSelectedPaymentMethodAPI.PutV1CartsMineSelectedpaymentmethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsMineSelectedpaymentmethodRequest struct via the builder pattern


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

