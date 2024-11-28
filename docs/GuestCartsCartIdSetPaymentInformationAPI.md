# \GuestCartsCartIdSetPaymentInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1GuestcartsCartIdSetpaymentinformation**](GuestCartsCartIdSetPaymentInformationAPI.md#PostV1GuestcartsCartIdSetpaymentinformation) | **Post** /V1/guest-carts/{cartId}/set-payment-information | guest-carts/{cartId}/set-payment-information



## PostV1GuestcartsCartIdSetpaymentinformation

> int32 PostV1GuestcartsCartIdSetpaymentinformation(ctx, cartId).PostV1GuestcartsCartIdPaymentinformationRequest(postV1GuestcartsCartIdPaymentinformationRequest).Execute()

guest-carts/{cartId}/set-payment-information



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
	postV1GuestcartsCartIdPaymentinformationRequest := *openapiclient.NewPostV1GuestcartsCartIdPaymentinformationRequest("Email_example", *openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PostV1GuestcartsCartIdPaymentinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdSetPaymentInformationAPI.PostV1GuestcartsCartIdSetpaymentinformation(context.Background(), cartId).PostV1GuestcartsCartIdPaymentinformationRequest(postV1GuestcartsCartIdPaymentinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdSetPaymentInformationAPI.PostV1GuestcartsCartIdSetpaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdSetpaymentinformation`: int32
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdSetPaymentInformationAPI.PostV1GuestcartsCartIdSetpaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdSetpaymentinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1GuestcartsCartIdPaymentinformationRequest** | [**PostV1GuestcartsCartIdPaymentinformationRequest**](PostV1GuestcartsCartIdPaymentinformationRequest.md) |  | 

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

