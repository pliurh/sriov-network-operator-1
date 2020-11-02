// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	operatorv1 "github.com/openshift/api/operator/v1"
	versioned "github.com/openshift/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/operator/listers/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceCatalogControllerManagerInformer provides access to a shared informer and lister for
// ServiceCatalogControllerManagers.
type ServiceCatalogControllerManagerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceCatalogControllerManagerLister
}

type serviceCatalogControllerManagerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewServiceCatalogControllerManagerInformer constructs a new informer for ServiceCatalogControllerManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceCatalogControllerManagerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceCatalogControllerManagerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredServiceCatalogControllerManagerInformer constructs a new informer for ServiceCatalogControllerManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceCatalogControllerManagerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().ServiceCatalogControllerManagers().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().ServiceCatalogControllerManagers().Watch(context.TODO(), options)
			},
		},
		&operatorv1.ServiceCatalogControllerManager{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceCatalogControllerManagerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceCatalogControllerManagerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceCatalogControllerManagerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1.ServiceCatalogControllerManager{}, f.defaultInformer)
}

func (f *serviceCatalogControllerManagerInformer) Lister() v1.ServiceCatalogControllerManagerLister {
	return v1.NewServiceCatalogControllerManagerLister(f.Informer().GetIndexer())
}
