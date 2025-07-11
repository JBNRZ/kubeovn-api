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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apiskubeovnv1 "github.com/JBNRZ/kubeovn-api/pkg/apis/kubeovn/v1"
	versioned "github.com/JBNRZ/kubeovn-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/JBNRZ/kubeovn-api/pkg/client/informers/externalversions/internalinterfaces"
	kubeovnv1 "github.com/JBNRZ/kubeovn-api/pkg/client/listers/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VlanInformer provides access to a shared informer and lister for
// Vlans.
type VlanInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeovnv1.VlanLister
}

type vlanInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVlanInformer constructs a new informer for Vlan type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVlanInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVlanInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVlanInformer constructs a new informer for Vlan type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVlanInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().Vlans().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().Vlans().Watch(context.TODO(), options)
			},
		},
		&apiskubeovnv1.Vlan{},
		resyncPeriod,
		indexers,
	)
}

func (f *vlanInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVlanInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vlanInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeovnv1.Vlan{}, f.defaultInformer)
}

func (f *vlanInformer) Lister() kubeovnv1.VlanLister {
	return kubeovnv1.NewVlanLister(f.Informer().GetIndexer())
}
