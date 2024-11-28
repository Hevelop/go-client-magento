# \PurchaseOrderCartsCartIdPaymentInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PurchaseordercartsCartIdPaymentinformation**](PurchaseOrderCartsCartIdPaymentInformationAPI.md#GetV1PurchaseordercartsCartIdPaymentinformation) | **Get** /V1/purchase-order-carts/{cartId}/payment-information | purchase-order-carts/{cartId}/payment-information
[**PostV1PurchaseordercartsCartIdPaymentinformation**](PurchaseOrderCartsCartIdPaymentInformationAPI.md#PostV1PurchaseordercartsCartIdPaymentinformation) | **Post** /V1/purchase-order-carts/{cartId}/payment-information | purchase-order-carts/{cartId}/payment-information



## GetV1PurchaseordercartsCartIdPaymentinformation

> CheckoutDataPaymentDetailsInterface GetV1PurchaseordercartsCartIdPaymentinformation(ctx, cartId).Execute()

purchase-order-carts/{cartId}/payment-information



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
	cartId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderCartsCartIdPaymentInformationAPI.GetV1PurchaseordercartsCartIdPaymentinformation(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdPaymentInformationAPI.GetV1PurchaseordercartsCartIdPaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PurchaseordercartsCartIdPaymentinformation`: CheckoutDataPaymentDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdPaymentInformationAPI.GetV1PurchaseordercartsCartIdPaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PurchaseordercartsCartIdPaymentinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutDataPaymentDetailsInterface**](CheckoutDataPaymentDetailsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1PurchaseordercartsCartIdPaymentinformation

> int32 PostV1PurchaseordercartsCartIdPaymentinformation(ctx, cartId).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()

purchase-order-carts/{cartId}/payment-information



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
	cartId := int32(56) // int32 | 
	postV1CartsMinePaymentinformationRequest := *openapiclient.NewPostV1CartsMinePaymentinformationRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PostV1CartsMinePaymentinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderCartsCartIdPaymentInformationAPI.PostV1PurchaseordercartsCartIdPaymentinformation(context.Background(), cartId).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdPaymentInformationAPI.PostV1PurchaseordercartsCartIdPaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1PurchaseordercartsCartIdPaymentinformation`: int32
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdPaymentInformationAPI.PostV1PurchaseordercartsCartIdPaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1PurchaseordercartsCartIdPaymentinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMinePaymentinformationRequest** | [**PostV1CartsMinePaymentinformationRequest**](PostV1CartsMinePaymentinformationRequest.md) |  | 

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

