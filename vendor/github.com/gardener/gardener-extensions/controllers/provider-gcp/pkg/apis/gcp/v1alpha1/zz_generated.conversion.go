// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	gcp "github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/apis/gcp"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CloudControllerManagerConfig)(nil), (*gcp.CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudControllerManagerConfig_To_gcp_CloudControllerManagerConfig(a.(*CloudControllerManagerConfig), b.(*gcp.CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.CloudControllerManagerConfig)(nil), (*CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(a.(*gcp.CloudControllerManagerConfig), b.(*CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudNAT)(nil), (*gcp.CloudNAT)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudNAT_To_gcp_CloudNAT(a.(*CloudNAT), b.(*gcp.CloudNAT), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.CloudNAT)(nil), (*CloudNAT)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_CloudNAT_To_v1alpha1_CloudNAT(a.(*gcp.CloudNAT), b.(*CloudNAT), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudProfileConfig)(nil), (*gcp.CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudProfileConfig_To_gcp_CloudProfileConfig(a.(*CloudProfileConfig), b.(*gcp.CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.CloudProfileConfig)(nil), (*CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(a.(*gcp.CloudProfileConfig), b.(*CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudRouter)(nil), (*gcp.CloudRouter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudRouter_To_gcp_CloudRouter(a.(*CloudRouter), b.(*gcp.CloudRouter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.CloudRouter)(nil), (*CloudRouter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_CloudRouter_To_v1alpha1_CloudRouter(a.(*gcp.CloudRouter), b.(*CloudRouter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneConfig)(nil), (*gcp.ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneConfig_To_gcp_ControlPlaneConfig(a.(*ControlPlaneConfig), b.(*gcp.ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.ControlPlaneConfig)(nil), (*ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(a.(*gcp.ControlPlaneConfig), b.(*ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*gcp.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_gcp_InfrastructureConfig(a.(*InfrastructureConfig), b.(*gcp.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*gcp.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*gcp.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_gcp_InfrastructureStatus(a.(*InfrastructureStatus), b.(*gcp.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*gcp.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImage)(nil), (*gcp.MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImage_To_gcp_MachineImage(a.(*MachineImage), b.(*gcp.MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.MachineImage)(nil), (*MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_MachineImage_To_v1alpha1_MachineImage(a.(*gcp.MachineImage), b.(*MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImageVersion)(nil), (*gcp.MachineImageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImageVersion_To_gcp_MachineImageVersion(a.(*MachineImageVersion), b.(*gcp.MachineImageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.MachineImageVersion)(nil), (*MachineImageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_MachineImageVersion_To_v1alpha1_MachineImageVersion(a.(*gcp.MachineImageVersion), b.(*MachineImageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImages)(nil), (*gcp.MachineImages)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImages_To_gcp_MachineImages(a.(*MachineImages), b.(*gcp.MachineImages), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.MachineImages)(nil), (*MachineImages)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_MachineImages_To_v1alpha1_MachineImages(a.(*gcp.MachineImages), b.(*MachineImages), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkConfig)(nil), (*gcp.NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkConfig_To_gcp_NetworkConfig(a.(*NetworkConfig), b.(*gcp.NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.NetworkConfig)(nil), (*NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_NetworkConfig_To_v1alpha1_NetworkConfig(a.(*gcp.NetworkConfig), b.(*NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkStatus)(nil), (*gcp.NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkStatus_To_gcp_NetworkStatus(a.(*NetworkStatus), b.(*gcp.NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.NetworkStatus)(nil), (*NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_NetworkStatus_To_v1alpha1_NetworkStatus(a.(*gcp.NetworkStatus), b.(*NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Subnet)(nil), (*gcp.Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Subnet_To_gcp_Subnet(a.(*Subnet), b.(*gcp.Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.Subnet)(nil), (*Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_Subnet_To_v1alpha1_Subnet(a.(*gcp.Subnet), b.(*Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VPC)(nil), (*gcp.VPC)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VPC_To_gcp_VPC(a.(*VPC), b.(*gcp.VPC), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.VPC)(nil), (*VPC)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_VPC_To_v1alpha1_VPC(a.(*gcp.VPC), b.(*VPC), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkerStatus)(nil), (*gcp.WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkerStatus_To_gcp_WorkerStatus(a.(*WorkerStatus), b.(*gcp.WorkerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*gcp.WorkerStatus)(nil), (*WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_gcp_WorkerStatus_To_v1alpha1_WorkerStatus(a.(*gcp.WorkerStatus), b.(*WorkerStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloudControllerManagerConfig_To_gcp_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *gcp.CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_CloudControllerManagerConfig_To_gcp_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudControllerManagerConfig_To_gcp_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *gcp.CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudControllerManagerConfig_To_gcp_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_gcp_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *gcp.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_gcp_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_gcp_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *gcp.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_gcp_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_v1alpha1_CloudNAT_To_gcp_CloudNAT(in *CloudNAT, out *gcp.CloudNAT, s conversion.Scope) error {
	out.MinPortsPerVM = (*int32)(unsafe.Pointer(in.MinPortsPerVM))
	return nil
}

