# \ProductsSkuMediaEntryIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsSkuMediaEntryId**](ProductsSkuMediaEntryIdAPI.md#DeleteV1ProductsSkuMediaEntryId) | **Delete** /V1/products/{sku}/media/{entryId} | products/{sku}/media/{entryId}
[**GetV1ProductsSkuMediaEntryId**](ProductsSkuMediaEntryIdAPI.md#GetV1ProductsSkuMediaEntryId) | **Get** /V1/products/{sku}/media/{entryId} | products/{sku}/media/{entryId}
[**PutV1ProductsSkuMediaEntryId**](ProductsSkuMediaEntryIdAPI.md#PutV1ProductsSkuMediaEntryId) | **Put** /V1/products/{sku}/media/{entryId} | products/{sku}/media/{entryId}



## DeleteV1ProductsSkuMediaEntryId

> bool DeleteV1ProductsSkuMediaEntryId(ctx, sku, entryId).Execute()

products/{sku}/media/{entryId}



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
	entryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuMediaEntryIdAPI.DeleteV1ProductsSkuMediaEntryId(context.Background(), sku, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuMediaEntryIdAPI.DeleteV1ProductsSkuMediaEntryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsSkuMediaEntryId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuMediaEntryIdAPI.DeleteV1ProductsSkuMediaEntryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**entryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsSkuMediaEntryIdRequest struct via the builder pattern


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


## GetV1ProductsSkuMediaEntryId

> CatalogDataProductAttributeMediaGalleryEntryInterface GetV1ProductsSkuMediaEntryId(ctx, sku, entryId).Execute()

products/{sku}/media/{entryId}



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
	entryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuMediaEntryIdAPI.GetV1ProductsSkuMediaEntryId(context.Background(), sku, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuMediaEntryIdAPI.GetV1ProductsSkuMediaEntryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuMediaEntryId`: CatalogDataProductAttributeMediaGalleryEntryInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuMediaEntryIdAPI.GetV1ProductsSkuMediaEntryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**entryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuMediaEntryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CatalogDataProductAttributeMediaGalleryEntryInterface**](CatalogDataProductAttributeMediaGalleryEntryInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ProductsSkuMediaEntryId

> bool PutV1ProductsSkuMediaEntryId(ctx, sku, entryId).PostV1ProductsSkuMediaRequest(postV1ProductsSkuMediaRequest).Execute()

products/{sku}/media/{entryId}



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
	entryId := "entryId_example" // string | 
	postV1ProductsSkuMediaRequest := *openapiclient.NewPostV1ProductsSkuMediaRequest(*openapiclient.NewCatalogDataProductAttributeMediaGalleryEntryInterface("MediaType_example", "Label_example", int32(123), false, []string{"Types_example"})) // PostV1ProductsSkuMediaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuMediaEntryIdAPI.PutV1ProductsSkuMediaEntryId(context.Background(), sku, entryId).PostV1ProductsSkuMediaRequest(postV1ProductsSkuMediaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuMediaEntryIdAPI.PutV1ProductsSkuMediaEntryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsSkuMediaEntryId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuMediaEntryIdAPI.PutV1ProductsSkuMediaEntryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsSkuMediaEntryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postV1ProductsSkuMediaRequest** | [**PostV1ProductsSkuMediaRequest**](PostV1ProductsSkuMediaRequest.md) |  | 

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

