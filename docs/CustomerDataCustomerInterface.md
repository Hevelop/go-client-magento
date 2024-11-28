# # CustomerDataCustomerInterface
Customer entity interface for API handling.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | Customer id  | [optional]
**GroupId**| **int32** | Group id  | [optional]
**DefaultBilling**| **string** | Default billing address id  | [optional]
**DefaultShipping**| **string** | Default shipping address id  | [optional]
**Confirmation**| **string** | Confirmation  | [optional]
**CreatedAt**| **string** | Created at time  | [optional]
**UpdatedAt**| **string** | Updated at time  | [optional]
**CreatedIn**| **string** | Created in area  | [optional]
**Dob**| **string** | In keeping with current security and privacy best practices, be sure you are aware of any potential legal and security risks associated with the storage of customers’ full date of birth (month, day, year) along with other personal identifiers (e.g., full name) before collecting or processing such data.  | [optional]
**Email**| **string** | Email address  |
**Firstname**| **string** | First name  |
**Lastname**| **string** | Last name  |
**Middlename**| **string** | Middle name  | [optional]
**Prefix**| **string** | Prefix  | [optional]
**Suffix**| **string** | Suffix  | [optional]
**Gender**| **int32** | Gender  | [optional]
**StoreId**| **int32** | Store id  | [optional]
**Taxvat**| **string** | Tax Vat  | [optional]
**WebsiteId**| **int32** | Website id  | [optional]
**Addresses**| [**[]CustomerDataAddressInterface**](CustomerDataAddressInterface.md) | Customer addresses.  | [optional]
**DisableAutoGroupChange**| **int32** | Disable auto group change flag.  | [optional]
**ExtensionAttributes**| [**CustomerDataCustomerExtensionInterface**](CustomerDataCustomerExtensionInterface.md) |   | [optional]
**CustomAttributes**| [**[]FrameworkAttributeInterface**](FrameworkAttributeInterface.md) | Custom attributes values.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

