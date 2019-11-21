/*
Copyright 2016 The Kubernetes Authors.

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

package fake

import (
	authorizationapi "sigs.k8s.io/kustomize/pseudo/k8s/api/authorization/v1beta1"
	core "sigs.k8s.io/kustomize/pseudo/k8s/client-go/testing"
)

func (c *FakeSubjectAccessReviews) Create(sar *authorizationapi.SubjectAccessReview) (result *authorizationapi.SubjectAccessReview, err error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(authorizationapi.SchemeGroupVersion.WithResource("subjectaccessreviews"), sar), &authorizationapi.SubjectAccessReview{})
	return obj.(*authorizationapi.SubjectAccessReview), err
}