# \TfaDefaultProviderCodeUserIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaDefaultprovidercodeUserId**](TfaDefaultProviderCodeUserIdAPI.md#GetV1TfaDefaultprovidercodeUserId) | **Get** /V1/tfa/default-provider-code/{userId} | tfa/default-provider-code/{userId}
[**PutV1TfaDefaultprovidercodeUserId**](TfaDefaultProviderCodeUserIdAPI.md#PutV1TfaDefaultprovidercodeUserId) | **Put** /V1/tfa/default-provider-code/{userId} | tfa/default-provider-code/{userId}



## GetV1TfaDefaultprovidercodeUserId

> string GetV1TfaDefaultprovidercodeUserId(ctx, userId).Execute()

tfa/default-provider-code/{userId}



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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaDefaultProviderCodeUserIdAPI.GetV1TfaDefaultprovidercodeUserId(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaDefaultProviderCodeUserIdAPI.GetV1TfaDefaultprovidercodeUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaDefaultprovidercodeUserId`: string
	fmt.Fprintf(os.Stdout, "Response from `TfaDefaultProviderCodeUserIdAPI.GetV1TfaDefaultprovidercodeUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaDefaultprovidercodeUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1TfaDefaultprovidercodeUserId

> bool PutV1TfaDefaultprovidercodeUserId(ctx, userId).PutV1TfaDefaultprovidercodeUserIdRequest(putV1TfaDefaultprovidercodeUserIdRequest).Execute()

tfa/default-provider-code/{userId}



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
	userId := int32(56) // int32 | 
	putV1TfaDefaultprovidercodeUserIdRequest := *openapiclient.NewPutV1TfaDefaultprovidercodeUserIdRequest("ProviderCode_example") // PutV1TfaDefaultprovidercodeUserIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaDefaultProviderCodeUserIdAPI.PutV1TfaDefaultprovidercodeUserId(context.Background(), userId).PutV1TfaDefaultprovidercodeUserIdRequest(putV1TfaDefaultprovidercodeUserIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaDefaultProviderCodeUserIdAPI.PutV1TfaDefaultprovidercodeUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1TfaDefaultprovidercodeUserId`: bool
	fmt.Fprintf(os.Stdout, "Response from `TfaDefaultProviderCodeUserIdAPI.PutV1TfaDefaultprovidercodeUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1TfaDefaultprovidercodeUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1TfaDefaultprovidercodeUserIdRequest** | [**PutV1TfaDefaultprovidercodeUserIdRequest**](PutV1TfaDefaultprovidercodeUserIdRequest.md) |  | 

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

