# \ProductsSkuLinksAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsSkuLinks**](ProductsSkuLinksAPI.md#PostV1ProductsSkuLinks) | **Post** /V1/products/{sku}/links | products/{sku}/links
[**PutV1ProductsSkuLinks**](ProductsSkuLinksAPI.md#PutV1ProductsSkuLinks) | **Put** /V1/products/{sku}/links | products/{sku}/links



## PostV1ProductsSkuLinks

> bool PostV1ProductsSkuLinks(ctx, sku).PostV1ProductsSkuLinksRequest(postV1ProductsSkuLinksRequest).Execute()

products/{sku}/links



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
	postV1ProductsSkuLinksRequest := *openapiclient.NewPostV1ProductsSkuLinksRequest([]openapiclient.CatalogDataProductLinkInterface{*openapiclient.NewCatalogDataProductLinkInterface("Sku_example", "LinkType_example", "LinkedProductSku_example", "LinkedProductType_example", int32(123))}) // PostV1ProductsSkuLinksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuLinksAPI.PostV1ProductsSkuLinks(context.Background(), sku).PostV1ProductsSkuLinksRequest(postV1ProductsSkuLinksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuLinksAPI.PostV1ProductsSkuLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSkuLinks`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuLinksAPI.PostV1ProductsSkuLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSkuLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsSkuLinksRequest** | [**PostV1ProductsSkuLinksRequest**](PostV1ProductsSkuLinksRequest.md) |  | 

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


## PutV1ProductsSkuLinks

> bool PutV1ProductsSkuLinks(ctx, sku).PutV1ProductsSkuLinksRequest(putV1ProductsSkuLinksRequest).Execute()

products/{sku}/links



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
	putV1ProductsSkuLinksRequest := *openapiclient.NewPutV1ProductsSkuLinksRequest(*openapiclient.NewCatalogDataProductLinkInterface("Sku_example", "LinkType_example", "LinkedProductSku_example", "LinkedProductType_example", int32(123))) // PutV1ProductsSkuLinksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuLinksAPI.PutV1ProductsSkuLinks(context.Background(), sku).PutV1ProductsSkuLinksRequest(putV1ProductsSkuLinksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuLinksAPI.PutV1ProductsSkuLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsSkuLinks`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuLinksAPI.PutV1ProductsSkuLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsSkuLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1ProductsSkuLinksRequest** | [**PutV1ProductsSkuLinksRequest**](PutV1ProductsSkuLinksRequest.md) |  | 

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

