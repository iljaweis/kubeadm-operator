// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Command) DeepCopyInto(out *Command) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Command.
func (in *Command) DeepCopy() *Command {
	if in == nil {
		return nil
	}
	out := new(Command)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Command) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandList) DeepCopyInto(out *CommandList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Command, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandList.
func (in *CommandList) DeepCopy() *CommandList {
	if in == nil {
		return nil
	}
	out := new(CommandList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommandList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandSpec) DeepCopyInto(out *CommandSpec) {
	*out = *in
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = new(Requires)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Require, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandSpec.
func (in *CommandSpec) DeepCopy() *CommandSpec {
	if in == nil {
		return nil
	}
	out := new(CommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandStatus) DeepCopyInto(out *CommandStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandStatus.
func (in *CommandStatus) DeepCopy() *CommandStatus {
	if in == nil {
		return nil
	}
	out := new(CommandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *File) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContent) DeepCopyInto(out *FileContent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContent.
func (in *FileContent) DeepCopy() *FileContent {
	if in == nil {
		return nil
	}
	out := new(FileContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileContent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContentList) DeepCopyInto(out *FileContentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContentList.
func (in *FileContentList) DeepCopy() *FileContentList {
	if in == nil {
		return nil
	}
	out := new(FileContentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileContentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContentSpec) DeepCopyInto(out *FileContentSpec) {
	*out = *in
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = new(Requires)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Require, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContentSpec.
func (in *FileContentSpec) DeepCopy() *FileContentSpec {
	if in == nil {
		return nil
	}
	out := new(FileContentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContentStatus) DeepCopyInto(out *FileContentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContentStatus.
func (in *FileContentStatus) DeepCopy() *FileContentStatus {
	if in == nil {
		return nil
	}
	out := new(FileContentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileList) DeepCopyInto(out *FileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileList.
func (in *FileList) DeepCopy() *FileList {
	if in == nil {
		return nil
	}
	out := new(FileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSource) DeepCopyInto(out *FileSource) {
	*out = *in
	if in.FileContent != nil {
		in, out := &in.FileContent, &out.FileContent
		*out = new(FileSourceFileContent)
		**out = **in
	}
	if in.CommandOutput != nil {
		in, out := &in.CommandOutput, &out.CommandOutput
		*out = new(FileSourceCommandOutput)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(FileSourceConfigMap)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(FileSourceSecret)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSource.
func (in *FileSource) DeepCopy() *FileSource {
	if in == nil {
		return nil
	}
	out := new(FileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSourceCommandOutput) DeepCopyInto(out *FileSourceCommandOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSourceCommandOutput.
func (in *FileSourceCommandOutput) DeepCopy() *FileSourceCommandOutput {
	if in == nil {
		return nil
	}
	out := new(FileSourceCommandOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSourceConfigMap) DeepCopyInto(out *FileSourceConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSourceConfigMap.
func (in *FileSourceConfigMap) DeepCopy() *FileSourceConfigMap {
	if in == nil {
		return nil
	}
	out := new(FileSourceConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSourceFileContent) DeepCopyInto(out *FileSourceFileContent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSourceFileContent.
func (in *FileSourceFileContent) DeepCopy() *FileSourceFileContent {
	if in == nil {
		return nil
	}
	out := new(FileSourceFileContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSourceSecret) DeepCopyInto(out *FileSourceSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSourceSecret.
func (in *FileSourceSecret) DeepCopy() *FileSourceSecret {
	if in == nil {
		return nil
	}
	out := new(FileSourceSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(FileSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = new(Requires)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Require, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStatus) DeepCopyInto(out *FileStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStatus.
func (in *FileStatus) DeepCopy() *FileStatus {
	if in == nil {
		return nil
	}
	out := new(FileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Host) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostList) DeepCopyInto(out *HostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Host, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostList.
func (in *HostList) DeepCopy() *HostList {
	if in == nil {
		return nil
	}
	out := new(HostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSpec) DeepCopyInto(out *HostSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSpec.
func (in *HostSpec) DeepCopy() *HostSpec {
	if in == nil {
		return nil
	}
	out := new(HostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	if in.Facts != nil {
		in, out := &in.Facts, &out.Facts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Require) DeepCopyInto(out *Require) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = new(RequireCommand)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(RequireFile)
		**out = **in
	}
	if in.FileContent != nil {
		in, out := &in.FileContent, &out.FileContent
		*out = new(RequireFileContent)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Require.
func (in *Require) DeepCopy() *Require {
	if in == nil {
		return nil
	}
	out := new(Require)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequireCommand) DeepCopyInto(out *RequireCommand) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequireCommand.
func (in *RequireCommand) DeepCopy() *RequireCommand {
	if in == nil {
		return nil
	}
	out := new(RequireCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequireFile) DeepCopyInto(out *RequireFile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequireFile.
func (in *RequireFile) DeepCopy() *RequireFile {
	if in == nil {
		return nil
	}
	out := new(RequireFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequireFileContent) DeepCopyInto(out *RequireFileContent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequireFileContent.
func (in *RequireFileContent) DeepCopy() *RequireFileContent {
	if in == nil {
		return nil
	}
	out := new(RequireFileContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Requires) DeepCopyInto(out *Requires) {
	{
		in := &in
		*out = make(Requires, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requires.
func (in Requires) DeepCopy() Requires {
	if in == nil {
		return nil
	}
	out := new(Requires)
	in.DeepCopyInto(out)
	return *out
}
