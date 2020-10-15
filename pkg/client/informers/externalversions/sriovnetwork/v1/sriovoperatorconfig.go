// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	sriovnetworkv1 "github.com/openshift/sriov-network-operator/api/v1"
	versioned "github.com/openshift/sriov-network-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift/sriov-network-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/sriov-network-operator/pkg/client/listers/sriovnetwork/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SriovOperatorConfigInformer provides access to a shared informer and lister for
// SriovOperatorConfigs.
type SriovOperatorConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SriovOperatorConfigLister
}

type sriovOperatorConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSriovOperatorConfigInformer constructs a new informer for SriovOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSriovOperatorConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSriovOperatorConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSriovOperatorConfigInformer constructs a new informer for SriovOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSriovOperatorConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SriovnetworkV1().SriovOperatorConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SriovnetworkV1().SriovOperatorConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&sriovnetworkv1.SriovOperatorConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *sriovOperatorConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSriovOperatorConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sriovOperatorConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sriovnetworkv1.SriovOperatorConfig{}, f.defaultInformer)
}

func (f *sriovOperatorConfigInformer) Lister() v1.SriovOperatorConfigLister {
	return v1.NewSriovOperatorConfigLister(f.Informer().GetIndexer())
}
