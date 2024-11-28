# \GuestCartsCartIdTotalsInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1GuestcartsCartIdTotalsinformation**](GuestCartsCartIdTotalsInformationAPI.md#PostV1GuestcartsCartIdTotalsinformation) | **Post** /V1/guest-carts/{cartId}/totals-information | guest-carts/{cartId}/totals-information



## PostV1GuestcartsCartIdTotalsinformation

> QuoteDataTotalsInterface PostV1GuestcartsCartIdTotalsinformation(ctx, cartId).PostV1CartsMineTotalsinformationRequest(postV1CartsMineTotalsinformationRequest).Execute()

guest-carts/{cartId}/totals-information



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
	postV1CartsMineTotalsinformationRequest := *openapiclient.NewPostV1CartsMineTotalsinformationRequest(*openapiclient.NewCheckoutDataTotalsInformationInterface(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example"))) // PostV1CartsMineTotalsinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdTotalsInformationAPI.PostV1GuestcartsCartIdTotalsinformation(context.Background(), cartId).PostV1CartsMineTotalsinformationRequest(postV1CartsMineTotalsinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdTotalsInformationAPI.PostV1GuestcartsCartIdTotalsinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdTotalsinformation`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdTotalsInformationAPI.PostV1GuestcartsCartIdTotalsinformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdTotalsinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMineTotalsinformationRequest** | [**PostV1CartsMineTotalsinformationRequest**](PostV1CartsMineTotalsinformationRequest.md) |  | 

### Return type

[**QuoteDataTotalsInterface**](QuoteDataTotalsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

