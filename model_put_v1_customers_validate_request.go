/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
	"fmt"
)

// checks if the PutV1CustomersValidateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1CustomersValidateRequest{}

// PutV1CustomersValidateRequest struct for PutV1CustomersValidateRequest
type PutV1CustomersValidateRequest struct {
	Customer             CustomerDataCustomerInterface `json:"customer"`
	AdditionalProperties map[string]interface{}
}

type _PutV1CustomersValidateRequest PutV1CustomersValidateRequest

// NewPutV1CustomersValidateRequest instantiates a new PutV1CustomersValidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1CustomersValidateRequest(customer CustomerDataCustomerInterface) *PutV1CustomersValidateRequest {
	this := PutV1CustomersValidateRequest{}
	this.Customer = customer
	return &this
}

// NewPutV1CustomersValidateRequestWithDefaults instantiates a new PutV1CustomersValidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1CustomersValidateRequestWithDefaults() *PutV1CustomersValidateRequest {
	this := PutV1CustomersValidateRequest{}
	return &this
}

// GetCustomer returns the Customer field value
func (o *PutV1CustomersValidateRequest) GetCustomer() CustomerDataCustomerInterface {
	if o == nil {
		var ret CustomerDataCustomerInterface
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *PutV1CustomersValidateRequest) GetCustomerOk() (*CustomerDataCustomerInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *PutV1CustomersValidateRequest) SetCustomer(v CustomerDataCustomerInterface) {
	o.Customer = v
}

func (o PutV1CustomersValidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1CustomersValidateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer"] = o.Customer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1CustomersValidateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPutV1CustomersValidateRequest := _PutV1CustomersValidateRequest{}

	err = json.Unmarshal(data, &varPutV1CustomersValidateRequest)

	if err != nil {
		return err
	}

	*o = PutV1CustomersValidateRequest(varPutV1CustomersValidateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1CustomersValidateRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1CustomersValidateRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1CustomersValidateRequest struct {
	value *PutV1CustomersValidateRequest
	isSet bool
}

func (v NullablePutV1CustomersValidateRequest) Get() *PutV1CustomersValidateRequest {
	return v.value
}

func (v *NullablePutV1CustomersValidateRequest) Set(val *PutV1CustomersValidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1CustomersValidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1CustomersValidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1CustomersValidateRequest(val *PutV1CustomersValidateRequest) *NullablePutV1CustomersValidateRequest {
	return &NullablePutV1CustomersValidateRequest{value: val, isSet: true}
}

func (v NullablePutV1CustomersValidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1CustomersValidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
