# \GuestCartsCartIdPaymentOrderIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdPaymentorderId**](GuestCartsCartIdPaymentOrderIdAPI.md#GetV1GuestcartsCartIdPaymentorderId) | **Get** /V1/guest-carts/{cartId}/payment-order/{id} | guest-carts/{cartId}/payment-order/{id}
[**PostV1GuestcartsCartIdPaymentorderId**](GuestCartsCartIdPaymentOrderIdAPI.md#PostV1GuestcartsCartIdPaymentorderId) | **Post** /V1/guest-carts/{cartId}/payment-order/{id} | guest-carts/{cartId}/payment-order/{id}



## GetV1GuestcartsCartIdPaymentorderId

> PaymentServicesPaypalDataPaymentOrderDetailsInterface GetV1GuestcartsCartIdPaymentorderId(ctx, cartId, id).Execute()

guest-carts/{cartId}/payment-order/{id}



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
	cartId := "cartId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdPaymentOrderIdAPI.GetV1GuestcartsCartIdPaymentorderId(context.Background(), cartId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdPaymentOrderIdAPI.GetV1GuestcartsCartIdPaymentorderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdPaymentorderId`: PaymentServicesPaypalDataPaymentOrderDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdPaymentOrderIdAPI.GetV1GuestcartsCartIdPaymentorderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdPaymentorderIdRequest struct via the builder pattern


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


## PostV1GuestcartsCartIdPaymentorderId

> bool PostV1GuestcartsCartIdPaymentorderId(ctx, cartId, id).Execute()

guest-carts/{cartId}/payment-order/{id}



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
	cartId := "cartId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdPaymentOrderIdAPI.PostV1GuestcartsCartIdPaymentorderId(context.Background(), cartId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdPaymentOrderIdAPI.PostV1GuestcartsCartIdPaymentorderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdPaymentorderId`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdPaymentOrderIdAPI.PostV1GuestcartsCartIdPaymentorderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdPaymentorderIdRequest struct via the builder pattern


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

