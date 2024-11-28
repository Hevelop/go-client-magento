# # SalesDataOrderPaymentInterface
Order payment interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountStatus**| **string** | Account status.  |
**AdditionalData**| **string** | Additional data.  | [optional]
**AdditionalInformation**| **[]string** | Array of additional information.  |
**AddressStatus**| **string** | Address status.  | [optional]
**AmountAuthorized**| **float32** | Amount authorized.  | [optional]
**AmountCanceled**| **float32** | Amount canceled.  | [optional]
**AmountOrdered**| **float32** | Amount ordered.  | [optional]
**AmountPaid**| **float32** | Amount paid.  | [optional]
**AmountRefunded**| **float32** | Amount refunded.  | [optional]
**AnetTransMethod**| **string** | Anet transaction method.  | [optional]
**BaseAmountAuthorized**| **float32** | Base amount authorized.  | [optional]
**BaseAmountCanceled**| **float32** | Base amount canceled.  | [optional]
**BaseAmountOrdered**| **float32** | Base amount ordered.  | [optional]
**BaseAmountPaid**| **float32** | Base amount paid.  | [optional]
**BaseAmountPaidOnline**| **float32** | Base amount paid online.  | [optional]
**BaseAmountRefunded**| **float32** | Base amount refunded.  | [optional]
**BaseAmountRefundedOnline**| **float32** | Base amount refunded online.  | [optional]
**BaseShippingAmount**| **float32** | Base shipping amount.  | [optional]
**BaseShippingCaptured**| **float32** | Base shipping captured amount.  | [optional]
**BaseShippingRefunded**| **float32** | Base shipping refunded amount.  | [optional]
**CcApproval**| **string** | Credit card approval.  | [optional]
**CcAvsStatus**| **string** | Credit card avs status.  | [optional]
**CcCidStatus**| **string** | Credit card CID status.  | [optional]
**CcDebugRequestBody**| **string** | Credit card debug request body.  | [optional]
**CcDebugResponseBody**| **string** | Credit card debug response body.  | [optional]
**CcDebugResponseSerialized**| **string** | Credit card debug response serialized.  | [optional]
**CcExpMonth**| **string** | Credit card expiration month.  | [optional]
**CcExpYear**| **string** | Credit card expiration year.  | [optional]
**CcLast4**| **string** | Last four digits of the credit card.  |
**CcNumberEnc**| **string** | Encrypted credit card number.  | [optional]
**CcOwner**| **string** | Credit card number.  | [optional]
**CcSecureVerify**| **string** | Credit card secure verify.  | [optional]
**CcSsIssue**| **string** | Credit card SS issue.  | [optional]
**CcSsStartMonth**| **string** | Credit card SS start month.  | [optional]
**CcSsStartYear**| **string** | Credit card SS start year.  | [optional]
**CcStatus**| **string** | Credit card status.  | [optional]
**CcStatusDescription**| **string** | Credit card status description.  | [optional]
**CcTransId**| **string** | Credit card transaction ID.  | [optional]
**CcType**| **string** | Credit card type.  | [optional]
**EcheckAccountName**| **string** | eCheck account name.  | [optional]
**EcheckAccountType**| **string** | eCheck account type.  | [optional]
**EcheckBankName**| **string** | eCheck bank name.  | [optional]
**EcheckRoutingNumber**| **string** | eCheck routing number.  | [optional]
**EcheckType**| **string** | eCheck type.  | [optional]
**EntityId**| **int32** | Entity ID.  | [optional]
**LastTransId**| **string** | Last transaction ID.  | [optional]
**Method**| **string** | Method.  |
**ParentId**| **int32** | Parent ID.  | [optional]
**PoNumber**| **string** | PO number.  | [optional]
**ProtectionEligibility**| **string** | Protection eligibility.  | [optional]
**QuotePaymentId**| **int32** | Quote payment ID.  | [optional]
**ShippingAmount**| **float32** | Shipping amount.  | [optional]
**ShippingCaptured**| **float32** | Shipping captured.  | [optional]
**ShippingRefunded**| **float32** | Shipping refunded.  | [optional]
**ExtensionAttributes**| [**SalesDataOrderPaymentExtensionInterface**](SalesDataOrderPaymentExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

