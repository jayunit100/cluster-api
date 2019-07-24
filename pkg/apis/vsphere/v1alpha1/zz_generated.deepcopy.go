// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyPair) DeepCopyInto(out *KeyPair) {
	*out = *in
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyPair.
func (in *KeyPair) DeepCopy() *KeyPair {
	if in == nil {
		return nil
	}
	out := new(KeyPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfiguration) DeepCopyInto(out *KubeadmConfiguration) {
	*out = *in
	in.Join.DeepCopyInto(&out.Join)
	in.Init.DeepCopyInto(&out.Init)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfiguration.
func (in *KubeadmConfiguration) DeepCopy() *KubeadmConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkDeviceSpec) DeepCopyInto(out *NetworkDeviceSpec) {
	*out = *in
	if in.IPAddrs != nil {
		in, out := &in.IPAddrs, &out.IPAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MTU != nil {
		in, out := &in.MTU, &out.MTU
		*out = new(int64)
		**out = **in
	}
	if in.Nameservers != nil {
		in, out := &in.Nameservers, &out.Nameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]NetworkRouteSpec, len(*in))
		copy(*out, *in)
	}
	if in.SearchDomains != nil {
		in, out := &in.SearchDomains, &out.SearchDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkDeviceSpec.
func (in *NetworkDeviceSpec) DeepCopy() *NetworkDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRouteSpec) DeepCopyInto(out *NetworkRouteSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRouteSpec.
func (in *NetworkRouteSpec) DeepCopy() *NetworkRouteSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]NetworkDeviceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]NetworkRouteSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.IPAddrs != nil {
		in, out := &in.IPAddrs, &out.IPAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereClusterProviderSpec) DeepCopyInto(out *VsphereClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.SSHAuthorizedKeys != nil {
		in, out := &in.SSHAuthorizedKeys, &out.SSHAuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.CAKeyPair.DeepCopyInto(&out.CAKeyPair)
	in.EtcdCAKeyPair.DeepCopyInto(&out.EtcdCAKeyPair)
	in.FrontProxyCAKeyPair.DeepCopyInto(&out.FrontProxyCAKeyPair)
	in.SAKeyPair.DeepCopyInto(&out.SAKeyPair)
	in.ClusterConfiguration.DeepCopyInto(&out.ClusterConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereClusterProviderSpec.
func (in *VsphereClusterProviderSpec) DeepCopy() *VsphereClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(VsphereClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereClusterProviderStatus) DeepCopyInto(out *VsphereClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereClusterProviderStatus.
func (in *VsphereClusterProviderStatus) DeepCopy() *VsphereClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(VsphereClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereMachineProviderSpec) DeepCopyInto(out *VsphereMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Network.DeepCopyInto(&out.Network)
	if in.TrustedCerts != nil {
		in, out := &in.TrustedCerts, &out.TrustedCerts
		*out = make([][]byte, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.NTPServers != nil {
		in, out := &in.NTPServers, &out.NTPServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.KubeadmConfiguration.DeepCopyInto(&out.KubeadmConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereMachineProviderSpec.
func (in *VsphereMachineProviderSpec) DeepCopy() *VsphereMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(VsphereMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereMachineProviderStatus) DeepCopyInto(out *VsphereMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereMachineProviderStatus.
func (in *VsphereMachineProviderStatus) DeepCopy() *VsphereMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(VsphereMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}