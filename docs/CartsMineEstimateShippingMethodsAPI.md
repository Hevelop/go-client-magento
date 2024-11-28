# \CartsMineEstimateShippingMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineEstimateshippingmethods**](CartsMineEstimateShippingMethodsAPI.md#PostV1CartsMineEstimateshippingmethods) | **Post** /V1/carts/mine/estimate-shipping-methods | carts/mine/estimate-shipping-methods



## PostV1CartsMineEstimateshippingmethods

> []QuoteDataShippingMethodInterface PostV1CartsMineEstimateshippingmethods(ctx).PostV1CartsMineEstimateshippingmethodsRequest(postV1CartsMineEstimateshippingmethodsRequest).Execute()

carts/mine/estimate-shipping-methods



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
	postV1CartsMineEstimateshippingmethodsRequest := *openapiclient.NewPostV1CartsMineEstimateshippingmethodsRequest(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example")) // PostV1CartsMineEstimateshippingmethodsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineEstimateShippingMethodsAPI.PostV1CartsMineEstimateshippingmethods(context.Background()).PostV1CartsMineEstimateshippingmethodsRequest(postV1CartsMineEstimateshippingmethodsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineEstimateShippingMethodsAPI.PostV1CartsMineEstimateshippingmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineEstimateshippingmethods`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineEstimateShippingMethodsAPI.PostV1CartsMineEstimateshippingmethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineEstimateshippingmethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CartsMineEstimateshippingmethodsRequest** | [**PostV1CartsMineEstimateshippingmethodsRequest**](PostV1CartsMineEstimateshippingmethodsRequest.md) |  | 

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

