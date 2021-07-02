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
// source: google/ads/googleads/v8/resources/detail_placement_view.proto

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

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *string `protobuf:"bytes,7,opt,name=placement,proto3,oneof" json:"placement,omitempty"`
	// Output only. The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *string `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	// Output only. URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *string `protobuf:"bytes,9,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3,oneof" json:"group_placement_target_url,omitempty"`
	// Output only. URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *string `protobuf:"bytes,10,opt,name=target_url,json=targetUrl,proto3,oneof" json:"target_url,omitempty"`
	// Output only. Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v8.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
}

func (x *DetailPlacementView) Reset() {
	*x = DetailPlacementView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_detail_placement_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailPlacementView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailPlacementView) ProtoMessage() {}

func (x *DetailPlacementView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_detail_placement_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailPlacementView.ProtoReflect.Descriptor instead.
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescGZIP(), []int{0}
}

func (x *DetailPlacementView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *DetailPlacementView) GetPlacement() string {
	if x != nil && x.Placement != nil {
		return *x.Placement
	}
	return ""
}

func (x *DetailPlacementView) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *DetailPlacementView) GetGroupPlacementTargetUrl() string {
	if x != nil && x.GroupPlacementTargetUrl != nil {
		return *x.GroupPlacementTargetUrl
	}
	return ""
}

func (x *DetailPlacementView) GetTargetUrl() string {
	if x != nil && x.TargetUrl != nil {
		return *x.TargetUrl
	}
	return ""
}

func (x *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if x != nil {
		return x.PlacementType
	}
	return enums.PlacementTypeEnum_PlacementType(0)
}

var File_google_ads_googleads_v8_resources_detail_placement_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf1, 0x04, 0x0a, 0x13, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x34, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x1a, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x02, 0x52, 0x17, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x27, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x6a, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x3a, 0x80, 0x01, 0xea, 0x41, 0x7d, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x4d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x42, 0x85, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x18,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescData = file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_detail_placement_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_resources_detail_placement_view_proto_goTypes = []interface{}{
	(*DetailPlacementView)(nil),                // 0: google.ads.googleads.v8.resources.DetailPlacementView
	(enums.PlacementTypeEnum_PlacementType)(0), // 1: google.ads.googleads.v8.enums.PlacementTypeEnum.PlacementType
}
var file_google_ads_googleads_v8_resources_detail_placement_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.DetailPlacementView.placement_type:type_name -> google.ads.googleads.v8.enums.PlacementTypeEnum.PlacementType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_detail_placement_view_proto_init() }
func file_google_ads_googleads_v8_resources_detail_placement_view_proto_init() {
	if File_google_ads_googleads_v8_resources_detail_placement_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_detail_placement_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailPlacementView); i {
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
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_detail_placement_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_detail_placement_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_detail_placement_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_detail_placement_view_proto = out.File
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_detail_placement_view_proto_depIdxs = nil
}
