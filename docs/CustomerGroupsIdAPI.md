# \CustomerGroupsIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CustomerGroupsId**](CustomerGroupsIdAPI.md#DeleteV1CustomerGroupsId) | **Delete** /V1/customerGroups/{id} | customerGroups/{id}
[**GetV1CustomerGroupsId**](CustomerGroupsIdAPI.md#GetV1CustomerGroupsId) | **Get** /V1/customerGroups/{id} | customerGroups/{id}
[**PutV1CustomerGroupsId**](CustomerGroupsIdAPI.md#PutV1CustomerGroupsId) | **Put** /V1/customerGroups/{id} | customerGroups/{id}



## DeleteV1CustomerGroupsId

> bool DeleteV1CustomerGroupsId(ctx, id).Execute()

customerGroups/{id}



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerGroupsIdAPI.DeleteV1CustomerGroupsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsIdAPI.DeleteV1CustomerGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CustomerGroupsId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsIdAPI.DeleteV1CustomerGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CustomerGroupsIdRequest struct via the builder pattern


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


## GetV1CustomerGroupsId

> CustomerDataGroupInterface GetV1CustomerGroupsId(ctx, id).Execute()

customerGroups/{id}



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerGroupsIdAPI.GetV1CustomerGroupsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsIdAPI.GetV1CustomerGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomerGroupsId`: CustomerDataGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsIdAPI.GetV1CustomerGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomerGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDataGroupInterface**](CustomerDataGroupInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CustomerGroupsId

> CustomerDataGroupInterface PutV1CustomerGroupsId(ctx, id).PostV1CustomerGroupsRequest(postV1CustomerGroupsRequest).Execute()

customerGroups/{id}



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
	id := "id_example" // string | 
	postV1CustomerGroupsRequest := *openapiclient.NewPostV1CustomerGroupsRequest(*openapiclient.NewCustomerDataGroupInterface("Code_example", int32(123))) // PostV1CustomerGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerGroupsIdAPI.PutV1CustomerGroupsId(context.Background(), id).PostV1CustomerGroupsRequest(postV1CustomerGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsIdAPI.PutV1CustomerGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomerGroupsId`: CustomerDataGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsIdAPI.PutV1CustomerGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomerGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CustomerGroupsRequest** | [**PostV1CustomerGroupsRequest**](PostV1CustomerGroupsRequest.md) |  | 

### Return type

[**CustomerDataGroupInterface**](CustomerDataGroupInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

