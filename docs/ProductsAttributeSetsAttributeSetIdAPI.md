# \ProductsAttributeSetsAttributeSetIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsAttributesetsAttributeSetId**](ProductsAttributeSetsAttributeSetIdAPI.md#DeleteV1ProductsAttributesetsAttributeSetId) | **Delete** /V1/products/attribute-sets/{attributeSetId} | products/attribute-sets/{attributeSetId}
[**GetV1ProductsAttributesetsAttributeSetId**](ProductsAttributeSetsAttributeSetIdAPI.md#GetV1ProductsAttributesetsAttributeSetId) | **Get** /V1/products/attribute-sets/{attributeSetId} | products/attribute-sets/{attributeSetId}
[**PutV1ProductsAttributesetsAttributeSetId**](ProductsAttributeSetsAttributeSetIdAPI.md#PutV1ProductsAttributesetsAttributeSetId) | **Put** /V1/products/attribute-sets/{attributeSetId} | products/attribute-sets/{attributeSetId}



## DeleteV1ProductsAttributesetsAttributeSetId

> bool DeleteV1ProductsAttributesetsAttributeSetId(ctx, attributeSetId).Execute()

products/attribute-sets/{attributeSetId}



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
	attributeSetId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAttributeSetIdAPI.DeleteV1ProductsAttributesetsAttributeSetId(context.Background(), attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAttributeSetIdAPI.DeleteV1ProductsAttributesetsAttributeSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsAttributesetsAttributeSetId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAttributeSetIdAPI.DeleteV1ProductsAttributesetsAttributeSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsAttributesetsAttributeSetIdRequest struct via the builder pattern


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


## GetV1ProductsAttributesetsAttributeSetId

> EavDataAttributeSetInterface GetV1ProductsAttributesetsAttributeSetId(ctx, attributeSetId).Execute()

products/attribute-sets/{attributeSetId}



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
	attributeSetId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAttributeSetIdAPI.GetV1ProductsAttributesetsAttributeSetId(context.Background(), attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAttributeSetIdAPI.GetV1ProductsAttributesetsAttributeSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsAttributesetsAttributeSetId`: EavDataAttributeSetInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAttributeSetIdAPI.GetV1ProductsAttributesetsAttributeSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsAttributesetsAttributeSetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EavDataAttributeSetInterface**](EavDataAttributeSetInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ProductsAttributesetsAttributeSetId

> EavDataAttributeSetInterface PutV1ProductsAttributesetsAttributeSetId(ctx, attributeSetId).PutV1EavAttributesetsAttributeSetIdRequest(putV1EavAttributesetsAttributeSetIdRequest).Execute()

products/attribute-sets/{attributeSetId}



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
	attributeSetId := "attributeSetId_example" // string | 
	putV1EavAttributesetsAttributeSetIdRequest := *openapiclient.NewPutV1EavAttributesetsAttributeSetIdRequest(*openapiclient.NewEavDataAttributeSetInterface("AttributeSetName_example", int32(123))) // PutV1EavAttributesetsAttributeSetIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAttributeSetIdAPI.PutV1ProductsAttributesetsAttributeSetId(context.Background(), attributeSetId).PutV1EavAttributesetsAttributeSetIdRequest(putV1EavAttributesetsAttributeSetIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAttributeSetIdAPI.PutV1ProductsAttributesetsAttributeSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsAttributesetsAttributeSetId`: EavDataAttributeSetInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAttributeSetIdAPI.PutV1ProductsAttributesetsAttributeSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsAttributesetsAttributeSetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1EavAttributesetsAttributeSetIdRequest** | [**PutV1EavAttributesetsAttributeSetIdRequest**](PutV1EavAttributesetsAttributeSetIdRequest.md) |  | 

### Return type

[**EavDataAttributeSetInterface**](EavDataAttributeSetInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

