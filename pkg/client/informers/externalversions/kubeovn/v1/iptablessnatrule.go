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

// IptablesSnatRuleInformer provides access to a shared informer and lister for
// IptablesSnatRules.
type IptablesSnatRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeovnv1.IptablesSnatRuleLister
}

type iptablesSnatRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIptablesSnatRuleInformer constructs a new informer for IptablesSnatRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIptablesSnatRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIptablesSnatRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIptablesSnatRuleInformer constructs a new informer for IptablesSnatRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIptablesSnatRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().IptablesSnatRules().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().IptablesSnatRules().Watch(context.TODO(), options)
			},
		},
		&apiskubeovnv1.IptablesSnatRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *iptablesSnatRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIptablesSnatRuleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iptablesSnatRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeovnv1.IptablesSnatRule{}, f.defaultInformer)
}

func (f *iptablesSnatRuleInformer) Lister() kubeovnv1.IptablesSnatRuleLister {
	return kubeovnv1.NewIptablesSnatRuleLister(f.Informer().GetIndexer())
}
