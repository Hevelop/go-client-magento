# \ReturnsIdTrackingNumbersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReturnsIdTrackingnumbers**](ReturnsIdTrackingNumbersAPI.md#GetV1ReturnsIdTrackingnumbers) | **Get** /V1/returns/{id}/tracking-numbers | returns/{id}/tracking-numbers
[**PostV1ReturnsIdTrackingnumbers**](ReturnsIdTrackingNumbersAPI.md#PostV1ReturnsIdTrackingnumbers) | **Post** /V1/returns/{id}/tracking-numbers | returns/{id}/tracking-numbers



## GetV1ReturnsIdTrackingnumbers

> RmaDataTrackSearchResultInterface GetV1ReturnsIdTrackingnumbers(ctx, id).Execute()

returns/{id}/tracking-numbers



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
	resp, r, err := apiClient.ReturnsIdTrackingNumbersAPI.GetV1ReturnsIdTrackingnumbers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdTrackingNumbersAPI.GetV1ReturnsIdTrackingnumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsIdTrackingnumbers`: RmaDataTrackSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdTrackingNumbersAPI.GetV1ReturnsIdTrackingnumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsIdTrackingnumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RmaDataTrackSearchResultInterface**](RmaDataTrackSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ReturnsIdTrackingnumbers

> bool PostV1ReturnsIdTrackingnumbers(ctx, id).PostV1ReturnsIdTrackingnumbersRequest(postV1ReturnsIdTrackingnumbersRequest).Execute()

returns/{id}/tracking-numbers



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
	postV1ReturnsIdTrackingnumbersRequest := *openapiclient.NewPostV1ReturnsIdTrackingnumbersRequest(*openapiclient.NewRmaDataTrackInterface(int32(123), int32(123), "TrackNumber_example", "CarrierTitle_example", "CarrierCode_example")) // PostV1ReturnsIdTrackingnumbersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsIdTrackingNumbersAPI.PostV1ReturnsIdTrackingnumbers(context.Background(), id).PostV1ReturnsIdTrackingnumbersRequest(postV1ReturnsIdTrackingnumbersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdTrackingNumbersAPI.PostV1ReturnsIdTrackingnumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ReturnsIdTrackingnumbers`: bool
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdTrackingNumbersAPI.PostV1ReturnsIdTrackingnumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ReturnsIdTrackingnumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ReturnsIdTrackingnumbersRequest** | [**PostV1ReturnsIdTrackingnumbersRequest**](PostV1ReturnsIdTrackingnumbersRequest.md) |  | 

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

