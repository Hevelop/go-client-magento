# # AsynchronousOperationsDataBulkOperationsStatusInterface
Interface BulkStatusInterface Bulk summary data with list of operations items summary data.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationsList**| [**[]AsynchronousOperationsDataSummaryOperationStatusInterface**](AsynchronousOperationsDataSummaryOperationStatusInterface.md) | List of operation with statuses (short data).  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\AsynchronousOperations\\Api\\Data\\BulkSummaryInterface  | [optional]
**UserType**| **int32** | User type  |
**BulkId**| **string** | Bulk uuid  |
**Description**| **string** | Bulk description  |
**StartTime**| **string** | Bulk scheduled time  |
**UserId**| **int32** | User id  |
**OperationCount**| **int32** | Total number of operations scheduled in scope of this bulk  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

