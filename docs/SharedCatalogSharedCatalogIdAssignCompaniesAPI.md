# \SharedCatalogSharedCatalogIdAssignCompaniesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SharedCatalogSharedCatalogIdAssignCompanies**](SharedCatalogSharedCatalogIdAssignCompaniesAPI.md#PostV1SharedCatalogSharedCatalogIdAssignCompanies) | **Post** /V1/sharedCatalog/{sharedCatalogId}/assignCompanies | sharedCatalog/{sharedCatalogId}/assignCompanies



## PostV1SharedCatalogSharedCatalogIdAssignCompanies

> bool PostV1SharedCatalogSharedCatalogIdAssignCompanies(ctx, sharedCatalogId).PostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest(postV1SharedCatalogSharedCatalogIdAssignCompaniesRequest).Execute()

sharedCatalog/{sharedCatalogId}/assignCompanies



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
	sharedCatalogId := int32(56) // int32 | 
	postV1SharedCatalogSharedCatalogIdAssignCompaniesRequest := *openapiclient.NewPostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest([]openapiclient.CompanyDataCompanyInterface{*openapiclient.NewCompanyDataCompanyInterface([]string{"Street_example"}, int32(123), int32(123), "RejectReason_example", "RejectedAt_example", int32(123))}) // PostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogSharedCatalogIdAssignCompaniesAPI.PostV1SharedCatalogSharedCatalogIdAssignCompanies(context.Background(), sharedCatalogId).PostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest(postV1SharedCatalogSharedCatalogIdAssignCompaniesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogSharedCatalogIdAssignCompaniesAPI.PostV1SharedCatalogSharedCatalogIdAssignCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SharedCatalogSharedCatalogIdAssignCompanies`: bool
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogSharedCatalogIdAssignCompaniesAPI.PostV1SharedCatalogSharedCatalogIdAssignCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedCatalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1SharedCatalogSharedCatalogIdAssignCompaniesRequest** | [**PostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest**](PostV1SharedCatalogSharedCatalogIdAssignCompaniesRequest.md) |  | 

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

