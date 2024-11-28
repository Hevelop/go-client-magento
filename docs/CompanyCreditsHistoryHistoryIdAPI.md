# \CompanyCreditsHistoryHistoryIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CompanyCreditsHistoryHistoryId**](CompanyCreditsHistoryHistoryIdAPI.md#PutV1CompanyCreditsHistoryHistoryId) | **Put** /V1/companyCredits/history/{historyId} | companyCredits/history/{historyId}



## PutV1CompanyCreditsHistoryHistoryId

> bool PutV1CompanyCreditsHistoryHistoryId(ctx, historyId).PutV1CompanyCreditsHistoryHistoryIdRequest(putV1CompanyCreditsHistoryHistoryIdRequest).Execute()

companyCredits/history/{historyId}



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
	historyId := int32(56) // int32 | 
	putV1CompanyCreditsHistoryHistoryIdRequest := *openapiclient.NewPutV1CompanyCreditsHistoryHistoryIdRequest() // PutV1CompanyCreditsHistoryHistoryIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCreditsHistoryHistoryIdAPI.PutV1CompanyCreditsHistoryHistoryId(context.Background(), historyId).PutV1CompanyCreditsHistoryHistoryIdRequest(putV1CompanyCreditsHistoryHistoryIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCreditsHistoryHistoryIdAPI.PutV1CompanyCreditsHistoryHistoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CompanyCreditsHistoryHistoryId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CompanyCreditsHistoryHistoryIdAPI.PutV1CompanyCreditsHistoryHistoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**historyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CompanyCreditsHistoryHistoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CompanyCreditsHistoryHistoryIdRequest** | [**PutV1CompanyCreditsHistoryHistoryIdRequest**](PutV1CompanyCreditsHistoryHistoryIdRequest.md) |  | 

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

