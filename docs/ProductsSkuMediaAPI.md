# \ProductsSkuMediaAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsSkuMedia**](ProductsSkuMediaAPI.md#GetV1ProductsSkuMedia) | **Get** /V1/products/{sku}/media | products/{sku}/media
[**PostV1ProductsSkuMedia**](ProductsSkuMediaAPI.md#PostV1ProductsSkuMedia) | **Post** /V1/products/{sku}/media | products/{sku}/media



## GetV1ProductsSkuMedia

> []CatalogDataProductAttributeMediaGalleryEntryInterface GetV1ProductsSkuMedia(ctx, sku).Execute()

products/{sku}/media



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuMediaAPI.GetV1ProductsSkuMedia(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuMediaAPI.GetV1ProductsSkuMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuMedia`: []CatalogDataProductAttributeMediaGalleryEntryInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuMediaAPI.GetV1ProductsSkuMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogDataProductAttributeMediaGalleryEntryInterface**](CatalogDataProductAttributeMediaGalleryEntryInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ProductsSkuMedia

> int32 PostV1ProductsSkuMedia(ctx, sku).PostV1ProductsSkuMediaRequest(postV1ProductsSkuMediaRequest).Execute()

products/{sku}/media



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
	postV1ProductsSkuMediaRequest := *openapiclient.NewPostV1ProductsSkuMediaRequest(*openapiclient.NewCatalogDataProductAttributeMediaGalleryEntryInterface("MediaType_example", "Label_example", int32(123), false, []string{"Types_example"})) // PostV1ProductsSkuMediaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuMediaAPI.PostV1ProductsSkuMedia(context.Background(), sku).PostV1ProductsSkuMediaRequest(postV1ProductsSkuMediaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuMediaAPI.PostV1ProductsSkuMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSkuMedia`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuMediaAPI.PostV1ProductsSkuMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSkuMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsSkuMediaRequest** | [**PostV1ProductsSkuMediaRequest**](PostV1ProductsSkuMediaRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

