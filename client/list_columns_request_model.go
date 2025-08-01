// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListColumnsRequest
	GetComment() *string
	SetName(v string) *ListColumnsRequest
	GetName() *string
	SetOrder(v string) *ListColumnsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListColumnsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListColumnsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListColumnsRequest
	GetSortBy() *string
	SetTableId(v string) *ListColumnsRequest
	GetTableId() *string
}

type ListColumnsRequest struct {
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// test_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Position
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the table to which the columns belong. You can call the ListTables operation to query the ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:11075xxxx::test_project:test_schema:test_table
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s ListColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListColumnsRequest) GetComment() *string {
	return s.Comment
}

func (s *ListColumnsRequest) GetName() *string {
	return s.Name
}

func (s *ListColumnsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListColumnsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListColumnsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListColumnsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListColumnsRequest) GetTableId() *string {
	return s.TableId
}

func (s *ListColumnsRequest) SetComment(v string) *ListColumnsRequest {
	s.Comment = &v
	return s
}

func (s *ListColumnsRequest) SetName(v string) *ListColumnsRequest {
	s.Name = &v
	return s
}

func (s *ListColumnsRequest) SetOrder(v string) *ListColumnsRequest {
	s.Order = &v
	return s
}

func (s *ListColumnsRequest) SetPageNumber(v int32) *ListColumnsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListColumnsRequest) SetPageSize(v int32) *ListColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *ListColumnsRequest) SetSortBy(v string) *ListColumnsRequest {
	s.SortBy = &v
	return s
}

func (s *ListColumnsRequest) SetTableId(v string) *ListColumnsRequest {
	s.TableId = &v
	return s
}

func (s *ListColumnsRequest) Validate() error {
	return dara.Validate(s)
}
