// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=config.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("apiservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().APIServers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("authentications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Authentications().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("builds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Builds().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusteroperators"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().ClusterOperators().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusterversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().ClusterVersions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("consoles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Consoles().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("dnses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().DNSes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("featuregates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().FeatureGates().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("images"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Images().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("imagecontentpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().ImageContentPolicies().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("imagedigestmirrorsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().ImageDigestMirrorSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("imagetagmirrorsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().ImageTagMirrorSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("infrastructures"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Infrastructures().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Ingresses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("networks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Networks().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Nodes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("oauths"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().OAuths().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("operatorhubs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().OperatorHubs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Projects().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("proxies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Proxies().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("schedulers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1().Schedulers().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
