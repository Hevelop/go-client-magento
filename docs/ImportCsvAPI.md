# \ImportCsvAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ImportCsv**](ImportCsvAPI.md#PostV1ImportCsv) | **Post** /V1/import/csv | import/csv



## PostV1ImportCsv

> []string PostV1ImportCsv(ctx).PostV1ImportCsvRequest(postV1ImportCsvRequest).Execute()

import/csv



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
	postV1ImportCsvRequest := *openapiclient.NewPostV1ImportCsvRequest(*openapiclient.NewImportCsvApiDataLocalizedSourceDataInterface("Entity_example", "Behavior_example", "ValidationStrategy_example", "AllowedErrorCount_example", "CsvData_example")) // PostV1ImportCsvRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportCsvAPI.PostV1ImportCsv(context.Background()).PostV1ImportCsvRequest(postV1ImportCsvRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportCsvAPI.PostV1ImportCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ImportCsv`: []string
	fmt.Fprintf(os.Stdout, "Response from `ImportCsvAPI.PostV1ImportCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ImportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ImportCsvRequest** | [**PostV1ImportCsvRequest**](PostV1ImportCsvRequest.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

