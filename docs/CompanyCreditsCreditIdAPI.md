# \CompanyCreditsCreditIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CompanyCreditsCreditId**](CompanyCreditsCreditIdAPI.md#GetV1CompanyCreditsCreditId) | **Get** /V1/companyCredits/{creditId} | companyCredits/{creditId}



## GetV1CompanyCreditsCreditId

> CompanyCreditDataCreditLimitInterface GetV1CompanyCreditsCreditId(ctx, creditId).Reload(reload).Execute()

companyCredits/{creditId}



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
	creditId := int32(56) // int32 | 
	reload := true // bool | [optional] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCreditsCreditIdAPI.GetV1CompanyCreditsCreditId(context.Background(), creditId).Reload(reload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCreditsCreditIdAPI.GetV1CompanyCreditsCreditId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CompanyCreditsCreditId`: CompanyCreditDataCreditLimitInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyCreditsCreditIdAPI.GetV1CompanyCreditsCreditId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CompanyCreditsCreditIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reload** | **bool** | [optional] | 

### Return type

[**CompanyCreditDataCreditLimitInterface**](CompanyCreditDataCreditLimitInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

