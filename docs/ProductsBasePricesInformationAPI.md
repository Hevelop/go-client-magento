# \ProductsBasePricesInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsBasepricesinformation**](ProductsBasePricesInformationAPI.md#PostV1ProductsBasepricesinformation) | **Post** /V1/products/base-prices-information | products/base-prices-information



## PostV1ProductsBasepricesinformation

> []CatalogDataBasePriceInterface PostV1ProductsBasepricesinformation(ctx).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()

products/base-prices-information



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
	resp, r, err := apiClient.ProductsBasePricesInformationAPI.PostV1ProductsBasepricesinformation(context.Background()).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsBasePricesInformationAPI.PostV1ProductsBasepricesinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsBasepricesinformation`: []CatalogDataBasePriceInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsBasePricesInformationAPI.PostV1ProductsBasepricesinformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsBasepricesinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsBasepricesinformationRequest** | [**PostV1ProductsBasepricesinformationRequest**](PostV1ProductsBasepricesinformationRequest.md) |  | 

### Return type

[**[]CatalogDataBasePriceInterface**](CatalogDataBasePriceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

