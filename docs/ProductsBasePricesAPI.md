# \ProductsBasePricesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsBaseprices**](ProductsBasePricesAPI.md#PostV1ProductsBaseprices) | **Post** /V1/products/base-prices | products/base-prices



## PostV1ProductsBaseprices

> []CatalogDataPriceUpdateResultInterface PostV1ProductsBaseprices(ctx).PostV1ProductsBasepricesRequest(postV1ProductsBasepricesRequest).Execute()

products/base-prices



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
	postV1ProductsBasepricesRequest := *openapiclient.NewPostV1ProductsBasepricesRequest([]openapiclient.CatalogDataBasePriceInterface{*openapiclient.NewCatalogDataBasePriceInterface(float32(123), int32(123), "Sku_example")}) // PostV1ProductsBasepricesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsBasePricesAPI.PostV1ProductsBaseprices(context.Background()).PostV1ProductsBasepricesRequest(postV1ProductsBasepricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsBasePricesAPI.PostV1ProductsBaseprices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsBaseprices`: []CatalogDataPriceUpdateResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsBasePricesAPI.PostV1ProductsBaseprices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsBasepricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsBasepricesRequest** | [**PostV1ProductsBasepricesRequest**](PostV1ProductsBasepricesRequest.md) |  | 

### Return type

[**[]CatalogDataPriceUpdateResultInterface**](CatalogDataPriceUpdateResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

