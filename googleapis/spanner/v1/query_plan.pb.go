// Code generated by protoc-gen-go.
// source: google/spanner/v1/query_plan.proto
// DO NOT EDIT!

package spanner

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The kind of [PlanNode][google.spanner.v1.PlanNode]. Distinguishes between the two different kinds of
// nodes that can appear in a query plan.
type PlanNode_Kind int32

const (
	// Not specified.
	PlanNode_KIND_UNSPECIFIED PlanNode_Kind = 0
	// Denotes a Relational operator node in the expression tree. Relational
	// operators represent iterative processing of rows during query execution.
	// For example, a `TableScan` operation that reads rows from a table.
	PlanNode_RELATIONAL PlanNode_Kind = 1
	// Denotes a Scalar node in the expression tree. Scalar nodes represent
	// non-iterable entities in the query plan. For example, constants or
	// arithmetic operators appearing inside predicate expressions or references
	// to column names.
	PlanNode_SCALAR PlanNode_Kind = 2
)

var PlanNode_Kind_name = map[int32]string{
	0: "KIND_UNSPECIFIED",
	1: "RELATIONAL",
	2: "SCALAR",
}
var PlanNode_Kind_value = map[string]int32{
	"KIND_UNSPECIFIED": 0,
	"RELATIONAL":       1,
	"SCALAR":           2,
}

