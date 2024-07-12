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

// checks if the UpdateVendorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVendorRequest{}

// UpdateVendorRequest Update Vendor Request
type UpdateVendorRequest struct {
	// Specify the status of vendor that should be updated. Possible values: ACTIVE,BLOCKED, DELETED
	Status string `json:"status"`
	// Specify the name of the vendor to be updated. Name should not have any special character except . / - &
	Name string `json:"name"`
	// Specify the vendor email ID that should be updated. String in email ID format (Ex:johndoe_1@cashfree.com) should contain @ and dot (.)
	Email string `json:"email"`
	// Specify the beneficiaries phone number to be updated. Phone number registered in India (only digits, 8 - 12 characters after excluding +91).
	Phone string `json:"phone"`
	// Specify if the vendor bank account details should be verified. Possible values: true or false
	VerifyAccount *bool `json:"verify_account,omitempty"`
	// Update if the vendor will have dashboard access or not. Possible values are: true or false
	DashboardAccess *bool `json:"dashboard_access,omitempty"`
	// Specify the settlement cycle to be updated. View the settlement cycle details from the \"Settlement Cycles Supported\" table.  If no schedule option is configured, the settlement cycle ID \"1\" will be in effect. Select \"8\" or \"9\" if you want to schedule instant vendor settlements.
	ScheduleOption float32 `json:"schedule_option"`
	// Specify the vendor bank account details to be updated.
	Bank []BankDetails `json:"bank,omitempty"`
	// Updated beneficiary upi vpa. Alphanumeric, dot (.), hyphen (-), at sign (@), and underscore allowed (100 character limit). Note: underscore and dot (.) gets accepted before and after @, but hyphen (-) is only accepted before @ sign.
	Upi []UpiDetails `json:"upi,omitempty"`
	// Specify the kyc details that should be updated.
	KycDetails []KycDetails `json:"kyc_details"`
}


func (o UpdateVendorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVendorRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["phone"] = o.Phone
	if !IsNil(o.VerifyAccount) {
		toSerialize["verify_account"] = o.VerifyAccount
	}
	if !IsNil(o.DashboardAccess) {
		toSerialize["dashboard_access"] = o.DashboardAccess
	}
	toSerialize["schedule_option"] = o.ScheduleOption
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.Upi) {
		toSerialize["upi"] = o.Upi
	}
	toSerialize["kyc_details"] = o.KycDetails
	return toSerialize, nil
}



