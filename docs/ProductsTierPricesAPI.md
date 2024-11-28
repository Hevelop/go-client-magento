# \ProductsTierPricesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsTierprices**](ProductsTierPricesAPI.md#PostV1ProductsTierprices) | **Post** /V1/products/tier-prices | products/tier-prices
[**PutV1ProductsTierprices**](ProductsTierPricesAPI.md#PutV1ProductsTierprices) | **Put** /V1/products/tier-prices | products/tier-prices



## PostV1ProductsTierprices

> []CatalogDataPriceUpdateResultInterface PostV1ProductsTierprices(ctx).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()

products/tier-prices



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
	resp, r, err := apiClient.ProductsTierPricesAPI.PostV1ProductsTierprices(context.Background()).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsTierPricesAPI.PostV1ProductsTierprices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsTierprices`: []CatalogDataPriceUpdateResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsTierPricesAPI.PostV1ProductsTierprices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsTierpricesRequest struct via the builder pattern


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


## PutV1ProductsTierprices

> []CatalogDataPriceUpdateResultInterface PutV1ProductsTierprices(ctx).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()

products/tier-prices



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
	resp, r, err := apiClient.ProductsTierPricesAPI.PutV1ProductsTierprices(context.Background()).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsTierPricesAPI.PutV1ProductsTierprices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsTierprices`: []CatalogDataPriceUpdateResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsTierPricesAPI.PutV1ProductsTierprices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsTierpricesRequest struct via the builder pattern


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

