# \GuestCartsCartIdPaymentInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdPaymentinformation**](GuestCartsCartIdPaymentInformationAPI.md#GetV1GuestcartsCartIdPaymentinformation) | **Get** /V1/guest-carts/{cartId}/payment-information | guest-carts/{cartId}/payment-information
[**PostV1GuestcartsCartIdPaymentinformation**](GuestCartsCartIdPaymentInformationAPI.md#PostV1GuestcartsCartIdPaymentinformation) | **Post** /V1/guest-carts/{cartId}/payment-information | guest-carts/{cartId}/payment-information



## GetV1GuestcartsCartIdPaymentinformation

> CheckoutDataPaymentDetailsInterface GetV1GuestcartsCartIdPaymentinformation(ctx, cartId).Execute()

guest-carts/{cartId}/payment-information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdPaymentInformationAPI.GetV1GuestcartsCartIdPaymentinformation(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdPaymentInformationAPI.GetV1GuestcartsCartIdPaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdPaymentinformation`: CheckoutDataPaymentDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdPaymentInformationAPI.GetV1GuestcartsCartIdPaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdPaymentinformationRequest struct via the builder pattern


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


## PostV1GuestcartsCartIdPaymentinformation

> int32 PostV1GuestcartsCartIdPaymentinformation(ctx, cartId).PostV1GuestcartsCartIdPaymentinformationRequest(postV1GuestcartsCartIdPaymentinformationRequest).Execute()

guest-carts/{cartId}/payment-information



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
	resp, r, err := apiClient.GuestCartsCartIdPaymentInformationAPI.PostV1GuestcartsCartIdPaymentinformation(context.Background(), cartId).PostV1GuestcartsCartIdPaymentinformationRequest(postV1GuestcartsCartIdPaymentinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdPaymentInformationAPI.PostV1GuestcartsCartIdPaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdPaymentinformation`: int32
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdPaymentInformationAPI.PostV1GuestcartsCartIdPaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdPaymentinformationRequest struct via the builder pattern


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

