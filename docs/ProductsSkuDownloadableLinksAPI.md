# \ProductsSkuDownloadableLinksAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsSkuDownloadablelinks**](ProductsSkuDownloadableLinksAPI.md#GetV1ProductsSkuDownloadablelinks) | **Get** /V1/products/{sku}/downloadable-links | products/{sku}/downloadable-links
[**PostV1ProductsSkuDownloadablelinks**](ProductsSkuDownloadableLinksAPI.md#PostV1ProductsSkuDownloadablelinks) | **Post** /V1/products/{sku}/downloadable-links | products/{sku}/downloadable-links



## GetV1ProductsSkuDownloadablelinks

> []DownloadableDataLinkInterface GetV1ProductsSkuDownloadablelinks(ctx, sku).Execute()

products/{sku}/downloadable-links



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
	resp, r, err := apiClient.ProductsSkuDownloadableLinksAPI.GetV1ProductsSkuDownloadablelinks(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuDownloadableLinksAPI.GetV1ProductsSkuDownloadablelinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuDownloadablelinks`: []DownloadableDataLinkInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuDownloadableLinksAPI.GetV1ProductsSkuDownloadablelinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuDownloadablelinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DownloadableDataLinkInterface**](DownloadableDataLinkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ProductsSkuDownloadablelinks

> int32 PostV1ProductsSkuDownloadablelinks(ctx, sku).PostV1ProductsSkuDownloadablelinksRequest(postV1ProductsSkuDownloadablelinksRequest).Execute()

products/{sku}/downloadable-links



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
	postV1ProductsSkuDownloadablelinksRequest := *openapiclient.NewPostV1ProductsSkuDownloadablelinksRequest(*openapiclient.NewDownloadableDataLinkInterface(int32(123), int32(123), float32(123), "LinkType_example", "SampleType_example")) // PostV1ProductsSkuDownloadablelinksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuDownloadableLinksAPI.PostV1ProductsSkuDownloadablelinks(context.Background(), sku).PostV1ProductsSkuDownloadablelinksRequest(postV1ProductsSkuDownloadablelinksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuDownloadableLinksAPI.PostV1ProductsSkuDownloadablelinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSkuDownloadablelinks`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuDownloadableLinksAPI.PostV1ProductsSkuDownloadablelinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSkuDownloadablelinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsSkuDownloadablelinksRequest** | [**PostV1ProductsSkuDownloadablelinksRequest**](PostV1ProductsSkuDownloadablelinksRequest.md) |  | 

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

