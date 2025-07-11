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

// fakeIptablesDnatRules implements IptablesDnatRuleInterface
type fakeIptablesDnatRules struct {
	*gentype.FakeClientWithList[*v1.IptablesDnatRule, *v1.IptablesDnatRuleList]
	Fake *FakeKubeovnV1
}

func newFakeIptablesDnatRules(fake *FakeKubeovnV1) kubeovnv1.IptablesDnatRuleInterface {
	return &fakeIptablesDnatRules{
		gentype.NewFakeClientWithList[*v1.IptablesDnatRule, *v1.IptablesDnatRuleList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("iptables-dnat-rules"),
			v1.SchemeGroupVersion.WithKind("IptablesDnatRule"),
			func() *v1.IptablesDnatRule { return &v1.IptablesDnatRule{} },
			func() *v1.IptablesDnatRuleList { return &v1.IptablesDnatRuleList{} },
			func(dst, src *v1.IptablesDnatRuleList) { dst.ListMeta = src.ListMeta },
			func(list *v1.IptablesDnatRuleList) []*v1.IptablesDnatRule { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.IptablesDnatRuleList, items []*v1.IptablesDnatRule) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
