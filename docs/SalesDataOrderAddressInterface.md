# # SalesDataOrderAddressInterface
Order address interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType**| **string** | Address type.  |
**City**| **string** | City.  |
**Company**| **string** | Company.  | [optional]
**CountryId**| **string** | Country ID.  |
**CustomerAddressId**| **int32** | Country address ID.  | [optional]
**CustomerId**| **int32** | Customer ID.  | [optional]
**Email**| **string** | Email address.  | [optional]
**EntityId**| **int32** | Order address ID.  | [optional]
**Fax**| **string** | Fax number.  | [optional]
**Firstname**| **string** | First name.  |
**Lastname**| **string** | Last name.  |
**Middlename**| **string** | Middle name.  | [optional]
**ParentId**| **int32** | Parent ID.  | [optional]
**Postcode**| **string** | Postal code.  |
**Prefix**| **string** | Prefix.  | [optional]
**Region**| **string** | Region.  | [optional]
**RegionCode**| **string** | Region code.  | [optional]
**RegionId**| **int32** | Region ID.  | [optional]
**Street**| **[]string** | Array of any street values. Otherwise, null.  | [optional]
**Suffix**| **string** | Suffix.  | [optional]
**Telephone**| **string** | Telephone number.  |
**VatId**| **string** | VAT ID.  | [optional]
**VatIsValid**| **int32** | VAT-is-valid flag value.  | [optional]
**VatRequestDate**| **string** | VAT request date.  | [optional]
**VatRequestId**| **string** | VAT request ID.  | [optional]
**VatRequestSuccess**| **int32** | VAT-request-success flag value.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\OrderAddressInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

