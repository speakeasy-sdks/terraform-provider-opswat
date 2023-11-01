// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NodesStatusRequest struct {
	// Generated `session_id` from [Login](/mdcore/metadefender-core/ref#userlogin) call can be used as an `apikey` for API calls that require authentication.
	//
	Apikey string `header:"style=simple,explode=false,name=apikey"`
}

func (o *NodesStatusRequest) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

// NodesStatus405ApplicationJSON - The user has no rights for this operation.
type NodesStatus405ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *NodesStatus405ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// NodesStatus403ApplicationJSON - Invalid user information or Not Allowed
type NodesStatus403ApplicationJSON struct {
	// Error reason
	Err *string `json:"err,omitempty"`
}

func (o *NodesStatus403ApplicationJSON) GetErr() *string {
	if o == nil {
		return nil
	}
	return o.Err
}

// NodesStatus200ApplicationJSONStatusesEnginesEngineType - Engine's type:
//   - av
//   - archive
//   - filetype
type NodesStatus200ApplicationJSONStatusesEnginesEngineType string

const (
	NodesStatus200ApplicationJSONStatusesEnginesEngineTypeAv       NodesStatus200ApplicationJSONStatusesEnginesEngineType = "av"
	NodesStatus200ApplicationJSONStatusesEnginesEngineTypeArchive  NodesStatus200ApplicationJSONStatusesEnginesEngineType = "archive"
	NodesStatus200ApplicationJSONStatusesEnginesEngineTypeFiletype NodesStatus200ApplicationJSONStatusesEnginesEngineType = "filetype"
)

func (e NodesStatus200ApplicationJSONStatusesEnginesEngineType) ToPointer() *NodesStatus200ApplicationJSONStatusesEnginesEngineType {
	return &e
}

func (e *NodesStatus200ApplicationJSONStatusesEnginesEngineType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "av":
		fallthrough
	case "archive":
		fallthrough
	case "filetype":
		*e = NodesStatus200ApplicationJSONStatusesEnginesEngineType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodesStatus200ApplicationJSONStatusesEnginesEngineType: %v", v)
	}
}

// NodesStatus200ApplicationJSONStatusesEngines - Engine summary
type NodesStatus200ApplicationJSONStatusesEngines struct {
	// If used by at least one engine
	Active *bool `json:"active,omitempty"`
	// The database version for this engine
	DbVer *string `json:"db_ver,omitempty"`
	// The database definition time for this engine
	DefTime *string `json:"def_time,omitempty"`
	// Engine name
	EngName *string `json:"eng_name,omitempty"`
	// Engine's version (format differs from one engine to another).
	EngVer *string `json:"eng_ver,omitempty"`
	// Engine's type:
	//   * av
	//   * archive
	//   * filetype
	//
	EngineType *NodesStatus200ApplicationJSONStatusesEnginesEngineType `json:"engine_type,omitempty"`
	// Engine internal ID
	ID *string `json:"id,omitempty"`
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetDbVer() *string {
	if o == nil {
		return nil
	}
	return o.DbVer
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetDefTime() *string {
	if o == nil {
		return nil
	}
	return o.DefTime
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetEngName() *string {
	if o == nil {
		return nil
	}
	return o.EngName
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetEngVer() *string {
	if o == nil {
		return nil
	}
	return o.EngVer
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetEngineType() *NodesStatus200ApplicationJSONStatusesEnginesEngineType {
	if o == nil {
		return nil
	}
	return o.EngineType
}

func (o *NodesStatus200ApplicationJSONStatusesEngines) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type NodesStatus200ApplicationJSONStatusesIssues struct {
	// Text detailing the issue.
	Description *string `json:"description,omitempty"`
	// How important is the reported issue.
	Severity *string `json:"severity,omitempty"`
}

func (o *NodesStatus200ApplicationJSONStatusesIssues) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *NodesStatus200ApplicationJSONStatusesIssues) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

// NodesStatus200ApplicationJSONStatuses - Node status
type NodesStatus200ApplicationJSONStatuses struct {
	// The location of the Node. If local, is empty string.
	Address *string `json:"address,omitempty"`
	// The number of available RAM in this system.
	AvailableMem *int64 `json:"available_mem,omitempty"`
	// The number of CPU Cores allocated to this Node.
	CPUCores *int64 `json:"cpu_cores,omitempty"`
	// Summary of each engine status deployed on this Node.
	Engines []NodesStatus200ApplicationJSONStatusesEngines `json:"engines,omitempty"`
	// Reported available disk on that Node (in bytes).
	FreeDiskSpace *int64 `json:"free_disk_space,omitempty"`
	// Node identfier
	ID *string `json:"id,omitempty"`
	// A list of all potentials problems on that Node.
	Issues []NodesStatus200ApplicationJSONStatusesIssues `json:"issues,omitempty"`
	// Current CPU utilization on this Node (in percentage).
	Load *int64 `json:"load,omitempty"`
	// Current used OS
	Os *string `json:"os,omitempty"`
	// Current load on the Node, how many files are in the queue
	ScanQueue *int64 `json:"scan_queue,omitempty"`
	// The amount of disk space is allocated on this Node (in Byte).
	TotalDiskSpace *int64 `json:"total_disk_space,omitempty"`
	// How much memory is allocated on this Node (in MB).
	TotalMem *int64 `json:"total_mem,omitempty"`
	// The maximum queue size is allowed per Node.
	TotalScanQueue *int64 `json:"total_scan_queue,omitempty"`
	// How long this Node is in operation (in second).
	Uptime *int64 `json:"uptime,omitempty"`
	// Product version
	Version *string `json:"version,omitempty"`
}

func (o *NodesStatus200ApplicationJSONStatuses) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *NodesStatus200ApplicationJSONStatuses) GetAvailableMem() *int64 {
	if o == nil {
		return nil
	}
	return o.AvailableMem
}

func (o *NodesStatus200ApplicationJSONStatuses) GetCPUCores() *int64 {
	if o == nil {
		return nil
	}
	return o.CPUCores
}

func (o *NodesStatus200ApplicationJSONStatuses) GetEngines() []NodesStatus200ApplicationJSONStatusesEngines {
	if o == nil {
		return nil
	}
	return o.Engines
}

func (o *NodesStatus200ApplicationJSONStatuses) GetFreeDiskSpace() *int64 {
	if o == nil {
		return nil
	}
	return o.FreeDiskSpace
}

func (o *NodesStatus200ApplicationJSONStatuses) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NodesStatus200ApplicationJSONStatuses) GetIssues() []NodesStatus200ApplicationJSONStatusesIssues {
	if o == nil {
		return nil
	}
	return o.Issues
}

func (o *NodesStatus200ApplicationJSONStatuses) GetLoad() *int64 {
	if o == nil {
		return nil
	}
	return o.Load
}

func (o *NodesStatus200ApplicationJSONStatuses) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *NodesStatus200ApplicationJSONStatuses) GetScanQueue() *int64 {
	if o == nil {
		return nil
	}
	return o.ScanQueue
}

func (o *NodesStatus200ApplicationJSONStatuses) GetTotalDiskSpace() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalDiskSpace
}

