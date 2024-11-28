# \ProductsCostAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsCost**](ProductsCostAPI.md#PostV1ProductsCost) | **Post** /V1/products/cost | products/cost



## PostV1ProductsCost

> []CatalogDataPriceUpdateResultInterface PostV1ProductsCost(ctx).PostV1ProductsCostRequest(postV1ProductsCostRequest).Execute()

products/cost



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
	postV1ProductsCostRequest := *openapiclient.NewPostV1ProductsCostRequest([]openapiclient.CatalogDataCostInterface{*openapiclient.NewCatalogDataCostInterface(float32(123), int32(123), "Sku_example")}) // PostV1ProductsCostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsCostAPI.PostV1ProductsCost(context.Background()).PostV1ProductsCostRequest(postV1ProductsCostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsCostAPI.PostV1ProductsCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsCost`: []CatalogDataPriceUpdateResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsCostAPI.PostV1ProductsCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsCostRequest** | [**PostV1ProductsCostRequest**](PostV1ProductsCostRequest.md) |  | 

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

