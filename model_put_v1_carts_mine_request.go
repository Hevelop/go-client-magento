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

// checks if the PutV1CartsMineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1CartsMineRequest{}

// PutV1CartsMineRequest struct for PutV1CartsMineRequest
type PutV1CartsMineRequest struct {
	Quote                QuoteDataCartInterface `json:"quote"`
	AdditionalProperties map[string]interface{}
}

type _PutV1CartsMineRequest PutV1CartsMineRequest

// NewPutV1CartsMineRequest instantiates a new PutV1CartsMineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1CartsMineRequest(quote QuoteDataCartInterface) *PutV1CartsMineRequest {
	this := PutV1CartsMineRequest{}
	this.Quote = quote
	return &this
}

// NewPutV1CartsMineRequestWithDefaults instantiates a new PutV1CartsMineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1CartsMineRequestWithDefaults() *PutV1CartsMineRequest {
	this := PutV1CartsMineRequest{}
	return &this
}

// GetQuote returns the Quote field value
func (o *PutV1CartsMineRequest) GetQuote() QuoteDataCartInterface {
	if o == nil {
		var ret QuoteDataCartInterface
		return ret
	}

	return o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value
// and a boolean to check if the value has been set.
func (o *PutV1CartsMineRequest) GetQuoteOk() (*QuoteDataCartInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quote, true
}

// SetQuote sets field value
func (o *PutV1CartsMineRequest) SetQuote(v QuoteDataCartInterface) {
	o.Quote = v
}

func (o PutV1CartsMineRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1CartsMineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quote"] = o.Quote

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1CartsMineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quote",
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

	varPutV1CartsMineRequest := _PutV1CartsMineRequest{}

	err = json.Unmarshal(data, &varPutV1CartsMineRequest)

	if err != nil {
		return err
	}

	*o = PutV1CartsMineRequest(varPutV1CartsMineRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "quote")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1CartsMineRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1CartsMineRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1CartsMineRequest struct {
	value *PutV1CartsMineRequest
	isSet bool
}

func (v NullablePutV1CartsMineRequest) Get() *PutV1CartsMineRequest {
	return v.value
}

func (v *NullablePutV1CartsMineRequest) Set(val *PutV1CartsMineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1CartsMineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1CartsMineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1CartsMineRequest(val *PutV1CartsMineRequest) *NullablePutV1CartsMineRequest {
	return &NullablePutV1CartsMineRequest{value: val, isSet: true}
}

func (v NullablePutV1CartsMineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1CartsMineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
