# \GuestGiftregistryCartIdEstimateShippingMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1GuestgiftregistryCartIdEstimateshippingmethods**](GuestGiftregistryCartIdEstimateShippingMethodsAPI.md#PostV1GuestgiftregistryCartIdEstimateshippingmethods) | **Post** /V1/guest-giftregistry/{cartId}/estimate-shipping-methods | guest-giftregistry/{cartId}/estimate-shipping-methods



## PostV1GuestgiftregistryCartIdEstimateshippingmethods

> []QuoteDataShippingMethodInterface PostV1GuestgiftregistryCartIdEstimateshippingmethods(ctx, cartId).PostV1GiftregistryMineEstimateshippingmethodsRequest(postV1GiftregistryMineEstimateshippingmethodsRequest).Execute()

guest-giftregistry/{cartId}/estimate-shipping-methods



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
	cartId := "cartId_example" // string | The shopping cart ID.
	postV1GiftregistryMineEstimateshippingmethodsRequest := *openapiclient.NewPostV1GiftregistryMineEstimateshippingmethodsRequest(int32(123)) // PostV1GiftregistryMineEstimateshippingmethodsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestGiftregistryCartIdEstimateShippingMethodsAPI.PostV1GuestgiftregistryCartIdEstimateshippingmethods(context.Background(), cartId).PostV1GiftregistryMineEstimateshippingmethodsRequest(postV1GiftregistryMineEstimateshippingmethodsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestGiftregistryCartIdEstimateShippingMethodsAPI.PostV1GuestgiftregistryCartIdEstimateshippingmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestgiftregistryCartIdEstimateshippingmethods`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestGiftregistryCartIdEstimateShippingMethodsAPI.PostV1GuestgiftregistryCartIdEstimateshippingmethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The shopping cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestgiftregistryCartIdEstimateshippingmethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1GiftregistryMineEstimateshippingmethodsRequest** | [**PostV1GiftregistryMineEstimateshippingmethodsRequest**](PostV1GiftregistryMineEstimateshippingmethodsRequest.md) |  | 

### Return type

[**[]QuoteDataShippingMethodInterface**](QuoteDataShippingMethodInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

