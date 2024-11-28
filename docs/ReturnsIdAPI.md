# \ReturnsIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ReturnsId**](ReturnsIdAPI.md#DeleteV1ReturnsId) | **Delete** /V1/returns/{id} | returns/{id}
[**GetV1ReturnsId**](ReturnsIdAPI.md#GetV1ReturnsId) | **Get** /V1/returns/{id} | returns/{id}
[**PutV1ReturnsId**](ReturnsIdAPI.md#PutV1ReturnsId) | **Put** /V1/returns/{id} | returns/{id}



## DeleteV1ReturnsId

> bool DeleteV1ReturnsId(ctx, id).PostV1ReturnsRequest(postV1ReturnsRequest).Execute()

returns/{id}



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
	postV1ReturnsRequest := *openapiclient.NewPostV1ReturnsRequest(*openapiclient.NewRmaDataRmaInterface("IncrementId_example", int32(123), int32(123), "OrderIncrementId_example", int32(123), int32(123), "DateRequested_example", "CustomerCustomEmail_example", []openapiclient.RmaDataItemInterface{*openapiclient.NewRmaDataItemInterface(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "Reason_example", "Condition_example", "Resolution_example", "Status_example")}, "Status_example", []openapiclient.RmaDataCommentInterface{*openapiclient.NewRmaDataCommentInterface("Comment_example", int32(123), "CreatedAt_example", int32(123), false, false, "Status_example", false)}, []openapiclient.RmaDataTrackInterface{*openapiclient.NewRmaDataTrackInterface(int32(123), int32(123), "TrackNumber_example", "CarrierTitle_example", "CarrierCode_example")})) // PostV1ReturnsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsIdAPI.DeleteV1ReturnsId(context.Background(), id).PostV1ReturnsRequest(postV1ReturnsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdAPI.DeleteV1ReturnsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ReturnsId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdAPI.DeleteV1ReturnsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ReturnsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ReturnsRequest** | [**PostV1ReturnsRequest**](PostV1ReturnsRequest.md) |  | 

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


## GetV1ReturnsId

> RmaDataRmaInterface GetV1ReturnsId(ctx, id).Execute()

returns/{id}



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
	resp, r, err := apiClient.ReturnsIdAPI.GetV1ReturnsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdAPI.GetV1ReturnsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsId`: RmaDataRmaInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdAPI.GetV1ReturnsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RmaDataRmaInterface**](RmaDataRmaInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ReturnsId

> RmaDataRmaInterface PutV1ReturnsId(ctx, id).PostV1ReturnsRequest(postV1ReturnsRequest).Execute()

returns/{id}



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
	postV1ReturnsRequest := *openapiclient.NewPostV1ReturnsRequest(*openapiclient.NewRmaDataRmaInterface("IncrementId_example", int32(123), int32(123), "OrderIncrementId_example", int32(123), int32(123), "DateRequested_example", "CustomerCustomEmail_example", []openapiclient.RmaDataItemInterface{*openapiclient.NewRmaDataItemInterface(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "Reason_example", "Condition_example", "Resolution_example", "Status_example")}, "Status_example", []openapiclient.RmaDataCommentInterface{*openapiclient.NewRmaDataCommentInterface("Comment_example", int32(123), "CreatedAt_example", int32(123), false, false, "Status_example", false)}, []openapiclient.RmaDataTrackInterface{*openapiclient.NewRmaDataTrackInterface(int32(123), int32(123), "TrackNumber_example", "CarrierTitle_example", "CarrierCode_example")})) // PostV1ReturnsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsIdAPI.PutV1ReturnsId(context.Background(), id).PostV1ReturnsRequest(postV1ReturnsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdAPI.PutV1ReturnsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ReturnsId`: RmaDataRmaInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdAPI.PutV1ReturnsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ReturnsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ReturnsRequest** | [**PostV1ReturnsRequest**](PostV1ReturnsRequest.md) |  | 

### Return type

[**RmaDataRmaInterface**](RmaDataRmaInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

