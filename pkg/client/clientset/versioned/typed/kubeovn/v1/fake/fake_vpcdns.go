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

// fakeVpcDnses implements VpcDnsInterface
type fakeVpcDnses struct {
	*gentype.FakeClientWithList[*v1.VpcDns, *v1.VpcDnsList]
	Fake *FakeKubeovnV1
}

func newFakeVpcDnses(fake *FakeKubeovnV1) kubeovnv1.VpcDnsInterface {
	return &fakeVpcDnses{
		gentype.NewFakeClientWithList[*v1.VpcDns, *v1.VpcDnsList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("vpc-dnses"),
			v1.SchemeGroupVersion.WithKind("VpcDns"),
			func() *v1.VpcDns { return &v1.VpcDns{} },
			func() *v1.VpcDnsList { return &v1.VpcDnsList{} },
			func(dst, src *v1.VpcDnsList) { dst.ListMeta = src.ListMeta },
			func(list *v1.VpcDnsList) []*v1.VpcDns { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.VpcDnsList, items []*v1.VpcDns) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
