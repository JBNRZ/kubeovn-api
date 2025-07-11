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

// fakeIPs implements IPInterface
type fakeIPs struct {
	*gentype.FakeClientWithList[*v1.IP, *v1.IPList]
	Fake *FakeKubeovnV1
}

func newFakeIPs(fake *FakeKubeovnV1) kubeovnv1.IPInterface {
	return &fakeIPs{
		gentype.NewFakeClientWithList[*v1.IP, *v1.IPList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("ips"),
			v1.SchemeGroupVersion.WithKind("IP"),
			func() *v1.IP { return &v1.IP{} },
			func() *v1.IPList { return &v1.IPList{} },
			func(dst, src *v1.IPList) { dst.ListMeta = src.ListMeta },
			func(list *v1.IPList) []*v1.IP { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.IPList, items []*v1.IP) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
