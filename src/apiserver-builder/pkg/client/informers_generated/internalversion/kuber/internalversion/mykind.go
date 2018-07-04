/*
 * licensed to vmware.
 */

// This file was automatically generated by informer-gen

package internalversion

import (
	kuber "apiserver-builder/pkg/apis/kuber"
	internalclientset "apiserver-builder/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "apiserver-builder/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "apiserver-builder/pkg/client/listers_generated/kuber/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// MyKindInformer provides access to a shared informer and lister for
// MyKinds.
type MyKindInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.MyKindLister
}

type myKindInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMyKindInformer constructs a new informer for MyKind type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMyKindInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMyKindInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMyKindInformer constructs a new informer for MyKind type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMyKindInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Kuber().MyKinds(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Kuber().MyKinds(namespace).Watch(options)
			},
		},
		&kuber.MyKind{},
		resyncPeriod,
		indexers,
	)
}

func (f *myKindInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMyKindInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *myKindInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kuber.MyKind{}, f.defaultInformer)
}

func (f *myKindInformer) Lister() internalversion.MyKindLister {
	return internalversion.NewMyKindLister(f.Informer().GetIndexer())
}