# \CartsMinePaymentOrderAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMinePaymentorder**](CartsMinePaymentOrderAPI.md#PostV1CartsMinePaymentorder) | **Post** /V1/carts/mine/payment-order | carts/mine/payment-order



## PostV1CartsMinePaymentorder

> PaymentServicesPaypalDataPaymentOrderInterface PostV1CartsMinePaymentorder(ctx).PostV1CartsMinePaymentorderRequest(postV1CartsMinePaymentorderRequest).Execute()

carts/mine/payment-order



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
	postV1CartsMinePaymentorderRequest := *openapiclient.NewPostV1CartsMinePaymentorderRequest("MethodCode_example", "PaymentSource_example", "Location_example") // PostV1CartsMinePaymentorderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMinePaymentOrderAPI.PostV1CartsMinePaymentorder(context.Background()).PostV1CartsMinePaymentorderRequest(postV1CartsMinePaymentorderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMinePaymentOrderAPI.PostV1CartsMinePaymentorder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMinePaymentorder`: PaymentServicesPaypalDataPaymentOrderInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMinePaymentOrderAPI.PostV1CartsMinePaymentorder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMinePaymentorderRequest struct via the builder pattern


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

