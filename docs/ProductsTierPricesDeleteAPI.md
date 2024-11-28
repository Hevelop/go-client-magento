# \ProductsTierPricesDeleteAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsTierpricesdelete**](ProductsTierPricesDeleteAPI.md#PostV1ProductsTierpricesdelete) | **Post** /V1/products/tier-prices-delete | products/tier-prices-delete



## PostV1ProductsTierpricesdelete

> []CatalogDataPriceUpdateResultInterface PostV1ProductsTierpricesdelete(ctx).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()

products/tier-prices-delete



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
	putV1ProductsTierpricesRequest := *openapiclient.NewPutV1ProductsTierpricesRequest([]openapiclient.CatalogDataTierPriceInterface{*openapiclient.NewCatalogDataTierPriceInterface(float32(123), "PriceType_example", int32(123), "Sku_example", "CustomerGroup_example", float32(123))}) // PutV1ProductsTierpricesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsTierPricesDeleteAPI.PostV1ProductsTierpricesdelete(context.Background()).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsTierPricesDeleteAPI.PostV1ProductsTierpricesdelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsTierpricesdelete`: []CatalogDataPriceUpdateResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsTierPricesDeleteAPI.PostV1ProductsTierpricesdelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsTierpricesdeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1ProductsTierpricesRequest** | [**PutV1ProductsTierpricesRequest**](PutV1ProductsTierpricesRequest.md) |  | 

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

