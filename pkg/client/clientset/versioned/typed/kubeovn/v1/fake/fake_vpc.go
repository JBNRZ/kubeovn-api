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

// fakeVpcs implements VpcInterface
type fakeVpcs struct {
	*gentype.FakeClientWithList[*v1.Vpc, *v1.VpcList]
	Fake *FakeKubeovnV1
}

func newFakeVpcs(fake *FakeKubeovnV1) kubeovnv1.VpcInterface {
	return &fakeVpcs{
		gentype.NewFakeClientWithList[*v1.Vpc, *v1.VpcList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("vpcs"),
			v1.SchemeGroupVersion.WithKind("Vpc"),
			func() *v1.Vpc { return &v1.Vpc{} },
			func() *v1.VpcList { return &v1.VpcList{} },
			func(dst, src *v1.VpcList) { dst.ListMeta = src.ListMeta },
			func(list *v1.VpcList) []*v1.Vpc { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.VpcList, items []*v1.Vpc) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
