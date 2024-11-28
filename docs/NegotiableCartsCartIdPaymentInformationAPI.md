# \NegotiableCartsCartIdPaymentInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1NegotiablecartsCartIdPaymentinformation**](NegotiableCartsCartIdPaymentInformationAPI.md#GetV1NegotiablecartsCartIdPaymentinformation) | **Get** /V1/negotiable-carts/{cartId}/payment-information | negotiable-carts/{cartId}/payment-information
[**PostV1NegotiablecartsCartIdPaymentinformation**](NegotiableCartsCartIdPaymentInformationAPI.md#PostV1NegotiablecartsCartIdPaymentinformation) | **Post** /V1/negotiable-carts/{cartId}/payment-information | negotiable-carts/{cartId}/payment-information



## GetV1NegotiablecartsCartIdPaymentinformation

> CheckoutDataPaymentDetailsInterface GetV1NegotiablecartsCartIdPaymentinformation(ctx, cartId).Execute()

negotiable-carts/{cartId}/payment-information



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
	resp, r, err := apiClient.NegotiableCartsCartIdPaymentInformationAPI.GetV1NegotiablecartsCartIdPaymentinformation(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdPaymentInformationAPI.GetV1NegotiablecartsCartIdPaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1NegotiablecartsCartIdPaymentinformation`: CheckoutDataPaymentDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdPaymentInformationAPI.GetV1NegotiablecartsCartIdPaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1NegotiablecartsCartIdPaymentinformationRequest struct via the builder pattern


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


## PostV1NegotiablecartsCartIdPaymentinformation

> int32 PostV1NegotiablecartsCartIdPaymentinformation(ctx, cartId).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()

negotiable-carts/{cartId}/payment-information



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
	postV1CartsMinePaymentinformationRequest := *openapiclient.NewPostV1CartsMinePaymentinformationRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PostV1CartsMinePaymentinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartsCartIdPaymentInformationAPI.PostV1NegotiablecartsCartIdPaymentinformation(context.Background(), cartId).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdPaymentInformationAPI.PostV1NegotiablecartsCartIdPaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiablecartsCartIdPaymentinformation`: int32
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdPaymentInformationAPI.PostV1NegotiablecartsCartIdPaymentinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiablecartsCartIdPaymentinformationRequest struct via the builder pattern


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

