/*
Copyright 2021.

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
	"context"
	time "time"

	webappv1 "example.com/foo-controller/apis/webapp/v1"
	versioned "example.com/foo-controller/generated/webapp/clientset/versioned"
	internalinterfaces "example.com/foo-controller/generated/webapp/informers/externalversions/internalinterfaces"
	v1 "example.com/foo-controller/generated/webapp/listers/webapp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GuestbookInformer provides access to a shared informer and lister for
// Guestbooks.
type GuestbookInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.GuestbookLister
}

type guestbookInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGuestbookInformer constructs a new informer for Guestbook type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGuestbookInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGuestbookInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGuestbookInformer constructs a new informer for Guestbook type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGuestbookInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WebappV1().Guestbooks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WebappV1().Guestbooks(namespace).Watch(context.TODO(), options)
			},
		},
		&webappv1.Guestbook{},
		resyncPeriod,
		indexers,
	)
}

func (f *guestbookInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGuestbookInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *guestbookInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&webappv1.Guestbook{}, f.defaultInformer)
}

func (f *guestbookInformer) Lister() v1.GuestbookLister {
	return v1.NewGuestbookLister(f.Informer().GetIndexer())
}
