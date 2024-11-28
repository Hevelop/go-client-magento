# \ProductsSpecialPriceInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsSpecialpriceinformation**](ProductsSpecialPriceInformationAPI.md#PostV1ProductsSpecialpriceinformation) | **Post** /V1/products/special-price-information | products/special-price-information



## PostV1ProductsSpecialpriceinformation

> []CatalogDataSpecialPriceInterface PostV1ProductsSpecialpriceinformation(ctx).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()

products/special-price-information



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
	postV1ProductsBasepricesinformationRequest := *openapiclient.NewPostV1ProductsBasepricesinformationRequest([]string{"Skus_example"}) // PostV1ProductsBasepricesinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSpecialPriceInformationAPI.PostV1ProductsSpecialpriceinformation(context.Background()).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSpecialPriceInformationAPI.PostV1ProductsSpecialpriceinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSpecialpriceinformation`: []CatalogDataSpecialPriceInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSpecialPriceInformationAPI.PostV1ProductsSpecialpriceinformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSpecialpriceinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsBasepricesinformationRequest** | [**PostV1ProductsBasepricesinformationRequest**](PostV1ProductsBasepricesinformationRequest.md) |  | 

### Return type

[**[]CatalogDataSpecialPriceInterface**](CatalogDataSpecialPriceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

