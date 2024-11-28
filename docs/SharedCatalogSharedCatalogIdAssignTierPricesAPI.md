# \SharedCatalogSharedCatalogIdAssignTierPricesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SharedCatalogSharedCatalogIdAssignTierPrices**](SharedCatalogSharedCatalogIdAssignTierPricesAPI.md#PostV1SharedCatalogSharedCatalogIdAssignTierPrices) | **Post** /V1/sharedCatalog/{sharedCatalogId}/assignTierPrices | sharedCatalog/{sharedCatalogId}/assignTierPrices



## PostV1SharedCatalogSharedCatalogIdAssignTierPrices

> ErrorResponse PostV1SharedCatalogSharedCatalogIdAssignTierPrices(ctx, sharedCatalogId).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()

sharedCatalog/{sharedCatalogId}/assignTierPrices



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
	sharedCatalogId := int32(56) // int32 | 
	putV1ProductsTierpricesRequest := *openapiclient.NewPutV1ProductsTierpricesRequest([]openapiclient.CatalogDataTierPriceInterface{*openapiclient.NewCatalogDataTierPriceInterface(float32(123), "PriceType_example", int32(123), "Sku_example", "CustomerGroup_example", float32(123))}) // PutV1ProductsTierpricesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogSharedCatalogIdAssignTierPricesAPI.PostV1SharedCatalogSharedCatalogIdAssignTierPrices(context.Background(), sharedCatalogId).PutV1ProductsTierpricesRequest(putV1ProductsTierpricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogSharedCatalogIdAssignTierPricesAPI.PostV1SharedCatalogSharedCatalogIdAssignTierPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SharedCatalogSharedCatalogIdAssignTierPrices`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogSharedCatalogIdAssignTierPricesAPI.PostV1SharedCatalogSharedCatalogIdAssignTierPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedCatalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SharedCatalogSharedCatalogIdAssignTierPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1ProductsTierpricesRequest** | [**PutV1ProductsTierpricesRequest**](PutV1ProductsTierpricesRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

