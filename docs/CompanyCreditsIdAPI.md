# \CompanyCreditsIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CompanyCreditsId**](CompanyCreditsIdAPI.md#PutV1CompanyCreditsId) | **Put** /V1/companyCredits/{id} | companyCredits/{id}



## PutV1CompanyCreditsId

> CompanyCreditDataCreditLimitInterface PutV1CompanyCreditsId(ctx, id).PutV1CompanyCreditsIdRequest(putV1CompanyCreditsIdRequest).Execute()

companyCredits/{id}



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
	putV1CompanyCreditsIdRequest := *openapiclient.NewPutV1CompanyCreditsIdRequest(*openapiclient.NewCompanyCreditDataCreditLimitInterface(false)) // PutV1CompanyCreditsIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCreditsIdAPI.PutV1CompanyCreditsId(context.Background(), id).PutV1CompanyCreditsIdRequest(putV1CompanyCreditsIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCreditsIdAPI.PutV1CompanyCreditsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CompanyCreditsId`: CompanyCreditDataCreditLimitInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyCreditsIdAPI.PutV1CompanyCreditsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CompanyCreditsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CompanyCreditsIdRequest** | [**PutV1CompanyCreditsIdRequest**](PutV1CompanyCreditsIdRequest.md) |  | 

### Return type

[**CompanyCreditDataCreditLimitInterface**](CompanyCreditDataCreditLimitInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

