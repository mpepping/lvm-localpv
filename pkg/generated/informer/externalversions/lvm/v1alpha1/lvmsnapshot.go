/*
Copyright 2021 The OpenEBS Authors

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

package v1alpha1

import (
	"context"
	time "time"

	lvmv1alpha1 "github.com/openebs/lvm-localpv/pkg/apis/openebs.io/lvm/v1alpha1"
	internalclientset "github.com/openebs/lvm-localpv/pkg/generated/clientset/internalclientset"
	internalinterfaces "github.com/openebs/lvm-localpv/pkg/generated/informer/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/lvm-localpv/pkg/generated/lister/lvm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LVMSnapshotInformer provides access to a shared informer and lister for
// LVMSnapshots.
type LVMSnapshotInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LVMSnapshotLister
}

type lVMSnapshotInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLVMSnapshotInformer constructs a new informer for LVMSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLVMSnapshotInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLVMSnapshotInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLVMSnapshotInformer constructs a new informer for LVMSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLVMSnapshotInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalV1alpha1().LVMSnapshots(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalV1alpha1().LVMSnapshots(namespace).Watch(context.TODO(), options)
			},
		},
		&lvmv1alpha1.LVMSnapshot{},
		resyncPeriod,
		indexers,
	)
}

func (f *lVMSnapshotInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLVMSnapshotInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *lVMSnapshotInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lvmv1alpha1.LVMSnapshot{}, f.defaultInformer)
}

func (f *lVMSnapshotInformer) Lister() v1alpha1.LVMSnapshotLister {
	return v1alpha1.NewLVMSnapshotLister(f.Informer().GetIndexer())
}
