# \SharedCatalogSharedCatalogIdCompaniesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1SharedCatalogSharedCatalogIdCompanies**](SharedCatalogSharedCatalogIdCompaniesAPI.md#GetV1SharedCatalogSharedCatalogIdCompanies) | **Get** /V1/sharedCatalog/{sharedCatalogId}/companies | sharedCatalog/{sharedCatalogId}/companies



## GetV1SharedCatalogSharedCatalogIdCompanies

> string GetV1SharedCatalogSharedCatalogIdCompanies(ctx, sharedCatalogId).Execute()

sharedCatalog/{sharedCatalogId}/companies



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogSharedCatalogIdCompaniesAPI.GetV1SharedCatalogSharedCatalogIdCompanies(context.Background(), sharedCatalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogSharedCatalogIdCompaniesAPI.GetV1SharedCatalogSharedCatalogIdCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1SharedCatalogSharedCatalogIdCompanies`: string
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogSharedCatalogIdCompaniesAPI.GetV1SharedCatalogSharedCatalogIdCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedCatalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1SharedCatalogSharedCatalogIdCompaniesRequest struct via the builder pattern


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

