# \ProductsSkuGroupPricesCustomerGroupIdTiersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsSkuGrouppricesCustomerGroupIdTiers**](ProductsSkuGroupPricesCustomerGroupIdTiersAPI.md#GetV1ProductsSkuGrouppricesCustomerGroupIdTiers) | **Get** /V1/products/{sku}/group-prices/{customerGroupId}/tiers | products/{sku}/group-prices/{customerGroupId}/tiers



## GetV1ProductsSkuGrouppricesCustomerGroupIdTiers

> []CatalogDataProductTierPriceInterface GetV1ProductsSkuGrouppricesCustomerGroupIdTiers(ctx, sku, customerGroupId).Execute()

products/{sku}/group-prices/{customerGroupId}/tiers



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
	sku := "sku_example" // string | 
	customerGroupId := "customerGroupId_example" // string | 'all' can be used to specify 'ALL GROUPS'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuGroupPricesCustomerGroupIdTiersAPI.GetV1ProductsSkuGrouppricesCustomerGroupIdTiers(context.Background(), sku, customerGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuGroupPricesCustomerGroupIdTiersAPI.GetV1ProductsSkuGrouppricesCustomerGroupIdTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuGrouppricesCustomerGroupIdTiers`: []CatalogDataProductTierPriceInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuGroupPricesCustomerGroupIdTiersAPI.GetV1ProductsSkuGrouppricesCustomerGroupIdTiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**customerGroupId** | **string** | &#39;all&#39; can be used to specify &#39;ALL GROUPS&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuGrouppricesCustomerGroupIdTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CatalogDataProductTierPriceInterface**](CatalogDataProductTierPriceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

