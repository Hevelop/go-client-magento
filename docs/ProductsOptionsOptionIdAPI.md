# \ProductsOptionsOptionIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1ProductsOptionsOptionId**](ProductsOptionsOptionIdAPI.md#PutV1ProductsOptionsOptionId) | **Put** /V1/products/options/{optionId} | products/options/{optionId}



## PutV1ProductsOptionsOptionId

> CatalogDataProductCustomOptionInterface PutV1ProductsOptionsOptionId(ctx, optionId).PostV1ProductsOptionsRequest(postV1ProductsOptionsRequest).Execute()

products/options/{optionId}



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
	optionId := "optionId_example" // string | 
	postV1ProductsOptionsRequest := *openapiclient.NewPostV1ProductsOptionsRequest(*openapiclient.NewCatalogDataProductCustomOptionInterface("ProductSku_example", "Title_example", "Type_example", int32(123), false)) // PostV1ProductsOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsOptionsOptionIdAPI.PutV1ProductsOptionsOptionId(context.Background(), optionId).PostV1ProductsOptionsRequest(postV1ProductsOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsOptionsOptionIdAPI.PutV1ProductsOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsOptionsOptionId`: CatalogDataProductCustomOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsOptionsOptionIdAPI.PutV1ProductsOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsOptionsOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsOptionsRequest** | [**PostV1ProductsOptionsRequest**](PostV1ProductsOptionsRequest.md) |  | 

### Return type

[**CatalogDataProductCustomOptionInterface**](CatalogDataProductCustomOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

