# \DirectoryCountriesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1DirectoryCountries**](DirectoryCountriesAPI.md#GetV1DirectoryCountries) | **Get** /V1/directory/countries | directory/countries



## GetV1DirectoryCountries

> []DirectoryDataCountryInformationInterface GetV1DirectoryCountries(ctx).Execute()

directory/countries



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectoryCountriesAPI.GetV1DirectoryCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryCountriesAPI.GetV1DirectoryCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1DirectoryCountries`: []DirectoryDataCountryInformationInterface
	fmt.Fprintf(os.Stdout, "Response from `DirectoryCountriesAPI.GetV1DirectoryCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1DirectoryCountriesRequest struct via the builder pattern


### Return type

[**[]DirectoryDataCountryInformationInterface**](DirectoryDataCountryInformationInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

