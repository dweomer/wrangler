/*
Copyright 2018 The Kubernetes Authors.

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

package lister

import (
	"github.com/rancher/wrangler/v3/pkg/summary"
	"k8s.io/apimachinery/pkg/labels"
)

// Lister helps list resources.
type Lister interface {
	// List lists all resources in the indexer.
	List(selector labels.Selector) (ret []*summary.SummarizedObject, err error)
	// Get retrieves a resource from the indexer with the given name
	Get(name string) (*summary.SummarizedObject, error)
	// Namespace returns an object that can list and get resources in a given namespace.
	Namespace(namespace string) NamespaceLister
}

// NamespaceLister helps list and get resources.
type NamespaceLister interface {
	// List lists all resources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*summary.SummarizedObject, err error)
	// Get retrieves a resource from the indexer for a given namespace and name.
	Get(name string) (*summary.SummarizedObject, error)
}
