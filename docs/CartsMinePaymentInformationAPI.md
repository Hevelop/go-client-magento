# \CartsMinePaymentInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMinePaymentinformation**](CartsMinePaymentInformationAPI.md#GetV1CartsMinePaymentinformation) | **Get** /V1/carts/mine/payment-information | carts/mine/payment-information
[**PostV1CartsMinePaymentinformation**](CartsMinePaymentInformationAPI.md#PostV1CartsMinePaymentinformation) | **Post** /V1/carts/mine/payment-information | carts/mine/payment-information



## GetV1CartsMinePaymentinformation

> CheckoutDataPaymentDetailsInterface GetV1CartsMinePaymentinformation(ctx).Execute()

carts/mine/payment-information



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
	resp, r, err := apiClient.CartsMinePaymentInformationAPI.GetV1CartsMinePaymentinformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMinePaymentInformationAPI.GetV1CartsMinePaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMinePaymentinformation`: CheckoutDataPaymentDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMinePaymentInformationAPI.GetV1CartsMinePaymentinformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMinePaymentinformationRequest struct via the builder pattern


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


## PostV1CartsMinePaymentinformation

> int32 PostV1CartsMinePaymentinformation(ctx).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()

carts/mine/payment-information



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
	postV1CartsMinePaymentinformationRequest := *openapiclient.NewPostV1CartsMinePaymentinformationRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PostV1CartsMinePaymentinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMinePaymentInformationAPI.PostV1CartsMinePaymentinformation(context.Background()).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMinePaymentInformationAPI.PostV1CartsMinePaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMinePaymentinformation`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsMinePaymentInformationAPI.PostV1CartsMinePaymentinformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMinePaymentinformationRequest struct via the builder pattern


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