// Convert_v1alpha1_CloudNAT_To_gcp_CloudNAT is an autogenerated conversion function.
func Convert_v1alpha1_CloudNAT_To_gcp_CloudNAT(in *CloudNAT, out *gcp.CloudNAT, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudNAT_To_gcp_CloudNAT(in, out, s)
}

func autoConvert_gcp_CloudNAT_To_v1alpha1_CloudNAT(in *gcp.CloudNAT, out *CloudNAT, s conversion.Scope) error {
	out.MinPortsPerVM = (*int32)(unsafe.Pointer(in.MinPortsPerVM))
	return nil
}

// Convert_gcp_CloudNAT_To_v1alpha1_CloudNAT is an autogenerated conversion function.
func Convert_gcp_CloudNAT_To_v1alpha1_CloudNAT(in *gcp.CloudNAT, out *CloudNAT, s conversion.Scope) error {
	return autoConvert_gcp_CloudNAT_To_v1alpha1_CloudNAT(in, out, s)
}

func autoConvert_v1alpha1_CloudProfileConfig_To_gcp_CloudProfileConfig(in *CloudProfileConfig, out *gcp.CloudProfileConfig, s conversion.Scope) error {
	out.MachineImages = *(*[]gcp.MachineImages)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_v1alpha1_CloudProfileConfig_To_gcp_CloudProfileConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudProfileConfig_To_gcp_CloudProfileConfig(in *CloudProfileConfig, out *gcp.CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudProfileConfig_To_gcp_CloudProfileConfig(in, out, s)
}

func autoConvert_gcp_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *gcp.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImages)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_gcp_CloudProfileConfig_To_v1alpha1_CloudProfileConfig is an autogenerated conversion function.
func Convert_gcp_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *gcp.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_gcp_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in, out, s)
}

