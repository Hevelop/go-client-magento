# \StoreStoreViewsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1StoreStoreViews**](StoreStoreViewsAPI.md#GetV1StoreStoreViews) | **Get** /V1/store/storeViews | store/storeViews



## GetV1StoreStoreViews

> []StoreDataStoreInterface GetV1StoreStoreViews(ctx).Execute()

store/storeViews



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
	resp, r, err := apiClient.StoreStoreViewsAPI.GetV1StoreStoreViews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreStoreViewsAPI.GetV1StoreStoreViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1StoreStoreViews`: []StoreDataStoreInterface
	fmt.Fprintf(os.Stdout, "Response from `StoreStoreViewsAPI.GetV1StoreStoreViews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1StoreStoreViewsRequest struct via the builder pattern


### Return type

[**[]StoreDataStoreInterface**](StoreDataStoreInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

