# \ProductsAttributesAttributeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsAttributesAttributeCode**](ProductsAttributesAttributeCodeAPI.md#DeleteV1ProductsAttributesAttributeCode) | **Delete** /V1/products/attributes/{attributeCode} | products/attributes/{attributeCode}
[**GetV1ProductsAttributesAttributeCode**](ProductsAttributesAttributeCodeAPI.md#GetV1ProductsAttributesAttributeCode) | **Get** /V1/products/attributes/{attributeCode} | products/attributes/{attributeCode}
[**PutV1ProductsAttributesAttributeCode**](ProductsAttributesAttributeCodeAPI.md#PutV1ProductsAttributesAttributeCode) | **Put** /V1/products/attributes/{attributeCode} | products/attributes/{attributeCode}



## DeleteV1ProductsAttributesAttributeCode

> bool DeleteV1ProductsAttributesAttributeCode(ctx, attributeCode).Execute()

products/attributes/{attributeCode}



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
	attributeCode := "attributeCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeAPI.DeleteV1ProductsAttributesAttributeCode(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeAPI.DeleteV1ProductsAttributesAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsAttributesAttributeCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeAPI.DeleteV1ProductsAttributesAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsAttributesAttributeCodeRequest struct via the builder pattern


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


## GetV1ProductsAttributesAttributeCode

> CatalogDataProductAttributeInterface GetV1ProductsAttributesAttributeCode(ctx, attributeCode).Execute()

products/attributes/{attributeCode}



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
	attributeCode := "attributeCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeAPI.GetV1ProductsAttributesAttributeCode(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeAPI.GetV1ProductsAttributesAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsAttributesAttributeCode`: CatalogDataProductAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeAPI.GetV1ProductsAttributesAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsAttributesAttributeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CatalogDataProductAttributeInterface**](CatalogDataProductAttributeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ProductsAttributesAttributeCode

> CatalogDataProductAttributeInterface PutV1ProductsAttributesAttributeCode(ctx, attributeCode).PostV1ProductsAttributesRequest(postV1ProductsAttributesRequest).Execute()

products/attributes/{attributeCode}



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
	attributeCode := "attributeCode_example" // string | 
	postV1ProductsAttributesRequest := *openapiclient.NewPostV1ProductsAttributesRequest(*openapiclient.NewCatalogDataProductAttributeInterface("AttributeCode_example", "FrontendInput_example", "EntityTypeId_example", false, []openapiclient.EavDataAttributeFrontendLabelInterface{*openapiclient.NewEavDataAttributeFrontendLabelInterface()})) // PostV1ProductsAttributesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeAPI.PutV1ProductsAttributesAttributeCode(context.Background(), attributeCode).PostV1ProductsAttributesRequest(postV1ProductsAttributesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeAPI.PutV1ProductsAttributesAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsAttributesAttributeCode`: CatalogDataProductAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeAPI.PutV1ProductsAttributesAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsAttributesAttributeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsAttributesRequest** | [**PostV1ProductsAttributesRequest**](PostV1ProductsAttributesRequest.md) |  | 

### Return type

[**CatalogDataProductAttributeInterface**](CatalogDataProductAttributeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