func (o *NodesStatus200ApplicationJSONStatuses) GetTotalMem() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalMem
}

func (o *NodesStatus200ApplicationJSONStatuses) GetTotalScanQueue() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalScanQueue
}

func (o *NodesStatus200ApplicationJSONStatuses) GetUptime() *int64 {
	if o == nil {
		return nil
	}
	return o.Uptime
}

func (o *NodesStatus200ApplicationJSONStatuses) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

type NodesStatus200ApplicationJSON struct {
	// Configuration option if additional external nodes could be connected.
	ExternalNodesAllowed *bool `json:"external_nodes_allowed,omitempty"`
	// Remaining available slots to connect to this instance.
	MaxNodeCount *int64 `json:"max_node_count,omitempty"`
	// Array with a status for each attached node.
	Statuses []NodesStatus200ApplicationJSONStatuses `json:"statuses,omitempty"`
}

func (o *NodesStatus200ApplicationJSON) GetExternalNodesAllowed() *bool {
	if o == nil {
		return nil
	}
	return o.ExternalNodesAllowed
}

func (o *NodesStatus200ApplicationJSON) GetMaxNodeCount() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxNodeCount
}

func (o *NodesStatus200ApplicationJSON) GetStatuses() []NodesStatus200ApplicationJSONStatuses {
	if o == nil {
		return nil
	}
	return o.Statuses
}

type NodesStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// An array with all the engines and their details.
	NodesStatus200ApplicationJSONObjects []NodesStatus200ApplicationJSON
	// Invalid user information or Not Allowed
	NodesStatus403ApplicationJSONObject *NodesStatus403ApplicationJSON
	// The user has no rights for this operation.
	NodesStatus405ApplicationJSONObject *NodesStatus405ApplicationJSON
}

func (o *NodesStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *NodesStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *NodesStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *NodesStatusResponse) GetNodesStatus200ApplicationJSONObjects() []NodesStatus200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.NodesStatus200ApplicationJSONObjects
}

func (o *NodesStatusResponse) GetNodesStatus403ApplicationJSONObject() *NodesStatus403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.NodesStatus403ApplicationJSONObject
}

func (o *NodesStatusResponse) GetNodesStatus405ApplicationJSONObject() *NodesStatus405ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.NodesStatus405ApplicationJSONObject
}
