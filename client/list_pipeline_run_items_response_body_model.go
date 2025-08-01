// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListPipelineRunItemsResponseBodyPagingInfo) *ListPipelineRunItemsResponseBody
	GetPagingInfo() *ListPipelineRunItemsResponseBodyPagingInfo
	SetRequestId(v string) *ListPipelineRunItemsResponseBody
	GetRequestId() *string
}

type ListPipelineRunItemsResponseBody struct {
	PagingInfo *ListPipelineRunItemsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPipelineRunItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunItemsResponseBody) GetPagingInfo() *ListPipelineRunItemsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListPipelineRunItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineRunItemsResponseBody) SetPagingInfo(v *ListPipelineRunItemsResponseBodyPagingInfo) *ListPipelineRunItemsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListPipelineRunItemsResponseBody) SetRequestId(v string) *ListPipelineRunItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPipelineRunItemsResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize         *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PipelineRunItems []*ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems `json:"PipelineRunItems,omitempty" xml:"PipelineRunItems,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelineRunItemsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunItemsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) GetPipelineRunItems() []*ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	return s.PipelineRunItems
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) SetPageNumber(v int32) *ListPipelineRunItemsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) SetPageSize(v int32) *ListPipelineRunItemsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) SetPipelineRunItems(v []*ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) *ListPipelineRunItemsResponseBodyPagingInfo {
	s.PipelineRunItems = v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) SetTotalCount(v int32) *ListPipelineRunItemsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems struct {
	// 发布包创建时间戳
	//
	// example:
	//
	// 1724984066000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 创建人
	//
	// example:
	//
	// Error Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1724984066000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// { "version": "1.1.0", "kind": "Node", "spec": { "nodes": [ { "recurrence": "Normal", "id": "860438872620113XXXX", "timeout": 0, "instanceMode": "T+1", "rerunMode": "Allowed", "rerunTimes": 3, "rerunInterval": 180000, "datasource": { "name": "odps_test", "type": "odps" }, "script": { "language": "odps-sql", "path": "XX/OpenAPI_Test/ODPS_SQL_Test", "runtime": { "command": "ODPS_SQL", "commandTypeId": 10 }, "content": "select now();", "id": "853573334108680XXXX" }, "trigger": { "type": "Scheduler", "id": "543680677872062XXXX", "cron": "00 00 00 	- 	- ?", "startTime": "1970-01-01 00:00:00", "endTime": "9999-01-01 00:00:00", "timezone": "Asia/Shanghai", "delaySeconds": 0 }, "runtimeResource": { "resourceGroup": "S_res_group_XXXX_XXXX", "id": "623731286945488XXXX", "resourceGroupId": "7201XXXX" }, "name": "ODPS_SQL_Test", "owner": "110755000425XXXX", "metadata": { "owner": "110755000425XXXX", "ownerName": "XXXXX@test.XXX.com", "projectId": "307XXX" }, "inputs": { "nodeOutputs": [ { "data": "lwttest_standard_root", "artifactType": "NodeOutput" } ] }, "outputs": { "nodeOutputs": [ { "data": "860438872620113XXXX", "artifactType": "NodeOutput", "refTableName": "ODPS_SQL_Test", "isDefault": true } ] } } ], "flow": [ { "nodeId": "860438872620113XXXX", "depends": [ { "type": "Normal", "output": "lwttest_standard_root" } ] } ] }, "metadata": { "uuid": "860438872620113XXXX" } }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// 发布流程状态
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Node
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 项目Id
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GoString() string {
	return s.String()
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetId() *int64 {
	return s.Id
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetMessage() *string {
	return s.Message
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetName() *string {
	return s.Name
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetSpec() *string {
	return s.Spec
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetType() *string {
	return s.Type
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) GetVersion() *int64 {
	return s.Version
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetCreateTime(v int64) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.CreateTime = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetId(v int64) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Id = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetMessage(v string) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Message = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetModifyTime(v int64) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.ModifyTime = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetName(v string) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Name = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetSpec(v string) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Spec = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetStatus(v string) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Status = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetType(v string) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Type = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) SetVersion(v int64) *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems {
	s.Version = &v
	return s
}

func (s *ListPipelineRunItemsResponseBodyPagingInfoPipelineRunItems) Validate() error {
	return dara.Validate(s)
}
