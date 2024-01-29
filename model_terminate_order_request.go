/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the TerminateOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateOrderRequest{}

// TerminateOrderRequest Request to terminate an active order at Cashfree
type TerminateOrderRequest struct {
	// To terminate an order, pass order_status as \"TERMINATE\". Please note, order might not be terminated - confirm with the order_status in response. \"TERMINATION_REQUESTED\" states that the request is recieved and we are working on it. If the order terminates successfully, status will change to \"TERMINATED\". Incase there's any active transaction which moved to success - order might not get terminated.
	OrderStatus string `json:"order_status"`
}


func (o TerminateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_status"] = o.OrderStatus
	return toSerialize, nil
}



