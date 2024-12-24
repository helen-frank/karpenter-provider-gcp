/*
Copyright 2024 The CloudPilot AI Authors.

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

// Package apis contains Kubernetes API groups.
package apis

import (
	"sigs.k8s.io/karpenter/pkg/apis"
)

var (
	Group              = "karpenter.k8s.gcp"
	CompatibilityGroup = "compatibility." + Group
	CRDs               = apis.CRDs // object.Unmarshal[apiextensionsv1.CustomResourceDefinition](crds.GCENodeClassCRD)
)
