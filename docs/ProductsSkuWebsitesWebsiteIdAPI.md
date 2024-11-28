# \ProductsSkuWebsitesWebsiteIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsSkuWebsitesWebsiteId**](ProductsSkuWebsitesWebsiteIdAPI.md#DeleteV1ProductsSkuWebsitesWebsiteId) | **Delete** /V1/products/{sku}/websites/{websiteId} | products/{sku}/websites/{websiteId}



## DeleteV1ProductsSkuWebsitesWebsiteId

> bool DeleteV1ProductsSkuWebsitesWebsiteId(ctx, sku, websiteId).Execute()

products/{sku}/websites/{websiteId}



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
	websiteId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuWebsitesWebsiteIdAPI.DeleteV1ProductsSkuWebsitesWebsiteId(context.Background(), sku, websiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuWebsitesWebsiteIdAPI.DeleteV1ProductsSkuWebsitesWebsiteId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsSkuWebsitesWebsiteId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuWebsitesWebsiteIdAPI.DeleteV1ProductsSkuWebsitesWebsiteId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**websiteId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsSkuWebsitesWebsiteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

