# \TeamTeamIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1TeamTeamId**](TeamTeamIdAPI.md#DeleteV1TeamTeamId) | **Delete** /V1/team/{teamId} | team/{teamId}
[**GetV1TeamTeamId**](TeamTeamIdAPI.md#GetV1TeamTeamId) | **Get** /V1/team/{teamId} | team/{teamId}
[**PutV1TeamTeamId**](TeamTeamIdAPI.md#PutV1TeamTeamId) | **Put** /V1/team/{teamId} | team/{teamId}



## DeleteV1TeamTeamId

> ErrorResponse DeleteV1TeamTeamId(ctx, teamId).Execute()

team/{teamId}



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
	teamId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamTeamIdAPI.DeleteV1TeamTeamId(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamTeamIdAPI.DeleteV1TeamTeamId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1TeamTeamId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamTeamIdAPI.DeleteV1TeamTeamId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1TeamTeamIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1TeamTeamId

> CompanyDataTeamInterface GetV1TeamTeamId(ctx, teamId).Execute()

team/{teamId}



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
	teamId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamTeamIdAPI.GetV1TeamTeamId(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamTeamIdAPI.GetV1TeamTeamId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TeamTeamId`: CompanyDataTeamInterface
	fmt.Fprintf(os.Stdout, "Response from `TeamTeamIdAPI.GetV1TeamTeamId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TeamTeamIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyDataTeamInterface**](CompanyDataTeamInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1TeamTeamId

> bool PutV1TeamTeamId(ctx, teamId).PostV1TeamCompanyIdRequest(postV1TeamCompanyIdRequest).Execute()

team/{teamId}



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
	teamId := "teamId_example" // string | 
	postV1TeamCompanyIdRequest := *openapiclient.NewPostV1TeamCompanyIdRequest(*openapiclient.NewCompanyDataTeamInterface()) // PostV1TeamCompanyIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamTeamIdAPI.PutV1TeamTeamId(context.Background(), teamId).PostV1TeamCompanyIdRequest(postV1TeamCompanyIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamTeamIdAPI.PutV1TeamTeamId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1TeamTeamId`: bool
	fmt.Fprintf(os.Stdout, "Response from `TeamTeamIdAPI.PutV1TeamTeamId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1TeamTeamIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1TeamCompanyIdRequest** | [**PostV1TeamCompanyIdRequest**](PostV1TeamCompanyIdRequest.md) |  | 

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

