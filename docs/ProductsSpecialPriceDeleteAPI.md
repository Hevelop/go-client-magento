# \ProductsSpecialPriceDeleteAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsSpecialpricedelete**](ProductsSpecialPriceDeleteAPI.md#PostV1ProductsSpecialpricedelete) | **Post** /V1/products/special-price-delete | products/special-price-delete



## PostV1ProductsSpecialpricedelete

> []CatalogDataPriceUpdateResultInterface PostV1ProductsSpecialpricedelete(ctx).PostV1ProductsSpecialpriceRequest(postV1ProductsSpecialpriceRequest).Execute()

products/special-price-delete



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
	postV1ProductsSpecialpriceRequest := *openapiclient.NewPostV1ProductsSpecialpriceRequest([]openapiclient.CatalogDataSpecialPriceInterface{*openapiclient.NewCatalogDataSpecialPriceInterface(float32(123), int32(123), "Sku_example", "PriceFrom_example", "PriceTo_example")}) // PostV1ProductsSpecialpriceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSpecialPriceDeleteAPI.PostV1ProductsSpecialpricedelete(context.Background()).PostV1ProductsSpecialpriceRequest(postV1ProductsSpecialpriceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSpecialPriceDeleteAPI.PostV1ProductsSpecialpricedelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSpecialpricedelete`: []CatalogDataPriceUpdateResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSpecialPriceDeleteAPI.PostV1ProductsSpecialpricedelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSpecialpricedeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsSpecialpriceRequest** | [**PostV1ProductsSpecialpriceRequest**](PostV1ProductsSpecialpriceRequest.md) |  | 

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

