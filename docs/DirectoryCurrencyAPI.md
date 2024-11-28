# \DirectoryCurrencyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1DirectoryCurrency**](DirectoryCurrencyAPI.md#GetV1DirectoryCurrency) | **Get** /V1/directory/currency | directory/currency



## GetV1DirectoryCurrency

> DirectoryDataCurrencyInformationInterface GetV1DirectoryCurrency(ctx).Execute()

directory/currency



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
	resp, r, err := apiClient.DirectoryCurrencyAPI.GetV1DirectoryCurrency(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryCurrencyAPI.GetV1DirectoryCurrency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1DirectoryCurrency`: DirectoryDataCurrencyInformationInterface
	fmt.Fprintf(os.Stdout, "Response from `DirectoryCurrencyAPI.GetV1DirectoryCurrency`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1DirectoryCurrencyRequest struct via the builder pattern


### Return type

[**DirectoryDataCurrencyInformationInterface**](DirectoryDataCurrencyInformationInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

