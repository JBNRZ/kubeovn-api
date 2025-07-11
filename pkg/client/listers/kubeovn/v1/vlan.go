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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/JBNRZ/kubeovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// VlanLister helps list Vlans.
// All objects returned here must be treated as read-only.
type VlanLister interface {
	// List lists all Vlans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.Vlan, err error)
	// Get retrieves the Vlan from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.Vlan, error)
	VlanListerExpansion
}

// vlanLister implements the VlanLister interface.
type vlanLister struct {
	listers.ResourceIndexer[*kubeovnv1.Vlan]
}

// NewVlanLister returns a new VlanLister.
func NewVlanLister(indexer cache.Indexer) VlanLister {
	return &vlanLister{listers.New[*kubeovnv1.Vlan](indexer, kubeovnv1.Resource("vlan"))}
}
