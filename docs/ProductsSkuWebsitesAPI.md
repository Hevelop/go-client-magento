# \ProductsSkuWebsitesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsSkuWebsites**](ProductsSkuWebsitesAPI.md#PostV1ProductsSkuWebsites) | **Post** /V1/products/{sku}/websites | products/{sku}/websites
[**PutV1ProductsSkuWebsites**](ProductsSkuWebsitesAPI.md#PutV1ProductsSkuWebsites) | **Put** /V1/products/{sku}/websites | products/{sku}/websites



## PostV1ProductsSkuWebsites

> bool PostV1ProductsSkuWebsites(ctx, sku).PutV1ProductsSkuWebsitesRequest(putV1ProductsSkuWebsitesRequest).Execute()

products/{sku}/websites



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
	putV1ProductsSkuWebsitesRequest := *openapiclient.NewPutV1ProductsSkuWebsitesRequest(*openapiclient.NewCatalogDataProductWebsiteLinkInterface("Sku_example", int32(123))) // PutV1ProductsSkuWebsitesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuWebsitesAPI.PostV1ProductsSkuWebsites(context.Background(), sku).PutV1ProductsSkuWebsitesRequest(putV1ProductsSkuWebsitesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuWebsitesAPI.PostV1ProductsSkuWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSkuWebsites`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuWebsitesAPI.PostV1ProductsSkuWebsites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSkuWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1ProductsSkuWebsitesRequest** | [**PutV1ProductsSkuWebsitesRequest**](PutV1ProductsSkuWebsitesRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ProductsSkuWebsites

> bool PutV1ProductsSkuWebsites(ctx, sku).PutV1ProductsSkuWebsitesRequest(putV1ProductsSkuWebsitesRequest).Execute()

products/{sku}/websites



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
	putV1ProductsSkuWebsitesRequest := *openapiclient.NewPutV1ProductsSkuWebsitesRequest(*openapiclient.NewCatalogDataProductWebsiteLinkInterface("Sku_example", int32(123))) // PutV1ProductsSkuWebsitesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuWebsitesAPI.PutV1ProductsSkuWebsites(context.Background(), sku).PutV1ProductsSkuWebsitesRequest(putV1ProductsSkuWebsitesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuWebsitesAPI.PutV1ProductsSkuWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsSkuWebsites`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuWebsitesAPI.PutV1ProductsSkuWebsites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsSkuWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1ProductsSkuWebsitesRequest** | [**PutV1ProductsSkuWebsitesRequest**](PutV1ProductsSkuWebsitesRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

