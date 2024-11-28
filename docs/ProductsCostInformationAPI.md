# \ProductsCostInformationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsCostinformation**](ProductsCostInformationAPI.md#PostV1ProductsCostinformation) | **Post** /V1/products/cost-information | products/cost-information



## PostV1ProductsCostinformation

> []CatalogDataCostInterface PostV1ProductsCostinformation(ctx).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()

products/cost-information



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
	postV1ProductsBasepricesinformationRequest := *openapiclient.NewPostV1ProductsBasepricesinformationRequest([]string{"Skus_example"}) // PostV1ProductsBasepricesinformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsCostInformationAPI.PostV1ProductsCostinformation(context.Background()).PostV1ProductsBasepricesinformationRequest(postV1ProductsBasepricesinformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsCostInformationAPI.PostV1ProductsCostinformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsCostinformation`: []CatalogDataCostInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsCostInformationAPI.PostV1ProductsCostinformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsCostinformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsBasepricesinformationRequest** | [**PostV1ProductsBasepricesinformationRequest**](PostV1ProductsBasepricesinformationRequest.md) |  | 

### Return type

[**[]CatalogDataCostInterface**](CatalogDataCostInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

