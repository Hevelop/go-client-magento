# # CompanyCreditDataHistoryDataInterface
History data transfer object interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | ID.  | [optional]
**CompanyCreditId**| **int32** | Company credit id.  | [optional]
**UserId**| **int32** | User Id.  | [optional]
**UserType**| **int32** | User type: integration, admin, customer.  | [optional]
**CurrencyCredit**| **string** | Currency code of credit.  | [optional]
**CurrencyOperation**| **string** | Currency code of operation.  | [optional]
**Rate**| **float32** | Currency rate between credit and operation currencies.  |
**RateCredit**| **float32** | Rate between credit and base currencies.  | [optional]
**Amount**| **float32** | Amount.  |
**Balance**| **float32** | Outstanding balance.  |
**CreditLimit**| **float32** | Credit limit.  |
**AvailableLimit**| **float32** | Available limit.  | [optional]
**Type**| **int32** | Type of operation.  | [optional]
**Datetime**| **string** | Operation datetime.  | [optional]
**PurchaseOrder**| **string** | Purchase Order number.  | [optional]
**CustomReferenceNumber**| **string** | Custom Reference number.  | [optional]
**Comment**| **string** | Comment.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

