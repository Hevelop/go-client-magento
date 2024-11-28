# \DirectoryCountriesCountryIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1DirectoryCountriesCountryId**](DirectoryCountriesCountryIdAPI.md#GetV1DirectoryCountriesCountryId) | **Get** /V1/directory/countries/{countryId} | directory/countries/{countryId}



## GetV1DirectoryCountriesCountryId

> DirectoryDataCountryInformationInterface GetV1DirectoryCountriesCountryId(ctx, countryId).Execute()

directory/countries/{countryId}



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
	countryId := "countryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectoryCountriesCountryIdAPI.GetV1DirectoryCountriesCountryId(context.Background(), countryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryCountriesCountryIdAPI.GetV1DirectoryCountriesCountryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1DirectoryCountriesCountryId`: DirectoryDataCountryInformationInterface
	fmt.Fprintf(os.Stdout, "Response from `DirectoryCountriesCountryIdAPI.GetV1DirectoryCountriesCountryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1DirectoryCountriesCountryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectoryDataCountryInformationInterface**](DirectoryDataCountryInformationInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

