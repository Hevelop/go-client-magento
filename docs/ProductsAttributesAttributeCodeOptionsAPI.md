# \ProductsAttributesAttributeCodeOptionsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsAttributesAttributeCodeOptions**](ProductsAttributesAttributeCodeOptionsAPI.md#GetV1ProductsAttributesAttributeCodeOptions) | **Get** /V1/products/attributes/{attributeCode}/options | products/attributes/{attributeCode}/options
[**PostV1ProductsAttributesAttributeCodeOptions**](ProductsAttributesAttributeCodeOptionsAPI.md#PostV1ProductsAttributesAttributeCodeOptions) | **Post** /V1/products/attributes/{attributeCode}/options | products/attributes/{attributeCode}/options



## GetV1ProductsAttributesAttributeCodeOptions

> []EavDataAttributeOptionInterface GetV1ProductsAttributesAttributeCodeOptions(ctx, attributeCode).Execute()

products/attributes/{attributeCode}/options



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
	resp, r, err := apiClient.ProductsAttributesAttributeCodeOptionsAPI.GetV1ProductsAttributesAttributeCodeOptions(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeOptionsAPI.GetV1ProductsAttributesAttributeCodeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsAttributesAttributeCodeOptions`: []EavDataAttributeOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeOptionsAPI.GetV1ProductsAttributesAttributeCodeOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsAttributesAttributeCodeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EavDataAttributeOptionInterface**](EavDataAttributeOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ProductsAttributesAttributeCodeOptions

> string PostV1ProductsAttributesAttributeCodeOptions(ctx, attributeCode).PostV1ProductsAttributesAttributeCodeOptionsRequest(postV1ProductsAttributesAttributeCodeOptionsRequest).Execute()

products/attributes/{attributeCode}/options



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
	postV1ProductsAttributesAttributeCodeOptionsRequest := *openapiclient.NewPostV1ProductsAttributesAttributeCodeOptionsRequest(*openapiclient.NewEavDataAttributeOptionInterface("Label_example", "Value_example")) // PostV1ProductsAttributesAttributeCodeOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeOptionsAPI.PostV1ProductsAttributesAttributeCodeOptions(context.Background(), attributeCode).PostV1ProductsAttributesAttributeCodeOptionsRequest(postV1ProductsAttributesAttributeCodeOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeOptionsAPI.PostV1ProductsAttributesAttributeCodeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsAttributesAttributeCodeOptions`: string
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeOptionsAPI.PostV1ProductsAttributesAttributeCodeOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsAttributesAttributeCodeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsAttributesAttributeCodeOptionsRequest** | [**PostV1ProductsAttributesAttributeCodeOptionsRequest**](PostV1ProductsAttributesAttributeCodeOptionsRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

