# # SalesDataTransactionInterface
Transaction interface. A transaction is an interaction between a merchant and a customer such as a purchase, a credit, a refund, and so on.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId**| **int32** | Transaction ID.  |
**ParentId**| **int32** | The parent ID for the transaction. Otherwise, null.  | [optional]
**OrderId**| **int32** | Order ID.  |
**PaymentId**| **int32** | Payment ID.  |
**TxnId**| **string** | Transaction business ID.  |
**ParentTxnId**| **string** | Parent transaction business ID.  |
**TxnType**| **string** | Transaction type.  |
**IsClosed**| **int32** | Is-closed flag value.  |
**AdditionalInformation**| **[]string** | Array of additional information. Otherwise, null.  | [optional]
**CreatedAt**| **string** | Created-at timestamp.  |
**ChildTransactions**| [**[]SalesDataTransactionInterface**](SalesDataTransactionInterface.md) | Array of child transactions.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\TransactionInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

