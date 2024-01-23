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
	"fmt"
)

// PaymentMethodInPaymentsEntityPaymentMethod - struct for PaymentMethodInPaymentsEntityPaymentMethod
type PaymentMethodInPaymentsEntityPaymentMethod struct {
	PaymentMethodAppInPaymentsEntity *PaymentMethodAppInPaymentsEntity
	PaymentMethodCardInPaymentsEntity *PaymentMethodCardInPaymentsEntity
	PaymentMethodCardlessEMIInPaymentsEntity *PaymentMethodCardlessEMIInPaymentsEntity
	PaymentMethodNetBankingInPaymentsEntity *PaymentMethodNetBankingInPaymentsEntity
	PaymentMethodPaylaterInPaymentsEntity *PaymentMethodPaylaterInPaymentsEntity
	PaymentMethodUPIInPaymentsEntity *PaymentMethodUPIInPaymentsEntity
}

// PaymentMethodAppInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod is a convenience function that returns PaymentMethodAppInPaymentsEntity wrapped in PaymentMethodInPaymentsEntityPaymentMethod
func PaymentMethodAppInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod(v *PaymentMethodAppInPaymentsEntity) PaymentMethodInPaymentsEntityPaymentMethod {
	return PaymentMethodInPaymentsEntityPaymentMethod{
		PaymentMethodAppInPaymentsEntity: v,
	}
}

// PaymentMethodCardInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod is a convenience function that returns PaymentMethodCardInPaymentsEntity wrapped in PaymentMethodInPaymentsEntityPaymentMethod
func PaymentMethodCardInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod(v *PaymentMethodCardInPaymentsEntity) PaymentMethodInPaymentsEntityPaymentMethod {
	return PaymentMethodInPaymentsEntityPaymentMethod{
		PaymentMethodCardInPaymentsEntity: v,
	}
}

// PaymentMethodCardlessEMIInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod is a convenience function that returns PaymentMethodCardlessEMIInPaymentsEntity wrapped in PaymentMethodInPaymentsEntityPaymentMethod
func PaymentMethodCardlessEMIInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod(v *PaymentMethodCardlessEMIInPaymentsEntity) PaymentMethodInPaymentsEntityPaymentMethod {
	return PaymentMethodInPaymentsEntityPaymentMethod{
		PaymentMethodCardlessEMIInPaymentsEntity: v,
	}
}

// PaymentMethodNetBankingInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod is a convenience function that returns PaymentMethodNetBankingInPaymentsEntity wrapped in PaymentMethodInPaymentsEntityPaymentMethod
func PaymentMethodNetBankingInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod(v *PaymentMethodNetBankingInPaymentsEntity) PaymentMethodInPaymentsEntityPaymentMethod {
	return PaymentMethodInPaymentsEntityPaymentMethod{
		PaymentMethodNetBankingInPaymentsEntity: v,
	}
}

// PaymentMethodPaylaterInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod is a convenience function that returns PaymentMethodPaylaterInPaymentsEntity wrapped in PaymentMethodInPaymentsEntityPaymentMethod
func PaymentMethodPaylaterInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod(v *PaymentMethodPaylaterInPaymentsEntity) PaymentMethodInPaymentsEntityPaymentMethod {
	return PaymentMethodInPaymentsEntityPaymentMethod{
		PaymentMethodPaylaterInPaymentsEntity: v,
	}
}

// PaymentMethodUPIInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod is a convenience function that returns PaymentMethodUPIInPaymentsEntity wrapped in PaymentMethodInPaymentsEntityPaymentMethod
func PaymentMethodUPIInPaymentsEntityAsPaymentMethodInPaymentsEntityPaymentMethod(v *PaymentMethodUPIInPaymentsEntity) PaymentMethodInPaymentsEntityPaymentMethod {
	return PaymentMethodInPaymentsEntityPaymentMethod{
		PaymentMethodUPIInPaymentsEntity: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentMethodInPaymentsEntityPaymentMethod) UnmarshalJSON(data []byte) error {
	fmt.Println("data")
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentMethodInPaymentsEntityPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodAppInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodAppInPaymentsEntity)
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
func (obj *PaymentMethodInPaymentsEntityPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodAppInPaymentsEntity != nil {
		return obj.PaymentMethodAppInPaymentsEntity
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



