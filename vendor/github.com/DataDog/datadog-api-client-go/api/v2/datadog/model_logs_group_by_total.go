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

// LogsGroupByTotal - A resulting object to put the given computes in over all the matching records.
type LogsGroupByTotal struct {
	bool    *bool
	float64 *float64
	string  *string
}

// boolAsLogsGroupByTotal is a convenience function that returns bool wrapped in LogsGroupByTotal
func boolAsLogsGroupByTotal(v *bool) LogsGroupByTotal {
	return LogsGroupByTotal{bool: v}
}

// float64AsLogsGroupByTotal is a convenience function that returns float64 wrapped in LogsGroupByTotal
func float64AsLogsGroupByTotal(v *float64) LogsGroupByTotal {
	return LogsGroupByTotal{float64: v}
}

// stringAsLogsGroupByTotal is a convenience function that returns string wrapped in LogsGroupByTotal
func stringAsLogsGroupByTotal(v *string) LogsGroupByTotal {
	return LogsGroupByTotal{string: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into bool
	err = json.Unmarshal(data, &dst.bool)
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			match++
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal data into float64
	err = json.Unmarshal(data, &dst.float64)
	if err == nil {
		jsonfloat64, _ := json.Marshal(dst.float64)
		if string(jsonfloat64) == "{}" { // empty struct
			dst.float64 = nil
		} else {
			match++
		}
	} else {
		dst.float64 = nil
	}

	// try to unmarshal data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			match++
		}
	} else {
		dst.string = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.bool = nil
		dst.float64 = nil
		dst.string = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(LogsGroupByTotal)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(LogsGroupByTotal)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsGroupByTotal) MarshalJSON() ([]byte, error) {
	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.float64 != nil {
		return json.Marshal(&src.float64)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsGroupByTotal) GetActualInstance() interface{} {
	if obj.bool != nil {
		return obj.bool
	}

	if obj.float64 != nil {
		return obj.float64
	}

	if obj.string != nil {
		return obj.string
	}

	// all schemas are nil
	return nil
}

type NullableLogsGroupByTotal struct {
	value *LogsGroupByTotal
	isSet bool
}

func (v NullableLogsGroupByTotal) Get() *LogsGroupByTotal {
	return v.value
}

func (v *NullableLogsGroupByTotal) Set(val *LogsGroupByTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGroupByTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsGroupByTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGroupByTotal(val *LogsGroupByTotal) *NullableLogsGroupByTotal {
	return &NullableLogsGroupByTotal{value: val, isSet: true}
}

func (v NullableLogsGroupByTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGroupByTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}