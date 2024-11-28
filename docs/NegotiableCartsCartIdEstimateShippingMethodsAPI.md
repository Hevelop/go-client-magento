# \NegotiableCartsCartIdEstimateShippingMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiablecartsCartIdEstimateshippingmethods**](NegotiableCartsCartIdEstimateShippingMethodsAPI.md#PostV1NegotiablecartsCartIdEstimateshippingmethods) | **Post** /V1/negotiable-carts/{cartId}/estimate-shipping-methods | negotiable-carts/{cartId}/estimate-shipping-methods



## PostV1NegotiablecartsCartIdEstimateshippingmethods

> []QuoteDataShippingMethodInterface PostV1NegotiablecartsCartIdEstimateshippingmethods(ctx, cartId).PostV1CartsMineEstimateshippingmethodsRequest(postV1CartsMineEstimateshippingmethodsRequest).Execute()

negotiable-carts/{cartId}/estimate-shipping-methods



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
	postV1CartsMineEstimateshippingmethodsRequest := *openapiclient.NewPostV1CartsMineEstimateshippingmethodsRequest(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example")) // PostV1CartsMineEstimateshippingmethodsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartsCartIdEstimateShippingMethodsAPI.PostV1NegotiablecartsCartIdEstimateshippingmethods(context.Background(), cartId).PostV1CartsMineEstimateshippingmethodsRequest(postV1CartsMineEstimateshippingmethodsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdEstimateShippingMethodsAPI.PostV1NegotiablecartsCartIdEstimateshippingmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiablecartsCartIdEstimateshippingmethods`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdEstimateShippingMethodsAPI.PostV1NegotiablecartsCartIdEstimateshippingmethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiablecartsCartIdEstimateshippingmethodsRequest struct via the builder pattern


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

