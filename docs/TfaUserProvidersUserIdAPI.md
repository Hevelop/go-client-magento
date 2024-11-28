# \TfaUserProvidersUserIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1TfaUserprovidersUserId**](TfaUserProvidersUserIdAPI.md#GetV1TfaUserprovidersUserId) | **Get** /V1/tfa/user-providers/{userId} | tfa/user-providers/{userId}
[**PutV1TfaUserprovidersUserId**](TfaUserProvidersUserIdAPI.md#PutV1TfaUserprovidersUserId) | **Put** /V1/tfa/user-providers/{userId} | tfa/user-providers/{userId}



## GetV1TfaUserprovidersUserId

> []TwoFactorAuthProviderInterface GetV1TfaUserprovidersUserId(ctx, userId).Execute()

tfa/user-providers/{userId}



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
	resp, r, err := apiClient.TfaUserProvidersUserIdAPI.GetV1TfaUserprovidersUserId(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaUserProvidersUserIdAPI.GetV1TfaUserprovidersUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TfaUserprovidersUserId`: []TwoFactorAuthProviderInterface
	fmt.Fprintf(os.Stdout, "Response from `TfaUserProvidersUserIdAPI.GetV1TfaUserprovidersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TfaUserprovidersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TwoFactorAuthProviderInterface**](TwoFactorAuthProviderInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1TfaUserprovidersUserId

> bool PutV1TfaUserprovidersUserId(ctx, userId).PutV1TfaUserprovidersUserIdRequest(putV1TfaUserprovidersUserIdRequest).Execute()

tfa/user-providers/{userId}



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
	putV1TfaUserprovidersUserIdRequest := *openapiclient.NewPutV1TfaUserprovidersUserIdRequest("ProvidersCodes_example") // PutV1TfaUserprovidersUserIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaUserProvidersUserIdAPI.PutV1TfaUserprovidersUserId(context.Background(), userId).PutV1TfaUserprovidersUserIdRequest(putV1TfaUserprovidersUserIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaUserProvidersUserIdAPI.PutV1TfaUserprovidersUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1TfaUserprovidersUserId`: bool
	fmt.Fprintf(os.Stdout, "Response from `TfaUserProvidersUserIdAPI.PutV1TfaUserprovidersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1TfaUserprovidersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1TfaUserprovidersUserIdRequest** | [**PutV1TfaUserprovidersUserIdRequest**](PutV1TfaUserprovidersUserIdRequest.md) |  | 

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