func autoConvert_v1alpha1_CloudRouter_To_gcp_CloudRouter(in *CloudRouter, out *gcp.CloudRouter, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_CloudRouter_To_gcp_CloudRouter is an autogenerated conversion function.
func Convert_v1alpha1_CloudRouter_To_gcp_CloudRouter(in *CloudRouter, out *gcp.CloudRouter, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudRouter_To_gcp_CloudRouter(in, out, s)
}

func autoConvert_gcp_CloudRouter_To_v1alpha1_CloudRouter(in *gcp.CloudRouter, out *CloudRouter, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_gcp_CloudRouter_To_v1alpha1_CloudRouter is an autogenerated conversion function.
func Convert_gcp_CloudRouter_To_v1alpha1_CloudRouter(in *gcp.CloudRouter, out *CloudRouter, s conversion.Scope) error {
	return autoConvert_gcp_CloudRouter_To_v1alpha1_CloudRouter(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneConfig_To_gcp_ControlPlaneConfig(in *ControlPlaneConfig, out *gcp.ControlPlaneConfig, s conversion.Scope) error {
	out.Zone = in.Zone
	out.CloudControllerManager = (*gcp.CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_v1alpha1_ControlPlaneConfig_To_gcp_ControlPlaneConfig is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneConfig_To_gcp_ControlPlaneConfig(in *ControlPlaneConfig, out *gcp.ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneConfig_To_gcp_ControlPlaneConfig(in, out, s)
}

func autoConvert_gcp_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *gcp.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	out.Zone = in.Zone
	out.CloudControllerManager = (*CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_gcp_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig is an autogenerated conversion function.
func Convert_gcp_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *gcp.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_gcp_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureConfig_To_gcp_InfrastructureConfig(in *InfrastructureConfig, out *gcp.InfrastructureConfig, s conversion.Scope) error {
	if err := Convert_v1alpha1_NetworkConfig_To_gcp_NetworkConfig(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_gcp_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_gcp_InfrastructureConfig(in *InfrastructureConfig, out *gcp.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_gcp_InfrastructureConfig(in, out, s)
}

func autoConvert_gcp_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *gcp.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	if err := Convert_gcp_NetworkConfig_To_v1alpha1_NetworkConfig(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_gcp_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_gcp_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *gcp.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_gcp_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_gcp_InfrastructureStatus(in *InfrastructureStatus, out *gcp.InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_NetworkStatus_To_gcp_NetworkStatus(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	out.ServiceAccountEmail = in.ServiceAccountEmail
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_gcp_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_gcp_InfrastructureStatus(in *InfrastructureStatus, out *gcp.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_gcp_InfrastructureStatus(in, out, s)
}

func autoConvert_gcp_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *gcp.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_gcp_NetworkStatus_To_v1alpha1_NetworkStatus(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	out.ServiceAccountEmail = in.ServiceAccountEmail
	return nil
}

// Convert_gcp_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_gcp_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *gcp.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_gcp_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_MachineImage_To_gcp_MachineImage(in *MachineImage, out *gcp.MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_v1alpha1_MachineImage_To_gcp_MachineImage is an autogenerated conversion function.
func Convert_v1alpha1_MachineImage_To_gcp_MachineImage(in *MachineImage, out *gcp.MachineImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImage_To_gcp_MachineImage(in, out, s)
}

func autoConvert_gcp_MachineImage_To_v1alpha1_MachineImage(in *gcp.MachineImage, out *MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_gcp_MachineImage_To_v1alpha1_MachineImage is an autogenerated conversion function.
func Convert_gcp_MachineImage_To_v1alpha1_MachineImage(in *gcp.MachineImage, out *MachineImage, s conversion.Scope) error {
	return autoConvert_gcp_MachineImage_To_v1alpha1_MachineImage(in, out, s)
}

func autoConvert_v1alpha1_MachineImageVersion_To_gcp_MachineImageVersion(in *MachineImageVersion, out *gcp.MachineImageVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_v1alpha1_MachineImageVersion_To_gcp_MachineImageVersion is an autogenerated conversion function.
func Convert_v1alpha1_MachineImageVersion_To_gcp_MachineImageVersion(in *MachineImageVersion, out *gcp.MachineImageVersion, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImageVersion_To_gcp_MachineImageVersion(in, out, s)
}

func autoConvert_gcp_MachineImageVersion_To_v1alpha1_MachineImageVersion(in *gcp.MachineImageVersion, out *MachineImageVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_gcp_MachineImageVersion_To_v1alpha1_MachineImageVersion is an autogenerated conversion function.
func Convert_gcp_MachineImageVersion_To_v1alpha1_MachineImageVersion(in *gcp.MachineImageVersion, out *MachineImageVersion, s conversion.Scope) error {
	return autoConvert_gcp_MachineImageVersion_To_v1alpha1_MachineImageVersion(in, out, s)
}

func autoConvert_v1alpha1_MachineImages_To_gcp_MachineImages(in *MachineImages, out *gcp.MachineImages, s conversion.Scope) error {
	out.Name = in.Name
	out.Versions = *(*[]gcp.MachineImageVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_v1alpha1_MachineImages_To_gcp_MachineImages is an autogenerated conversion function.
func Convert_v1alpha1_MachineImages_To_gcp_MachineImages(in *MachineImages, out *gcp.MachineImages, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImages_To_gcp_MachineImages(in, out, s)
}

func autoConvert_gcp_MachineImages_To_v1alpha1_MachineImages(in *gcp.MachineImages, out *MachineImages, s conversion.Scope) error {
	out.Name = in.Name
	out.Versions = *(*[]MachineImageVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_gcp_MachineImages_To_v1alpha1_MachineImages is an autogenerated conversion function.
func Convert_gcp_MachineImages_To_v1alpha1_MachineImages(in *gcp.MachineImages, out *MachineImages, s conversion.Scope) error {
	return autoConvert_gcp_MachineImages_To_v1alpha1_MachineImages(in, out, s)
}

func autoConvert_v1alpha1_NetworkConfig_To_gcp_NetworkConfig(in *NetworkConfig, out *gcp.NetworkConfig, s conversion.Scope) error {
	out.VPC = (*gcp.VPC)(unsafe.Pointer(in.VPC))
	out.CloudNAT = (*gcp.CloudNAT)(unsafe.Pointer(in.CloudNAT))
	out.Internal = (*string)(unsafe.Pointer(in.Internal))
	out.Worker = in.Worker
	return nil
}

// Convert_v1alpha1_NetworkConfig_To_gcp_NetworkConfig is an autogenerated conversion function.
func Convert_v1alpha1_NetworkConfig_To_gcp_NetworkConfig(in *NetworkConfig, out *gcp.NetworkConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkConfig_To_gcp_NetworkConfig(in, out, s)
}

func autoConvert_gcp_NetworkConfig_To_v1alpha1_NetworkConfig(in *gcp.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	out.VPC = (*VPC)(unsafe.Pointer(in.VPC))
	out.CloudNAT = (*CloudNAT)(unsafe.Pointer(in.CloudNAT))
	out.Internal = (*string)(unsafe.Pointer(in.Internal))
	out.Worker = in.Worker
	return nil
}

// Convert_gcp_NetworkConfig_To_v1alpha1_NetworkConfig is an autogenerated conversion function.
func Convert_gcp_NetworkConfig_To_v1alpha1_NetworkConfig(in *gcp.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	return autoConvert_gcp_NetworkConfig_To_v1alpha1_NetworkConfig(in, out, s)
}

func autoConvert_v1alpha1_NetworkStatus_To_gcp_NetworkStatus(in *NetworkStatus, out *gcp.NetworkStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_VPC_To_gcp_VPC(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	out.Subnets = *(*[]gcp.Subnet)(unsafe.Pointer(&in.Subnets))
	return nil
}

// Convert_v1alpha1_NetworkStatus_To_gcp_NetworkStatus is an autogenerated conversion function.
func Convert_v1alpha1_NetworkStatus_To_gcp_NetworkStatus(in *NetworkStatus, out *gcp.NetworkStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkStatus_To_gcp_NetworkStatus(in, out, s)
}

func autoConvert_gcp_NetworkStatus_To_v1alpha1_NetworkStatus(in *gcp.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	if err := Convert_gcp_VPC_To_v1alpha1_VPC(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	out.Subnets = *(*[]Subnet)(unsafe.Pointer(&in.Subnets))
	return nil
}

// Convert_gcp_NetworkStatus_To_v1alpha1_NetworkStatus is an autogenerated conversion function.
func Convert_gcp_NetworkStatus_To_v1alpha1_NetworkStatus(in *gcp.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return autoConvert_gcp_NetworkStatus_To_v1alpha1_NetworkStatus(in, out, s)
}

func autoConvert_v1alpha1_Subnet_To_gcp_Subnet(in *Subnet, out *gcp.Subnet, s conversion.Scope) error {
	out.Name = in.Name
	out.Purpose = gcp.SubnetPurpose(in.Purpose)
	return nil
}

// Convert_v1alpha1_Subnet_To_gcp_Subnet is an autogenerated conversion function.
func Convert_v1alpha1_Subnet_To_gcp_Subnet(in *Subnet, out *gcp.Subnet, s conversion.Scope) error {
	return autoConvert_v1alpha1_Subnet_To_gcp_Subnet(in, out, s)
}

func autoConvert_gcp_Subnet_To_v1alpha1_Subnet(in *gcp.Subnet, out *Subnet, s conversion.Scope) error {
	out.Name = in.Name
	out.Purpose = SubnetPurpose(in.Purpose)
	return nil
}

// Convert_gcp_Subnet_To_v1alpha1_Subnet is an autogenerated conversion function.
func Convert_gcp_Subnet_To_v1alpha1_Subnet(in *gcp.Subnet, out *Subnet, s conversion.Scope) error {
	return autoConvert_gcp_Subnet_To_v1alpha1_Subnet(in, out, s)
}

func autoConvert_v1alpha1_VPC_To_gcp_VPC(in *VPC, out *gcp.VPC, s conversion.Scope) error {
	out.Name = in.Name
	out.CloudRouter = (*gcp.CloudRouter)(unsafe.Pointer(in.CloudRouter))
	return nil
}

// Convert_v1alpha1_VPC_To_gcp_VPC is an autogenerated conversion function.
func Convert_v1alpha1_VPC_To_gcp_VPC(in *VPC, out *gcp.VPC, s conversion.Scope) error {
	return autoConvert_v1alpha1_VPC_To_gcp_VPC(in, out, s)
}

func autoConvert_gcp_VPC_To_v1alpha1_VPC(in *gcp.VPC, out *VPC, s conversion.Scope) error {
	out.Name = in.Name
	out.CloudRouter = (*CloudRouter)(unsafe.Pointer(in.CloudRouter))
	return nil
}

// Convert_gcp_VPC_To_v1alpha1_VPC is an autogenerated conversion function.
func Convert_gcp_VPC_To_v1alpha1_VPC(in *gcp.VPC, out *VPC, s conversion.Scope) error {
	return autoConvert_gcp_VPC_To_v1alpha1_VPC(in, out, s)
}

func autoConvert_v1alpha1_WorkerStatus_To_gcp_WorkerStatus(in *WorkerStatus, out *gcp.WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]gcp.MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_v1alpha1_WorkerStatus_To_gcp_WorkerStatus is an autogenerated conversion function.
func Convert_v1alpha1_WorkerStatus_To_gcp_WorkerStatus(in *WorkerStatus, out *gcp.WorkerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkerStatus_To_gcp_WorkerStatus(in, out, s)
}

func autoConvert_gcp_WorkerStatus_To_v1alpha1_WorkerStatus(in *gcp.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_gcp_WorkerStatus_To_v1alpha1_WorkerStatus is an autogenerated conversion function.
func Convert_gcp_WorkerStatus_To_v1alpha1_WorkerStatus(in *gcp.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	return autoConvert_gcp_WorkerStatus_To_v1alpha1_WorkerStatus(in, out, s)
}
