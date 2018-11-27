// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	garden "github.com/gardener/gardener/pkg/apis/garden"
	clientsetinternalversion "github.com/gardener/gardener/pkg/client/garden/clientset/internalversion"
	internalinterfaces "github.com/gardener/gardener/pkg/client/garden/informers/internalversion/internalinterfaces"
	internalversion "github.com/gardener/gardener/pkg/client/garden/listers/garden/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProjectInformer provides access to a shared informer and lister for
// Projects.
type ProjectInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ProjectLister
}

type projectInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewProjectInformer constructs a new informer for Project type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProjectInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProjectInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredProjectInformer constructs a new informer for Project type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProjectInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Garden().Projects().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Garden().Projects().Watch(options)
			},
		},
		&garden.Project{},
		resyncPeriod,
		indexers,
	)
}

func (f *projectInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProjectInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *projectInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&garden.Project{}, f.defaultInformer)
}

func (f *projectInformer) Lister() internalversion.ProjectLister {
	return internalversion.NewProjectLister(f.Informer().GetIndexer())
}
