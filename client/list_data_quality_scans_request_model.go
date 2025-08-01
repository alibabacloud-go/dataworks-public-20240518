// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListDataQualityScansRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataQualityScansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityScansRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityScansRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDataQualityScansRequest
	GetSortBy() *string
	SetTable(v string) *ListDataQualityScansRequest
	GetTable() *string
}

type ListDataQualityScansRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// ModifyTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// video_album
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s ListDataQualityScansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansRequest) GetName() *string {
	return s.Name
}

func (s *ListDataQualityScansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityScansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityScansRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityScansRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDataQualityScansRequest) GetTable() *string {
	return s.Table
}

func (s *ListDataQualityScansRequest) SetName(v string) *ListDataQualityScansRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityScansRequest) SetPageNumber(v int32) *ListDataQualityScansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityScansRequest) SetPageSize(v int32) *ListDataQualityScansRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityScansRequest) SetProjectId(v int64) *ListDataQualityScansRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityScansRequest) SetSortBy(v string) *ListDataQualityScansRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataQualityScansRequest) SetTable(v string) *ListDataQualityScansRequest {
	s.Table = &v
	return s
}

func (s *ListDataQualityScansRequest) Validate() error {
	return dara.Validate(s)
}
