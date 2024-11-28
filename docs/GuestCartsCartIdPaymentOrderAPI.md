# \GuestCartsCartIdPaymentOrderAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1GuestcartsCartIdPaymentorder**](GuestCartsCartIdPaymentOrderAPI.md#PostV1GuestcartsCartIdPaymentorder) | **Post** /V1/guest-carts/{cartId}/payment-order | guest-carts/{cartId}/payment-order



## PostV1GuestcartsCartIdPaymentorder

> PaymentServicesPaypalDataPaymentOrderInterface PostV1GuestcartsCartIdPaymentorder(ctx, cartId).PostV1CartsMinePaymentorderRequest(postV1CartsMinePaymentorderRequest).Execute()

guest-carts/{cartId}/payment-order



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
	postV1CartsMinePaymentorderRequest := *openapiclient.NewPostV1CartsMinePaymentorderRequest("MethodCode_example", "PaymentSource_example", "Location_example") // PostV1CartsMinePaymentorderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdPaymentOrderAPI.PostV1GuestcartsCartIdPaymentorder(context.Background(), cartId).PostV1CartsMinePaymentorderRequest(postV1CartsMinePaymentorderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdPaymentOrderAPI.PostV1GuestcartsCartIdPaymentorder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdPaymentorder`: PaymentServicesPaypalDataPaymentOrderInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdPaymentOrderAPI.PostV1GuestcartsCartIdPaymentorder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdPaymentorderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMinePaymentorderRequest** | [**PostV1CartsMinePaymentorderRequest**](PostV1CartsMinePaymentorderRequest.md) |  | 

### Return type

[**PaymentServicesPaypalDataPaymentOrderInterface**](PaymentServicesPaypalDataPaymentOrderInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

