# # CheckoutAgreementsDataAgreementInterface
Interface AgreementInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementId**| **int32** | Agreement ID.  |
**Name**| **string** | Agreement name.  |
**Content**| **string** | Agreement content.  |
**ContentHeight**| **string** | Agreement content height. Otherwise, null.  | [optional]
**CheckboxText**| **string** | Agreement checkbox text.  |
**IsActive**| **bool** | Agreement status.  |
**IsHtml**| **bool** | * true - HTML. * false - plain text.  |
**Mode**| **int32** | The agreement applied mode.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\CheckoutAgreements\\Api\\Data\\AgreementInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

