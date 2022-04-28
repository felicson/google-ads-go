// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: google/ads/googleads/v9/resources/custom_interest.proto

package resources

import (
	enums "github.com/felicson/google-ads-go/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A custom interest. This is a list of users by interest.
type CustomInterest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the custom interest.
	// Custom interest resource names have the form:
	//
	// `customers/{customer_id}/customInterests/{custom_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Id of the custom interest.
	Id *int64 `protobuf:"varint,8,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Status of this custom interest. Indicates whether the custom interest is
	// enabled or removed.
	Status enums.CustomInterestStatusEnum_CustomInterestStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v9.enums.CustomInterestStatusEnum_CustomInterestStatus" json:"status,omitempty"`
	// Name of the custom interest. It should be unique across the same custom
	// affinity audience.
	// This field is required for create operations.
	Name *string `protobuf:"bytes,9,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Type of the custom interest, CUSTOM_AFFINITY or CUSTOM_INTENT.
	// By default the type is set to CUSTOM_AFFINITY.
	Type enums.CustomInterestTypeEnum_CustomInterestType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v9.enums.CustomInterestTypeEnum_CustomInterestType" json:"type,omitempty"`
	// Description of this custom interest audience.
	Description *string `protobuf:"bytes,10,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// List of custom interest members that this custom interest is composed of.
	// Members can be added during CustomInterest creation. If members are
	// presented in UPDATE operation, existing members will be overridden.
	Members []*CustomInterestMember `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *CustomInterest) Reset() {
	*x = CustomInterest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomInterest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomInterest) ProtoMessage() {}

func (x *CustomInterest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomInterest.ProtoReflect.Descriptor instead.
func (*CustomInterest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescGZIP(), []int{0}
}

func (x *CustomInterest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomInterest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *CustomInterest) GetStatus() enums.CustomInterestStatusEnum_CustomInterestStatus {
	if x != nil {
		return x.Status
	}
	return enums.CustomInterestStatusEnum_CustomInterestStatus(0)
}

func (x *CustomInterest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CustomInterest) GetType() enums.CustomInterestTypeEnum_CustomInterestType {
	if x != nil {
		return x.Type
	}
	return enums.CustomInterestTypeEnum_CustomInterestType(0)
}

func (x *CustomInterest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CustomInterest) GetMembers() []*CustomInterestMember {
	if x != nil {
		return x.Members
	}
	return nil
}

// A member of custom interest audience. A member can be a keyword or url.
// It is immutable, that is, it can only be created or removed but not changed.
type CustomInterestMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of custom interest member, KEYWORD or URL.
	MemberType enums.CustomInterestMemberTypeEnum_CustomInterestMemberType `protobuf:"varint,1,opt,name=member_type,json=memberType,proto3,enum=google.ads.googleads.v9.enums.CustomInterestMemberTypeEnum_CustomInterestMemberType" json:"member_type,omitempty"`
	// Keyword text when member_type is KEYWORD or URL string when
	// member_type is URL.
	Parameter *string `protobuf:"bytes,3,opt,name=parameter,proto3,oneof" json:"parameter,omitempty"`
}

func (x *CustomInterestMember) Reset() {
	*x = CustomInterestMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomInterestMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomInterestMember) ProtoMessage() {}

func (x *CustomInterestMember) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomInterestMember.ProtoReflect.Descriptor instead.
func (*CustomInterestMember) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescGZIP(), []int{1}
}

func (x *CustomInterestMember) GetMemberType() enums.CustomInterestMemberTypeEnum_CustomInterestMemberType {
	if x != nil {
		return x.MemberType
	}
	return enums.CustomInterestMemberTypeEnum_CustomInterestMemberType(0)
}

func (x *CustomInterestMember) GetParameter() string {
	if x != nil && x.Parameter != nil {
		return *x.Parameter
	}
	return ""
}

var File_google_ads_googleads_v9_resources_custom_interest_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_resources_custom_interest_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x04,
	0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x29, 0x0a,
	0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x64, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x5c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x6a, 0xea, 0x41, 0x67, 0x0a, 0x27, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x0b,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x42, 0x80, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x13,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea,
	0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescData = file_google_ads_googleads_v9_resources_custom_interest_proto_rawDesc
)

func file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_resources_custom_interest_proto_rawDescData
}

var file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v9_resources_custom_interest_proto_goTypes = []interface{}{
	(*CustomInterest)(nil),                                           // 0: google.ads.googleads.v9.resources.CustomInterest
	(*CustomInterestMember)(nil),                                     // 1: google.ads.googleads.v9.resources.CustomInterestMember
	(enums.CustomInterestStatusEnum_CustomInterestStatus)(0),         // 2: google.ads.googleads.v9.enums.CustomInterestStatusEnum.CustomInterestStatus
	(enums.CustomInterestTypeEnum_CustomInterestType)(0),             // 3: google.ads.googleads.v9.enums.CustomInterestTypeEnum.CustomInterestType
	(enums.CustomInterestMemberTypeEnum_CustomInterestMemberType)(0), // 4: google.ads.googleads.v9.enums.CustomInterestMemberTypeEnum.CustomInterestMemberType
}
var file_google_ads_googleads_v9_resources_custom_interest_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v9.resources.CustomInterest.status:type_name -> google.ads.googleads.v9.enums.CustomInterestStatusEnum.CustomInterestStatus
	3, // 1: google.ads.googleads.v9.resources.CustomInterest.type:type_name -> google.ads.googleads.v9.enums.CustomInterestTypeEnum.CustomInterestType
	1, // 2: google.ads.googleads.v9.resources.CustomInterest.members:type_name -> google.ads.googleads.v9.resources.CustomInterestMember
	4, // 3: google.ads.googleads.v9.resources.CustomInterestMember.member_type:type_name -> google.ads.googleads.v9.enums.CustomInterestMemberTypeEnum.CustomInterestMemberType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_resources_custom_interest_proto_init() }
func file_google_ads_googleads_v9_resources_custom_interest_proto_init() {
	if File_google_ads_googleads_v9_resources_custom_interest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomInterest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomInterestMember); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v9_resources_custom_interest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_resources_custom_interest_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_resources_custom_interest_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v9_resources_custom_interest_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_resources_custom_interest_proto = out.File
	file_google_ads_googleads_v9_resources_custom_interest_proto_rawDesc = nil
	file_google_ads_googleads_v9_resources_custom_interest_proto_goTypes = nil
	file_google_ads_googleads_v9_resources_custom_interest_proto_depIdxs = nil
}
