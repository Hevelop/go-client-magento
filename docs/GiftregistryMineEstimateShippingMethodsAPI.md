# \GiftregistryMineEstimateShippingMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1GiftregistryMineEstimateshippingmethods**](GiftregistryMineEstimateShippingMethodsAPI.md#PostV1GiftregistryMineEstimateshippingmethods) | **Post** /V1/giftregistry/mine/estimate-shipping-methods | giftregistry/mine/estimate-shipping-methods



## PostV1GiftregistryMineEstimateshippingmethods

> []QuoteDataShippingMethodInterface PostV1GiftregistryMineEstimateshippingmethods(ctx).PostV1GiftregistryMineEstimateshippingmethodsRequest(postV1GiftregistryMineEstimateshippingmethodsRequest).Execute()

giftregistry/mine/estimate-shipping-methods



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
	postV1GiftregistryMineEstimateshippingmethodsRequest := *openapiclient.NewPostV1GiftregistryMineEstimateshippingmethodsRequest(int32(123)) // PostV1GiftregistryMineEstimateshippingmethodsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftregistryMineEstimateShippingMethodsAPI.PostV1GiftregistryMineEstimateshippingmethods(context.Background()).PostV1GiftregistryMineEstimateshippingmethodsRequest(postV1GiftregistryMineEstimateshippingmethodsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftregistryMineEstimateShippingMethodsAPI.PostV1GiftregistryMineEstimateshippingmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GiftregistryMineEstimateshippingmethods`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `GiftregistryMineEstimateShippingMethodsAPI.PostV1GiftregistryMineEstimateshippingmethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GiftregistryMineEstimateshippingmethodsRequest struct via the builder pattern


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

