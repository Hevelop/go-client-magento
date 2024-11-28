# \ProductsAttributesAttributeCodeOptionsOptionIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsAttributesAttributeCodeOptionsOptionId**](ProductsAttributesAttributeCodeOptionsOptionIdAPI.md#DeleteV1ProductsAttributesAttributeCodeOptionsOptionId) | **Delete** /V1/products/attributes/{attributeCode}/options/{optionId} | products/attributes/{attributeCode}/options/{optionId}
[**PutV1ProductsAttributesAttributeCodeOptionsOptionId**](ProductsAttributesAttributeCodeOptionsOptionIdAPI.md#PutV1ProductsAttributesAttributeCodeOptionsOptionId) | **Put** /V1/products/attributes/{attributeCode}/options/{optionId} | products/attributes/{attributeCode}/options/{optionId}



## DeleteV1ProductsAttributesAttributeCodeOptionsOptionId

> bool DeleteV1ProductsAttributesAttributeCodeOptionsOptionId(ctx, attributeCode, optionId).Execute()

products/attributes/{attributeCode}/options/{optionId}



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
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeOptionsOptionIdAPI.DeleteV1ProductsAttributesAttributeCodeOptionsOptionId(context.Background(), attributeCode, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeOptionsOptionIdAPI.DeleteV1ProductsAttributesAttributeCodeOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsAttributesAttributeCodeOptionsOptionId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeOptionsOptionIdAPI.DeleteV1ProductsAttributesAttributeCodeOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsAttributesAttributeCodeOptionsOptionIdRequest struct via the builder pattern


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


## PutV1ProductsAttributesAttributeCodeOptionsOptionId

> bool PutV1ProductsAttributesAttributeCodeOptionsOptionId(ctx, attributeCode, optionId).PostV1ProductsAttributesAttributeCodeOptionsRequest(postV1ProductsAttributesAttributeCodeOptionsRequest).Execute()

products/attributes/{attributeCode}/options/{optionId}



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
	optionId := int32(56) // int32 | 
	postV1ProductsAttributesAttributeCodeOptionsRequest := *openapiclient.NewPostV1ProductsAttributesAttributeCodeOptionsRequest(*openapiclient.NewEavDataAttributeOptionInterface("Label_example", "Value_example")) // PostV1ProductsAttributesAttributeCodeOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeOptionsOptionIdAPI.PutV1ProductsAttributesAttributeCodeOptionsOptionId(context.Background(), attributeCode, optionId).PostV1ProductsAttributesAttributeCodeOptionsRequest(postV1ProductsAttributesAttributeCodeOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeOptionsOptionIdAPI.PutV1ProductsAttributesAttributeCodeOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsAttributesAttributeCodeOptionsOptionId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeOptionsOptionIdAPI.PutV1ProductsAttributesAttributeCodeOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsAttributesAttributeCodeOptionsOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postV1ProductsAttributesAttributeCodeOptionsRequest** | [**PostV1ProductsAttributesAttributeCodeOptionsRequest**](PostV1ProductsAttributesAttributeCodeOptionsRequest.md) |  | 

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

