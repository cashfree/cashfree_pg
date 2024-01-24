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
	"fmt"
)

// PaymentEntityPaymentMethod - struct for PaymentEntityPaymentMethod
type PaymentEntityPaymentMethod struct {
	PaymentMethodAppInPaymentsEntity *PaymentMethodAppInPaymentsEntity
	PaymentMethodCardEMIInPaymentsEntity *PaymentMethodCardEMIInPaymentsEntity
	PaymentMethodCardInPaymentsEntity *PaymentMethodCardInPaymentsEntity
	PaymentMethodCardlessEMIInPaymentsEntity *PaymentMethodCardlessEMIInPaymentsEntity
	PaymentMethodNetBankingInPaymentsEntity *PaymentMethodNetBankingInPaymentsEntity
	PaymentMethodPaylaterInPaymentsEntity *PaymentMethodPaylaterInPaymentsEntity
	PaymentMethodUPIInPaymentsEntity *PaymentMethodUPIInPaymentsEntity
}

// PaymentMethodAppInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodAppInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodAppInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodAppInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodAppInPaymentsEntity: v,
	}
}

// PaymentMethodCardEMIInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodCardEMIInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodCardEMIInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodCardEMIInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodCardEMIInPaymentsEntity: v,
	}
}

// PaymentMethodCardInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodCardInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodCardInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodCardInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodCardInPaymentsEntity: v,
	}
}

// PaymentMethodCardlessEMIInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodCardlessEMIInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodCardlessEMIInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodCardlessEMIInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodCardlessEMIInPaymentsEntity: v,
	}
}

// PaymentMethodNetBankingInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodNetBankingInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodNetBankingInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodNetBankingInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodNetBankingInPaymentsEntity: v,
	}
}

// PaymentMethodPaylaterInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodPaylaterInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodPaylaterInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodPaylaterInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodPaylaterInPaymentsEntity: v,
	}
}

// PaymentMethodUPIInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodUPIInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodUPIInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodUPIInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodUPIInPaymentsEntity: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentEntityPaymentMethod) UnmarshalJSON(data []byte) error {
	fmt.Println("data")
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentEntityPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodAppInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodAppInPaymentsEntity)
	}

	if src.PaymentMethodCardEMIInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodCardEMIInPaymentsEntity)
	}

	if src.PaymentMethodCardInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodCardInPaymentsEntity)
	}

	if src.PaymentMethodCardlessEMIInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodCardlessEMIInPaymentsEntity)
	}

	if src.PaymentMethodNetBankingInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodNetBankingInPaymentsEntity)
	}

	if src.PaymentMethodPaylaterInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodPaylaterInPaymentsEntity)
	}

	if src.PaymentMethodUPIInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodUPIInPaymentsEntity)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaymentEntityPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodAppInPaymentsEntity != nil {
		return obj.PaymentMethodAppInPaymentsEntity
	}

	if obj.PaymentMethodCardEMIInPaymentsEntity != nil {
		return obj.PaymentMethodCardEMIInPaymentsEntity
	}

	if obj.PaymentMethodCardInPaymentsEntity != nil {
		return obj.PaymentMethodCardInPaymentsEntity
	}

	if obj.PaymentMethodCardlessEMIInPaymentsEntity != nil {
		return obj.PaymentMethodCardlessEMIInPaymentsEntity
	}

	if obj.PaymentMethodNetBankingInPaymentsEntity != nil {
		return obj.PaymentMethodNetBankingInPaymentsEntity
	}

	if obj.PaymentMethodPaylaterInPaymentsEntity != nil {
		return obj.PaymentMethodPaylaterInPaymentsEntity
	}

	if obj.PaymentMethodUPIInPaymentsEntity != nil {
		return obj.PaymentMethodUPIInPaymentsEntity
	}

	// all schemas are nil
	return nil
}


