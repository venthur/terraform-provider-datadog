/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// FormulaAndFunctionsAggregation The aggregation method.
type FormulaAndFunctionsAggregation string

// List of FormulaAndFunctionsAggregation
const (
	FORMULAANDFUNCTIONSAGGREGATION_AVG  FormulaAndFunctionsAggregation = "avg"
	FORMULAANDFUNCTIONSAGGREGATION_MIN  FormulaAndFunctionsAggregation = "min"
	FORMULAANDFUNCTIONSAGGREGATION_MAX  FormulaAndFunctionsAggregation = "max"
	FORMULAANDFUNCTIONSAGGREGATION_SUM  FormulaAndFunctionsAggregation = "sum"
	FORMULAANDFUNCTIONSAGGREGATION_LAST FormulaAndFunctionsAggregation = "last"
)

func (v *FormulaAndFunctionsAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FormulaAndFunctionsAggregation(value)
	for _, existing := range []FormulaAndFunctionsAggregation{"avg", "min", "max", "sum", "last"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FormulaAndFunctionsAggregation", value)
}

// Ptr returns reference to FormulaAndFunctionsAggregation value
func (v FormulaAndFunctionsAggregation) Ptr() *FormulaAndFunctionsAggregation {
	return &v
}

type NullableFormulaAndFunctionsAggregation struct {
	value *FormulaAndFunctionsAggregation
	isSet bool
}

func (v NullableFormulaAndFunctionsAggregation) Get() *FormulaAndFunctionsAggregation {
	return v.value
}

func (v *NullableFormulaAndFunctionsAggregation) Set(val *FormulaAndFunctionsAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionsAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionsAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionsAggregation(val *FormulaAndFunctionsAggregation) *NullableFormulaAndFunctionsAggregation {
	return &NullableFormulaAndFunctionsAggregation{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionsAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionsAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}