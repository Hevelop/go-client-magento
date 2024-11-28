# \TeamCompanyIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TeamCompanyId**](TeamCompanyIdAPI.md#PostV1TeamCompanyId) | **Post** /V1/team/{companyId} | team/{companyId}



## PostV1TeamCompanyId

> ErrorResponse PostV1TeamCompanyId(ctx, companyId).PostV1TeamCompanyIdRequest(postV1TeamCompanyIdRequest).Execute()

team/{companyId}



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
	companyId := int32(56) // int32 | 
	postV1TeamCompanyIdRequest := *openapiclient.NewPostV1TeamCompanyIdRequest(*openapiclient.NewCompanyDataTeamInterface()) // PostV1TeamCompanyIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamCompanyIdAPI.PostV1TeamCompanyId(context.Background(), companyId).PostV1TeamCompanyIdRequest(postV1TeamCompanyIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamCompanyIdAPI.PostV1TeamCompanyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TeamCompanyId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamCompanyIdAPI.PostV1TeamCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TeamCompanyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1TeamCompanyIdRequest** | [**PostV1TeamCompanyIdRequest**](PostV1TeamCompanyIdRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

