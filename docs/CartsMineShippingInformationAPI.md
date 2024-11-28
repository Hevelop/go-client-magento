# \CartsMineShippingInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineShippinginformation**](CartsMineShippingInformationAPI.md#PostV1CartsMineShippinginformation) | **Post** /V1/carts/mine/shipping-information | carts/mine/shipping-information



## PostV1CartsMineShippinginformation

> CheckoutDataPaymentDetailsInterface PostV1CartsMineShippinginformation(ctx).PostV1CartsMineShippinginformationRequest(postV1CartsMineShippinginformationRequest).Execute()

carts/mine/shipping-information



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
	postV1CartsMineShippinginformationRequest := *openapiclient.NewPostV1CartsMineShippinginformationRequest(*openapiclient.NewCheckoutDataShippingInformationInterface(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example"), "ShippingMethodCode_example", "ShippingCarrierCode_example")) // PostV1CartsMineShippinginformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineShippingInformationAPI.PostV1CartsMineShippinginformation(context.Background()).PostV1CartsMineShippinginformationRequest(postV1CartsMineShippinginformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineShippingInformationAPI.PostV1CartsMineShippinginformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineShippinginformation`: CheckoutDataPaymentDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineShippingInformationAPI.PostV1CartsMineShippinginformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineShippinginformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CartsMineShippinginformationRequest** | [**PostV1CartsMineShippinginformationRequest**](PostV1CartsMineShippinginformationRequest.md) |  | 

### Return type

[**CheckoutDataPaymentDetailsInterface**](CheckoutDataPaymentDetailsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