func (x PlanNode_Kind) String() string {
	return proto.EnumName(PlanNode_Kind_name, int32(x))
}
func (PlanNode_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

// Node information for nodes appearing in a [QueryPlan.plan_nodes][google.spanner.v1.QueryPlan.plan_nodes].
type PlanNode struct {
	// The `PlanNode`'s index in [node list][google.spanner.v1.QueryPlan.plan_nodes].
	Index int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// Used to determine the type of node. May be needed for visualizing
	// different kinds of nodes differently. For example, If the node is a
	// [SCALAR][google.spanner.v1.PlanNode.Kind.SCALAR] node, it will have a condensed representation
	// which can be used to directly embed a description of the node in its
	// parent.
	Kind PlanNode_Kind `protobuf:"varint,2,opt,name=kind,enum=google.spanner.v1.PlanNode_Kind" json:"kind,omitempty"`
	// The display name for the node.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// List of child node `index`es and their relationship to this parent.
	ChildLinks []*PlanNode_ChildLink `protobuf:"bytes,4,rep,name=child_links,json=childLinks" json:"child_links,omitempty"`
	// Condensed representation for [SCALAR][google.spanner.v1.PlanNode.Kind.SCALAR] nodes.
	ShortRepresentation *PlanNode_ShortRepresentation `protobuf:"bytes,5,opt,name=short_representation,json=shortRepresentation" json:"short_representation,omitempty"`
	// Attributes relevant to the node contained in a group of key-value pairs.
	// For example, a Parameter Reference node could have the following
	// information in its metadata:
	//
	//     {
	//       "parameter_reference": "param1",
	//       "parameter_type": "array"
	//     }
	Metadata *google_protobuf1.Struct `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	// The execution statistics associated with the node, contained in a group of
	// key-value pairs. Only present if the plan was returned as a result of a
	// profile query. For example, number of executions, number of rows/time per
	// execution etc.
	ExecutionStats *google_protobuf1.Struct `protobuf:"bytes,7,opt,name=execution_stats,json=executionStats" json:"execution_stats,omitempty"`
}

func (m *PlanNode) Reset()                    { *m = PlanNode{} }
func (m *PlanNode) String() string            { return proto.CompactTextString(m) }
func (*PlanNode) ProtoMessage()               {}
func (*PlanNode) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PlanNode) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PlanNode) GetKind() PlanNode_Kind {
	if m != nil {
		return m.Kind
	}
	return PlanNode_KIND_UNSPECIFIED
}

func (m *PlanNode) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *PlanNode) GetChildLinks() []*PlanNode_ChildLink {
	if m != nil {
		return m.ChildLinks
	}
	return nil
}

func (m *PlanNode) GetShortRepresentation() *PlanNode_ShortRepresentation {
	if m != nil {
		return m.ShortRepresentation
	}
	return nil
}

func (m *PlanNode) GetMetadata() *google_protobuf1.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PlanNode) GetExecutionStats() *google_protobuf1.Struct {
	if m != nil {
		return m.ExecutionStats
	}
	return nil
}

// Metadata associated with a parent-child relationship appearing in a
// [PlanNode][google.spanner.v1.PlanNode].
type PlanNode_ChildLink struct {
	// The node to which the link points.
	ChildIndex int32 `protobuf:"varint,1,opt,name=child_index,json=childIndex" json:"child_index,omitempty"`
	// The type of the link. For example, in Hash Joins this could be used to
	// distinguish between the build child and the probe child, or in the case
	// of the child being an output variable, to represent the tag associated
	// with the output variable.
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// Only present if the child node is [SCALAR][google.spanner.v1.PlanNode.Kind.SCALAR] and corresponds
	// to an output variable of the parent node. The field carries the name of
	// the output variable.
	// For example, a `TableScan` operator that reads rows from a table will
	// have child links to the `SCALAR` nodes representing the output variables
	// created for each column that is read by the operator. The corresponding
	// `variable` fields will be set to the variable names assigned to the
	// columns.
	Variable string `protobuf:"bytes,3,opt,name=variable" json:"variable,omitempty"`
}

func (m *PlanNode_ChildLink) Reset()                    { *m = PlanNode_ChildLink{} }
func (m *PlanNode_ChildLink) String() string            { return proto.CompactTextString(m) }
func (*PlanNode_ChildLink) ProtoMessage()               {}
func (*PlanNode_ChildLink) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *PlanNode_ChildLink) GetChildIndex() int32 {
	if m != nil {
		return m.ChildIndex
	}
	return 0
}

func (m *PlanNode_ChildLink) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlanNode_ChildLink) GetVariable() string {
	if m != nil {
		return m.Variable
	}
	return ""
}

// Condensed representation of a node and its subtree. Only present for
// `SCALAR` [PlanNode(s)][google.spanner.v1.PlanNode].
type PlanNode_ShortRepresentation struct {
	// A string representation of the expression subtree rooted at this node.
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// A mapping of (subquery variable name) -> (subquery node id) for cases
	// where the `description` string of this node references a `SCALAR`
	// subquery contained in the expression subtree rooted at this node. The
	// referenced `SCALAR` subquery may not necessarily be a direct child of
	// this node.
	Subqueries map[string]int32 `protobuf:"bytes,2,rep,name=subqueries" json:"subqueries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *PlanNode_ShortRepresentation) Reset()                    { *m = PlanNode_ShortRepresentation{} }
func (m *PlanNode_ShortRepresentation) String() string            { return proto.CompactTextString(m) }
func (*PlanNode_ShortRepresentation) ProtoMessage()               {}
func (*PlanNode_ShortRepresentation) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

func (m *PlanNode_ShortRepresentation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PlanNode_ShortRepresentation) GetSubqueries() map[string]int32 {
	if m != nil {
		return m.Subqueries
	}
	return nil
}

// Contains an ordered list of nodes appearing in the query plan.
type QueryPlan struct {
	// The nodes in the query plan. Plan nodes are returned in pre-order starting
	// with the plan root. Each [PlanNode][google.spanner.v1.PlanNode]'s `id` corresponds to its index in
	// `plan_nodes`.
	PlanNodes []*PlanNode `protobuf:"bytes,1,rep,name=plan_nodes,json=planNodes" json:"plan_nodes,omitempty"`
}

func (m *QueryPlan) Reset()                    { *m = QueryPlan{} }
func (m *QueryPlan) String() string            { return proto.CompactTextString(m) }
func (*QueryPlan) ProtoMessage()               {}
func (*QueryPlan) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *QueryPlan) GetPlanNodes() []*PlanNode {
	if m != nil {
		return m.PlanNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*PlanNode)(nil), "google.spanner.v1.PlanNode")
	proto.RegisterType((*PlanNode_ChildLink)(nil), "google.spanner.v1.PlanNode.ChildLink")
	proto.RegisterType((*PlanNode_ShortRepresentation)(nil), "google.spanner.v1.PlanNode.ShortRepresentation")
	proto.RegisterType((*QueryPlan)(nil), "google.spanner.v1.QueryPlan")
	proto.RegisterEnum("google.spanner.v1.PlanNode_Kind", PlanNode_Kind_name, PlanNode_Kind_value)
}

