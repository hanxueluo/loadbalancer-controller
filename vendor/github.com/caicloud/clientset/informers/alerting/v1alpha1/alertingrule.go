/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/alerting/v1alpha1"
	alertingv1alpha1 "github.com/caicloud/clientset/pkg/apis/alerting/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// AlertingRuleInformer provides access to a shared informer and lister for
// AlertingRules.
type AlertingRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AlertingRuleLister
}

type alertingRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAlertingRuleInformer constructs a new informer for AlertingRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAlertingRuleInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAlertingRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAlertingRuleInformer constructs a new informer for AlertingRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAlertingRuleInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertingV1alpha1().AlertingRules().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertingV1alpha1().AlertingRules().Watch(options)
			},
		},
		&alertingv1alpha1.AlertingRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *alertingRuleInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAlertingRuleInformer(client.(kubernetes.Interface), resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *alertingRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&alertingv1alpha1.AlertingRule{}, f.defaultInformer)
}

func (f *alertingRuleInformer) Lister() v1alpha1.AlertingRuleLister {
	return v1alpha1.NewAlertingRuleLister(f.Informer().GetIndexer())
}
