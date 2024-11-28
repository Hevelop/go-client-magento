# \ImportJsonAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ImportJson**](ImportJsonAPI.md#PostV1ImportJson) | **Post** /V1/import/json | import/json



## PostV1ImportJson

> []string PostV1ImportJson(ctx).PostV1ImportJsonRequest(postV1ImportJsonRequest).Execute()

import/json



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
	postV1ImportJsonRequest := *openapiclient.NewPostV1ImportJsonRequest(*openapiclient.NewImportJsonApiDataSourceDataInterface("Entity_example", "Behavior_example", "ValidationStrategy_example", "AllowedErrorCount_example", map[string]interface{}(123))) // PostV1ImportJsonRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportJsonAPI.PostV1ImportJson(context.Background()).PostV1ImportJsonRequest(postV1ImportJsonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportJsonAPI.PostV1ImportJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ImportJson`: []string
	fmt.Fprintf(os.Stdout, "Response from `ImportJsonAPI.PostV1ImportJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ImportJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ImportJsonRequest** | [**PostV1ImportJsonRequest**](PostV1ImportJsonRequest.md) |  | 

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

