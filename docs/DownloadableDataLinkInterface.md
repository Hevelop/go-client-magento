# # DownloadableDataLinkInterface


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | Sample(or link) id  | [optional]
**Title**| **string** |   | [optional]
**SortOrder**| **int32** |   |
**IsShareable**| **int32** | Shareable status  |
**Price**| **float32** | Price  |
**NumberOfDownloads**| **int32** | Of downloads per user  | [optional]
**LinkType**| **string** |   |
**LinkFile**| **string** | relative file path  | [optional]
**LinkFileContent**| [**DownloadableDataFileContentInterface**](DownloadableDataFileContentInterface.md) |   | [optional]
**LinkUrl**| **string** | Link url or null when type is &#39;file&#39;  | [optional]
**SampleType**| **string** |   |
**SampleFile**| **string** | relative file path  | [optional]
**SampleFileContent**| [**DownloadableDataFileContentInterface**](DownloadableDataFileContentInterface.md) |   | [optional]
**SampleUrl**| **string** | file URL  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Downloadable\\Api\\Data\\LinkInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

