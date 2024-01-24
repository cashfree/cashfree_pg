/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
)

// checks if the CardOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardOffer{}

// CardOffer struct for CardOffer
type CardOffer struct {
	Type []string `json:"type"`
	// Bank Name of Card.
	BankName string `json:"bank_name"`
	SchemeName []string `json:"scheme_name"`
}


func (o CardOffer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["bank_name"] = o.BankName
	toSerialize["scheme_name"] = o.SchemeName
	return toSerialize, nil
}



