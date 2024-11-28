# \CartsMineTotalsInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineTotalsinformation**](CartsMineTotalsInformationAPI.md#PostV1CartsMineTotalsinformation) | **Post** /V1/carts/mine/totals-information | carts/mine/totals-information



## PostV1CartsMineTotalsinformation

> QuoteDataTotalsInterface PostV1CartsMineTotalsinformation(ctx).PostV1CartsMineTotalsinformationRequest(postV1CartsMineTotalsinformationRequest).Execute()

carts/mine/totals-information



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
	postV1CartsMineTotalsinformationRequest := *openapiclient.NewPostV1CartsMineTotalsinformationRequest(*openapiclient.NewCheckoutDataTotalsInformationInterface(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example"))) // PostV1CartsMineTotalsinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineTotalsInformationAPI.PostV1CartsMineTotalsinformation(context.Background()).PostV1CartsMineTotalsinformationRequest(postV1CartsMineTotalsinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineTotalsInformationAPI.PostV1CartsMineTotalsinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineTotalsinformation`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineTotalsInformationAPI.PostV1CartsMineTotalsinformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineTotalsinformationRequest struct via the builder pattern


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

