# \CompanyCreditsCompanyCompanyIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CompanyCreditsCompanyCompanyId**](CompanyCreditsCompanyCompanyIdAPI.md#GetV1CompanyCreditsCompanyCompanyId) | **Get** /V1/companyCredits/company/{companyId} | companyCredits/company/{companyId}



## GetV1CompanyCreditsCompanyCompanyId

> CompanyCreditDataCreditLimitInterface GetV1CompanyCreditsCompanyCompanyId(ctx, companyId).Execute()

companyCredits/company/{companyId}



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCreditsCompanyCompanyIdAPI.GetV1CompanyCreditsCompanyCompanyId(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCreditsCompanyCompanyIdAPI.GetV1CompanyCreditsCompanyCompanyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CompanyCreditsCompanyCompanyId`: CompanyCreditDataCreditLimitInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyCreditsCompanyCompanyIdAPI.GetV1CompanyCreditsCompanyCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CompanyCreditsCompanyCompanyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

