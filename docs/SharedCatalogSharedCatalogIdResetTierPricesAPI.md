# \SharedCatalogSharedCatalogIdResetTierPricesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SharedCatalogSharedCatalogIdResetTierPrices**](SharedCatalogSharedCatalogIdResetTierPricesAPI.md#PostV1SharedCatalogSharedCatalogIdResetTierPrices) | **Post** /V1/sharedCatalog/{sharedCatalogId}/resetTierPrices | sharedCatalog/{sharedCatalogId}/resetTierPrices



## PostV1SharedCatalogSharedCatalogIdResetTierPrices

> ErrorResponse PostV1SharedCatalogSharedCatalogIdResetTierPrices(ctx, sharedCatalogId).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()

sharedCatalog/{sharedCatalogId}/resetTierPrices



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
	postV1ProductsBasepricesinformationRequest := *openapiclient.NewPostV1ProductsBasepricesinformationRequest([]string{"Skus_example"}) // PostV1ProductsBasepricesinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogSharedCatalogIdResetTierPricesAPI.PostV1SharedCatalogSharedCatalogIdResetTierPrices(context.Background(), sharedCatalogId).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogSharedCatalogIdResetTierPricesAPI.PostV1SharedCatalogSharedCatalogIdResetTierPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SharedCatalogSharedCatalogIdResetTierPrices`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogSharedCatalogIdResetTierPricesAPI.PostV1SharedCatalogSharedCatalogIdResetTierPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedCatalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SharedCatalogSharedCatalogIdResetTierPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsBasepricesinformationRequest** | [**PostV1ProductsBasepricesinformationRequest**](PostV1ProductsBasepricesinformationRequest.md) |  | 

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

