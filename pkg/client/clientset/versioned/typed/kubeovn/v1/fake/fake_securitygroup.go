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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/JBNRZ/kubeovn-api/pkg/apis/kubeovn/v1"
	kubeovnv1 "github.com/JBNRZ/kubeovn-api/pkg/client/clientset/versioned/typed/kubeovn/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeSecurityGroups implements SecurityGroupInterface
type fakeSecurityGroups struct {
	*gentype.FakeClientWithList[*v1.SecurityGroup, *v1.SecurityGroupList]
	Fake *FakeKubeovnV1
}

func newFakeSecurityGroups(fake *FakeKubeovnV1) kubeovnv1.SecurityGroupInterface {
	return &fakeSecurityGroups{
		gentype.NewFakeClientWithList[*v1.SecurityGroup, *v1.SecurityGroupList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("security-groups"),
			v1.SchemeGroupVersion.WithKind("SecurityGroup"),
			func() *v1.SecurityGroup { return &v1.SecurityGroup{} },
			func() *v1.SecurityGroupList { return &v1.SecurityGroupList{} },
			func(dst, src *v1.SecurityGroupList) { dst.ListMeta = src.ListMeta },
			func(list *v1.SecurityGroupList) []*v1.SecurityGroup { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.SecurityGroupList, items []*v1.SecurityGroup) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
