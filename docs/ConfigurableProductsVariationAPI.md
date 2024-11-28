# \ConfigurableProductsVariationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1ConfigurableproductsVariation**](ConfigurableProductsVariationAPI.md#PutV1ConfigurableproductsVariation) | **Put** /V1/configurable-products/variation | configurable-products/variation



## PutV1ConfigurableproductsVariation

> []CatalogDataProductInterface PutV1ConfigurableproductsVariation(ctx).PutV1ConfigurableproductsVariationRequest(putV1ConfigurableproductsVariationRequest).Execute()

configurable-products/variation



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
	putV1ConfigurableproductsVariationRequest := *openapiclient.NewPutV1ConfigurableproductsVariationRequest(*openapiclient.NewCatalogDataProductInterface("Sku_example"), []openapiclient.ConfigurableProductDataOptionInterface{*openapiclient.NewConfigurableProductDataOptionInterface()}) // PutV1ConfigurableproductsVariationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsVariationAPI.PutV1ConfigurableproductsVariation(context.Background()).PutV1ConfigurableproductsVariationRequest(putV1ConfigurableproductsVariationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsVariationAPI.PutV1ConfigurableproductsVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ConfigurableproductsVariation`: []CatalogDataProductInterface
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsVariationAPI.PutV1ConfigurableproductsVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ConfigurableproductsVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1ConfigurableproductsVariationRequest** | [**PutV1ConfigurableproductsVariationRequest**](PutV1ConfigurableproductsVariationRequest.md) |  | 

### Return type

[**[]CatalogDataProductInterface**](CatalogDataProductInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

