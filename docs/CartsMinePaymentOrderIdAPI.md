# \CartsMinePaymentOrderIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMinePaymentorderId**](CartsMinePaymentOrderIdAPI.md#GetV1CartsMinePaymentorderId) | **Get** /V1/carts/mine/payment-order/{id} | carts/mine/payment-order/{id}
[**PostV1CartsMinePaymentorderId**](CartsMinePaymentOrderIdAPI.md#PostV1CartsMinePaymentorderId) | **Post** /V1/carts/mine/payment-order/{id} | carts/mine/payment-order/{id}



## GetV1CartsMinePaymentorderId

> PaymentServicesPaypalDataPaymentOrderDetailsInterface GetV1CartsMinePaymentorderId(ctx, id).Execute()

carts/mine/payment-order/{id}



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMinePaymentOrderIdAPI.GetV1CartsMinePaymentorderId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMinePaymentOrderIdAPI.GetV1CartsMinePaymentorderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMinePaymentorderId`: PaymentServicesPaypalDataPaymentOrderDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMinePaymentOrderIdAPI.GetV1CartsMinePaymentorderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMinePaymentorderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentServicesPaypalDataPaymentOrderDetailsInterface**](PaymentServicesPaypalDataPaymentOrderDetailsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1CartsMinePaymentorderId

> bool PostV1CartsMinePaymentorderId(ctx, id).Execute()

carts/mine/payment-order/{id}



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMinePaymentOrderIdAPI.PostV1CartsMinePaymentorderId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMinePaymentOrderIdAPI.PostV1CartsMinePaymentorderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMinePaymentorderId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMinePaymentOrderIdAPI.PostV1CartsMinePaymentorderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMinePaymentorderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

