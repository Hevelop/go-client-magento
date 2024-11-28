# \CartsMineEstimateShippingMethodsByAddressIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineEstimateshippingmethodsbyaddressid**](CartsMineEstimateShippingMethodsByAddressIdAPI.md#PostV1CartsMineEstimateshippingmethodsbyaddressid) | **Post** /V1/carts/mine/estimate-shipping-methods-by-address-id | carts/mine/estimate-shipping-methods-by-address-id



## PostV1CartsMineEstimateshippingmethodsbyaddressid

> []QuoteDataShippingMethodInterface PostV1CartsMineEstimateshippingmethodsbyaddressid(ctx).PostV1CartsMineEstimateshippingmethodsbyaddressidRequest(postV1CartsMineEstimateshippingmethodsbyaddressidRequest).Execute()

carts/mine/estimate-shipping-methods-by-address-id



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
	postV1CartsMineEstimateshippingmethodsbyaddressidRequest := *openapiclient.NewPostV1CartsMineEstimateshippingmethodsbyaddressidRequest(int32(123)) // PostV1CartsMineEstimateshippingmethodsbyaddressidRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineEstimateShippingMethodsByAddressIdAPI.PostV1CartsMineEstimateshippingmethodsbyaddressid(context.Background()).PostV1CartsMineEstimateshippingmethodsbyaddressidRequest(postV1CartsMineEstimateshippingmethodsbyaddressidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineEstimateShippingMethodsByAddressIdAPI.PostV1CartsMineEstimateshippingmethodsbyaddressid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineEstimateshippingmethodsbyaddressid`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineEstimateShippingMethodsByAddressIdAPI.PostV1CartsMineEstimateshippingmethodsbyaddressid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineEstimateshippingmethodsbyaddressidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CartsMineEstimateshippingmethodsbyaddressidRequest** | [**PostV1CartsMineEstimateshippingmethodsbyaddressidRequest**](PostV1CartsMineEstimateshippingmethodsbyaddressidRequest.md) |  | 

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