func init() { proto.RegisterFile("google/spanner/v1/query_plan.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0x26, 0xf9, 0x9a, 0x09, 0x4a, 0xc3, 0xb6, 0x08, 0x2b, 0x20, 0x61, 0x22, 0x21,
	0xe5, 0xca, 0x56, 0x5b, 0x2e, 0xaa, 0x22, 0x04, 0x69, 0x9b, 0xa2, 0xa8, 0x51, 0x08, 0x1b, 0xb8,
	0x41, 0x48, 0xd6, 0xc6, 0x5e, 0xdc, 0x55, 0x9c, 0x5d, 0xe3, 0x5d, 0x47, 0xcd, 0x4b, 0xf0, 0x7a,
	0xbc, 0x0e, 0xda, 0xf5, 0x0f, 0x85, 0xa0, 0x48, 0xdc, 0xcd, 0xec, 0x9c, 0x39, 0x9a, 0x39, 0x67,
	0x6c, 0xe8, 0x47, 0x42, 0x44, 0x31, 0xf5, 0x64, 0x42, 0x38, 0xa7, 0xa9, 0xb7, 0x3e, 0xf6, 0xbe,
	0x65, 0x34, 0xdd, 0xf8, 0x49, 0x4c, 0xb8, 0x9b, 0xa4, 0x42, 0x09, 0xf4, 0x30, 0xc7, 0xb8, 0x05,
	0xc6, 0x5d, 0x1f, 0xf7, 0x9e, 0x16, 0x6d, 0x24, 0x61, 0x1e, 0xe1, 0x5c, 0x28, 0xa2, 0x98, 0xe0,
	0x32, 0x6f, 0xa8, 0xaa, 0x26, 0x5b, 0x64, 0x5f, 0x3d, 0xa9, 0xd2, 0x2c, 0x50, 0x79, 0xb5, 0xff,
	0xbd, 0x09, 0xfb, 0xb3, 0x98, 0xf0, 0xa9, 0x08, 0x29, 0x3a, 0x82, 0x06, 0xe3, 0x21, 0xbd, 0xb3,
	0x2d, 0xc7, 0x1a, 0x34, 0x70, 0x9e, 0xa0, 0x97, 0x50, 0x5f, 0x32, 0x1e, 0xda, 0x35, 0xc7, 0x1a,
	0x74, 0x4e, 0x1c, 0x77, 0x6b, 0x00, 0xb7, 0x24, 0x70, 0x6f, 0x18, 0x0f, 0xb1, 0x41, 0xa3, 0xe7,
	0xf0, 0x20, 0x64, 0x32, 0x89, 0xc9, 0xc6, 0xe7, 0x64, 0x45, 0xed, 0x3d, 0xc7, 0x1a, 0xb4, 0x70,
	0xbb, 0x78, 0x9b, 0x92, 0x15, 0x45, 0xd7, 0xd0, 0x0e, 0x6e, 0x59, 0x1c, 0xfa, 0x31, 0xe3, 0x4b,
	0x69, 0xd7, 0x9d, 0xbd, 0x41, 0xfb, 0xe4, 0xc5, 0x2e, 0xfe, 0x4b, 0x0d, 0x9f, 0x30, 0xbe, 0xc4,
	0x10, 0x94, 0xa1, 0x44, 0x0b, 0x38, 0x92, 0xb7, 0x22, 0x55, 0x7e, 0x4a, 0x93, 0x94, 0x4a, 0xca,
	0x73, 0x01, 0xec, 0x86, 0x63, 0x0d, 0xda, 0x27, 0xde, 0x2e, 0xc2, 0xb9, 0xee, 0xc3, 0xbf, 0xb5,
	0xe1, 0x43, 0xb9, 0xfd, 0x88, 0x4e, 0x61, 0x7f, 0x45, 0x15, 0x09, 0x89, 0x22, 0x76, 0xd3, 0xf0,
	0x3e, 0x2e, 0x79, 0x4b, 0x61, 0xdd, 0xb9, 0x11, 0x16, 0x57, 0x40, 0xf4, 0x16, 0x0e, 0xe8, 0x1d,
	0x0d, 0x32, 0xcd, 0xe0, 0x4b, 0x45, 0x94, 0xb4, 0xff, 0xdf, 0xdd, 0xdb, 0xa9, 0xf0, 0x73, 0x0d,
	0xef, 0x7d, 0x81, 0x56, 0xb5, 0x33, 0x7a, 0x56, 0xea, 0x75, 0xdf, 0xa4, 0x5c, 0x88, 0xb1, 0x71,
	0x0a, 0x41, 0x5d, 0x6d, 0x12, 0x6a, 0x9c, 0x6a, 0x61, 0x13, 0xa3, 0x1e, 0xec, 0xaf, 0x49, 0xca,
	0xc8, 0x22, 0x2e, 0x3d, 0xa8, 0xf2, 0xde, 0x0f, 0x0b, 0x0e, 0xff, 0xa2, 0x00, 0x72, 0xa0, 0x1d,
	0x52, 0x19, 0xa4, 0x2c, 0x31, 0x3a, 0x5a, 0x85, 0x75, 0xbf, 0x9e, 0x90, 0x0f, 0x20, 0xb3, 0x85,
	0x3e, 0x4e, 0x46, 0xa5, 0x5d, 0x33, 0xce, 0xbd, 0xf9, 0x47, 0xa1, 0xdd, 0x79, 0xc5, 0x30, 0xe2,
	0x2a, 0xdd, 0xe0, 0x7b, 0x94, 0xbd, 0xd7, 0x70, 0xf0, 0x47, 0x19, 0x75, 0x61, 0x6f, 0x49, 0x37,
	0xc5, 0x34, 0x3a, 0xd4, 0xf7, 0xba, 0x26, 0x71, 0x96, 0x2f, 0xdc, 0xc0, 0x79, 0x72, 0x5e, 0x3b,
	0xb3, 0xfa, 0x67, 0x50, 0xd7, 0xb7, 0x88, 0x8e, 0xa0, 0x7b, 0x33, 0x9e, 0x5e, 0xf9, 0x9f, 0xa6,
	0xf3, 0xd9, 0xe8, 0x72, 0x7c, 0x3d, 0x1e, 0x5d, 0x75, 0xff, 0x43, 0x1d, 0x00, 0x3c, 0x9a, 0x0c,
	0x3f, 0x8e, 0xdf, 0x4f, 0x87, 0x93, 0xae, 0x85, 0x00, 0x9a, 0xf3, 0xcb, 0xe1, 0x64, 0x88, 0xbb,
	0xb5, 0xfe, 0x3b, 0x68, 0x7d, 0xd0, 0xdf, 0x9c, 0x9e, 0x1c, 0x9d, 0x03, 0xe8, 0x4f, 0xcf, 0xe7,
	0x22, 0xa4, 0xd2, 0xb6, 0xcc, 0x9a, 0x4f, 0x76, 0xac, 0x89, 0x5b, 0x49, 0x11, 0xc9, 0x8b, 0x00,
	0x1e, 0x05, 0x62, 0xb5, 0x0d, 0xbe, 0xe8, 0x54, 0xfc, 0x33, 0xed, 0xfe, 0xcc, 0xfa, 0x7c, 0x56,
	0x80, 0x22, 0x11, 0x13, 0x1e, 0xb9, 0x22, 0x8d, 0xbc, 0x88, 0x72, 0x73, 0x1b, 0x5e, 0x5e, 0x22,
	0x09, 0x93, 0xf7, 0x7e, 0x0b, 0xaf, 0x8a, 0x70, 0xd1, 0x34, 0xa0, 0xd3, 0x9f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x31, 0xcf, 0x6d, 0x19, 0x3a, 0x04, 0x00, 0x00,
}
