# \EavAttributeSetsAttributeSetIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1EavAttributesetsAttributeSetId**](EavAttributeSetsAttributeSetIdAPI.md#DeleteV1EavAttributesetsAttributeSetId) | **Delete** /V1/eav/attribute-sets/{attributeSetId} | eav/attribute-sets/{attributeSetId}
[**GetV1EavAttributesetsAttributeSetId**](EavAttributeSetsAttributeSetIdAPI.md#GetV1EavAttributesetsAttributeSetId) | **Get** /V1/eav/attribute-sets/{attributeSetId} | eav/attribute-sets/{attributeSetId}
[**PutV1EavAttributesetsAttributeSetId**](EavAttributeSetsAttributeSetIdAPI.md#PutV1EavAttributesetsAttributeSetId) | **Put** /V1/eav/attribute-sets/{attributeSetId} | eav/attribute-sets/{attributeSetId}



## DeleteV1EavAttributesetsAttributeSetId

> bool DeleteV1EavAttributesetsAttributeSetId(ctx, attributeSetId).Execute()

eav/attribute-sets/{attributeSetId}



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
	resp, r, err := apiClient.EavAttributeSetsAttributeSetIdAPI.DeleteV1EavAttributesetsAttributeSetId(context.Background(), attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EavAttributeSetsAttributeSetIdAPI.DeleteV1EavAttributesetsAttributeSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1EavAttributesetsAttributeSetId`: bool
	fmt.Fprintf(os.Stdout, "Response from `EavAttributeSetsAttributeSetIdAPI.DeleteV1EavAttributesetsAttributeSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1EavAttributesetsAttributeSetIdRequest struct via the builder pattern


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


## GetV1EavAttributesetsAttributeSetId

> EavDataAttributeSetInterface GetV1EavAttributesetsAttributeSetId(ctx, attributeSetId).Execute()

eav/attribute-sets/{attributeSetId}



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
	resp, r, err := apiClient.EavAttributeSetsAttributeSetIdAPI.GetV1EavAttributesetsAttributeSetId(context.Background(), attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EavAttributeSetsAttributeSetIdAPI.GetV1EavAttributesetsAttributeSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1EavAttributesetsAttributeSetId`: EavDataAttributeSetInterface
	fmt.Fprintf(os.Stdout, "Response from `EavAttributeSetsAttributeSetIdAPI.GetV1EavAttributesetsAttributeSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1EavAttributesetsAttributeSetIdRequest struct via the builder pattern


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


## PutV1EavAttributesetsAttributeSetId

> EavDataAttributeSetInterface PutV1EavAttributesetsAttributeSetId(ctx, attributeSetId).PutV1EavAttributesetsAttributeSetIdRequest(putV1EavAttributesetsAttributeSetIdRequest).Execute()

eav/attribute-sets/{attributeSetId}



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
	resp, r, err := apiClient.EavAttributeSetsAttributeSetIdAPI.PutV1EavAttributesetsAttributeSetId(context.Background(), attributeSetId).PutV1EavAttributesetsAttributeSetIdRequest(putV1EavAttributesetsAttributeSetIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EavAttributeSetsAttributeSetIdAPI.PutV1EavAttributesetsAttributeSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1EavAttributesetsAttributeSetId`: EavDataAttributeSetInterface
	fmt.Fprintf(os.Stdout, "Response from `EavAttributeSetsAttributeSetIdAPI.PutV1EavAttributesetsAttributeSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1EavAttributesetsAttributeSetIdRequest struct via the builder pattern


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

