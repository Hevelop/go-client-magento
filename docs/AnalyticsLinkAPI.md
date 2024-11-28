# \AnalyticsLinkAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1AnalyticsLink**](AnalyticsLinkAPI.md#GetV1AnalyticsLink) | **Get** /V1/analytics/link | analytics/link



## GetV1AnalyticsLink

> AnalyticsDataLinkInterface GetV1AnalyticsLink(ctx).Execute()

analytics/link



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
	resp, r, err := apiClient.AnalyticsLinkAPI.GetV1AnalyticsLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsLinkAPI.GetV1AnalyticsLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AnalyticsLink`: AnalyticsDataLinkInterface
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsLinkAPI.GetV1AnalyticsLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AnalyticsLinkRequest struct via the builder pattern


### Return type

[**AnalyticsDataLinkInterface**](AnalyticsDataLinkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

