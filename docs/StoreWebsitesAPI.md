# \StoreWebsitesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1StoreWebsites**](StoreWebsitesAPI.md#GetV1StoreWebsites) | **Get** /V1/store/websites | store/websites



## GetV1StoreWebsites

> []StoreDataWebsiteInterface GetV1StoreWebsites(ctx).Execute()

store/websites



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
	resp, r, err := apiClient.StoreWebsitesAPI.GetV1StoreWebsites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreWebsitesAPI.GetV1StoreWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1StoreWebsites`: []StoreDataWebsiteInterface
	fmt.Fprintf(os.Stdout, "Response from `StoreWebsitesAPI.GetV1StoreWebsites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1StoreWebsitesRequest struct via the builder pattern


### Return type

[**[]StoreDataWebsiteInterface**](StoreDataWebsiteInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

