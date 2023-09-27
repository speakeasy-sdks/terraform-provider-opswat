// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// YaraSourcesObjectHTTPSourcesState - Defining if this source is being used or not
type YaraSourcesObjectHTTPSourcesState string

const (
	YaraSourcesObjectHTTPSourcesStateDisabled YaraSourcesObjectHTTPSourcesState = "disabled"
	YaraSourcesObjectHTTPSourcesStateEnabled  YaraSourcesObjectHTTPSourcesState = "enabled"
)

func (e YaraSourcesObjectHTTPSourcesState) ToPointer() *YaraSourcesObjectHTTPSourcesState {
	return &e
}

func (e *YaraSourcesObjectHTTPSourcesState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled":
		fallthrough
	case "enabled":
		*e = YaraSourcesObjectHTTPSourcesState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for YaraSourcesObjectHTTPSourcesState: %v", v)
	}
}

// YaraSourcesObjectHTTPSources - Object defining the URL and if this source is being used.
type YaraSourcesObjectHTTPSources struct {
	// The location (URL) of the package.
	Source *string `json:"source,omitempty"`
	// Defining if this source is being used or not
	State *YaraSourcesObjectHTTPSourcesState `json:"state,omitempty"`
}

// YaraSourcesObjectLocalSourcesState - Defining if this source is being used or not
type YaraSourcesObjectLocalSourcesState string

const (
	YaraSourcesObjectLocalSourcesStateDisabled YaraSourcesObjectLocalSourcesState = "disabled"
	YaraSourcesObjectLocalSourcesStateEnabled  YaraSourcesObjectLocalSourcesState = "enabled"
)

func (e YaraSourcesObjectLocalSourcesState) ToPointer() *YaraSourcesObjectLocalSourcesState {
	return &e
}

func (e *YaraSourcesObjectLocalSourcesState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled":
		fallthrough
	case "enabled":
		*e = YaraSourcesObjectLocalSourcesState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for YaraSourcesObjectLocalSourcesState: %v", v)
	}
}

// YaraSourcesObjectLocalSources - Object defining the location of the files and if this source is being used.
type YaraSourcesObjectLocalSources struct {
	// The location (file path) of the files.
	Source *string `json:"source,omitempty"`
	// Defining if this source is being used or not
	State *YaraSourcesObjectLocalSourcesState `json:"state,omitempty"`
}

// YaraSourcesObject - Object describing the yara sources.
type YaraSourcesObject struct {
	// A list of all remote sources.
	HTTPSources []YaraSourcesObjectHTTPSources `json:"http_sources,omitempty"`
	// A list of all locals sources
	LocalSources []YaraSourcesObjectLocalSources `json:"local_sources,omitempty"`
}
