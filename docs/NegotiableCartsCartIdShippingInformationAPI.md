# \NegotiableCartsCartIdShippingInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiablecartsCartIdShippinginformation**](NegotiableCartsCartIdShippingInformationAPI.md#PostV1NegotiablecartsCartIdShippinginformation) | **Post** /V1/negotiable-carts/{cartId}/shipping-information | negotiable-carts/{cartId}/shipping-information



## PostV1NegotiablecartsCartIdShippinginformation

> CheckoutDataPaymentDetailsInterface PostV1NegotiablecartsCartIdShippinginformation(ctx, cartId).PostV1CartsMineShippinginformationRequest(postV1CartsMineShippinginformationRequest).Execute()

negotiable-carts/{cartId}/shipping-information



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
	postV1CartsMineShippinginformationRequest := *openapiclient.NewPostV1CartsMineShippinginformationRequest(*openapiclient.NewCheckoutDataShippingInformationInterface(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example"), "ShippingMethodCode_example", "ShippingCarrierCode_example")) // PostV1CartsMineShippinginformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartsCartIdShippingInformationAPI.PostV1NegotiablecartsCartIdShippinginformation(context.Background(), cartId).PostV1CartsMineShippinginformationRequest(postV1CartsMineShippinginformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdShippingInformationAPI.PostV1NegotiablecartsCartIdShippinginformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiablecartsCartIdShippinginformation`: CheckoutDataPaymentDetailsInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdShippingInformationAPI.PostV1NegotiablecartsCartIdShippinginformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiablecartsCartIdShippinginformationRequest struct via the builder pattern


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

