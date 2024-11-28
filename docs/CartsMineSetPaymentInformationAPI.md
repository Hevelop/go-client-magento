# \CartsMineSetPaymentInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineSetpaymentinformation**](CartsMineSetPaymentInformationAPI.md#PostV1CartsMineSetpaymentinformation) | **Post** /V1/carts/mine/set-payment-information | carts/mine/set-payment-information



## PostV1CartsMineSetpaymentinformation

> int32 PostV1CartsMineSetpaymentinformation(ctx).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()

carts/mine/set-payment-information



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
	resp, r, err := apiClient.CartsMineSetPaymentInformationAPI.PostV1CartsMineSetpaymentinformation(context.Background()).PostV1CartsMinePaymentinformationRequest(postV1CartsMinePaymentinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineSetPaymentInformationAPI.PostV1CartsMineSetpaymentinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineSetpaymentinformation`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsMineSetPaymentInformationAPI.PostV1CartsMineSetpaymentinformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineSetpaymentinformationRequest struct via the builder pattern


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

