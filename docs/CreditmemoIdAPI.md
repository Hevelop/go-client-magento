# \CreditmemoIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CreditmemoId**](CreditmemoIdAPI.md#GetV1CreditmemoId) | **Get** /V1/creditmemo/{id} | creditmemo/{id}
[**PutV1CreditmemoId**](CreditmemoIdAPI.md#PutV1CreditmemoId) | **Put** /V1/creditmemo/{id} | creditmemo/{id}



## GetV1CreditmemoId

> SalesDataCreditmemoInterface GetV1CreditmemoId(ctx, id).Execute()

creditmemo/{id}



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
	id := int32(56) // int32 | The credit memo ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoIdAPI.GetV1CreditmemoId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoIdAPI.GetV1CreditmemoId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CreditmemoId`: SalesDataCreditmemoInterface
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoIdAPI.GetV1CreditmemoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The credit memo ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CreditmemoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataCreditmemoInterface**](SalesDataCreditmemoInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CreditmemoId

> bool PutV1CreditmemoId(ctx, id).Execute()

creditmemo/{id}



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
	id := int32(56) // int32 | The credit memo ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoIdAPI.PutV1CreditmemoId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoIdAPI.PutV1CreditmemoId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CreditmemoId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoIdAPI.PutV1CreditmemoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The credit memo ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CreditmemoIdRequest struct via the builder pattern


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

