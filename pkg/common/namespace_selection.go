// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package common

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//=================================================================
// GetSelectedNamespaces returns the list of filtered namespaces according to the policy namespace selector
func GetSelectedNamespaces(included, excluded, allNamespaces []string) []string {
	//get all namespaces
	//allNamespaces := getAllNamespaces() //TODO change this to call the func
	//then get the list of included
	includedNamespaces := []string{}
	for _, value := range included {
		found := FindPattern(value, allNamespaces)
		if found != nil {
			includedNamespaces = append(includedNamespaces, found...)
		}
	}
	//then get the list of excluded
	excludedNamespaces := []string{}
	for _, value := range excluded {
		found := FindPattern(value, allNamespaces)
		if found != nil {
			excludedNamespaces = append(excludedNamespaces, found...)
		}
	}
	//then get the list of deduplicated
	finalList := DeduplicateItems(includedNamespaces, excludedNamespaces)
	return finalList
}

//=================================================================
//GetAllNamespaces gets the list of all namespaces from k8s
func GetAllNamespaces() (list []string, err error) {
	namespaces := KubeClient.CoreV1().Namespaces()
	namespaceList, err := namespaces.List(metav1.ListOptions{})

	namespacesNames := []string{}
	for _, n := range namespaceList.Items {
		namespacesNames = append(namespacesNames, n.Name)
	}
	return namespacesNames, err
}
