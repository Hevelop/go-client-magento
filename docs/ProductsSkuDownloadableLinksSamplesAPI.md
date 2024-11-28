# \ProductsSkuDownloadableLinksSamplesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsSkuDownloadablelinksSamples**](ProductsSkuDownloadableLinksSamplesAPI.md#GetV1ProductsSkuDownloadablelinksSamples) | **Get** /V1/products/{sku}/downloadable-links/samples | products/{sku}/downloadable-links/samples
[**PostV1ProductsSkuDownloadablelinksSamples**](ProductsSkuDownloadableLinksSamplesAPI.md#PostV1ProductsSkuDownloadablelinksSamples) | **Post** /V1/products/{sku}/downloadable-links/samples | products/{sku}/downloadable-links/samples



## GetV1ProductsSkuDownloadablelinksSamples

> []DownloadableDataSampleInterface GetV1ProductsSkuDownloadablelinksSamples(ctx, sku).Execute()

products/{sku}/downloadable-links/samples



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
	resp, r, err := apiClient.ProductsSkuDownloadableLinksSamplesAPI.GetV1ProductsSkuDownloadablelinksSamples(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuDownloadableLinksSamplesAPI.GetV1ProductsSkuDownloadablelinksSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuDownloadablelinksSamples`: []DownloadableDataSampleInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuDownloadableLinksSamplesAPI.GetV1ProductsSkuDownloadablelinksSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuDownloadablelinksSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DownloadableDataSampleInterface**](DownloadableDataSampleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ProductsSkuDownloadablelinksSamples

> int32 PostV1ProductsSkuDownloadablelinksSamples(ctx, sku).PostV1ProductsSkuDownloadablelinksSamplesRequest(postV1ProductsSkuDownloadablelinksSamplesRequest).Execute()

products/{sku}/downloadable-links/samples



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
	postV1ProductsSkuDownloadablelinksSamplesRequest := *openapiclient.NewPostV1ProductsSkuDownloadablelinksSamplesRequest(*openapiclient.NewDownloadableDataSampleInterface("Title_example", int32(123), "SampleType_example")) // PostV1ProductsSkuDownloadablelinksSamplesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuDownloadableLinksSamplesAPI.PostV1ProductsSkuDownloadablelinksSamples(context.Background(), sku).PostV1ProductsSkuDownloadablelinksSamplesRequest(postV1ProductsSkuDownloadablelinksSamplesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuDownloadableLinksSamplesAPI.PostV1ProductsSkuDownloadablelinksSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsSkuDownloadablelinksSamples`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuDownloadableLinksSamplesAPI.PostV1ProductsSkuDownloadablelinksSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsSkuDownloadablelinksSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsSkuDownloadablelinksSamplesRequest** | [**PostV1ProductsSkuDownloadablelinksSamplesRequest**](PostV1ProductsSkuDownloadablelinksSamplesRequest.md) |  | 

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

