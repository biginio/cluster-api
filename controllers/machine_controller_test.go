/*
Copyright 2019 The Kubernetes Authors.

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

package controllers

import (
	"testing" //testing library 불러오기

	. "github.com/onsi/ginkgo" //
	. "github.com/onsi/gomega" //
	corev1 "k8s.io/api/core/v1"  // kubernetes api 핵심 라이브러리
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // kubernetes api
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured" // kubernetes api
	"k8s.io/apimachinery/pkg/types" // kubernetes api
	"k8s.io/client-go/kubernetes/scheme" // kubernetes
	"k8s.io/utils/pointer" //
	"sigs.k8s.io/cluster-api/api/v1alpha2" // kubernetes
	"sigs.k8s.io/controller-runtime/pkg/client/fake" // kubernetes
	"sigs.k8s.io/controller-runtime/pkg/log" // kubernetes
	"sigs.k8s.io/controller-runtime/pkg/reconcile" // kubernetes
)

var _ = Describe("Machine Reconciler", func() {
	It("Should create a Machine", func() {
		// TODO
	})
})

func TestReconcileRequest(t *testing.T) {
	RegisterTestingT(t)

	// Infrastructure configuration 설정하기

	infraConfig := unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "InfrastructureConfig",
			"apiVersion": "infrastructure.cluster.x-k8s.io/v1alpha2",
			"metadata": map[string]interface{}{
				"name":      "infra-config1",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"providerID": "test://id-1",
			},
			"status": map[string]interface{}{
				"ready": true,
				"addresses": []interface{}{
					map[string]interface{}{
						"type":    "InternalIP",
						"address": "10.0.0.1",
					},
				},
			},
		},
	}
	machine1 := v1alpha2.Machine{ // 사용할 cluster 설정하기
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:       "create", // 생성하기
			Namespace:  "default",
			Finalizers: []string{v1alpha2.MachineFinalizer, metav1.FinalizerDeleteDependents},
		},
		Spec: v1alpha2.MachineSpec{
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha2",
				Kind:       "InfrastructureConfig",
				Name:       "infra-config1",
			},
			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")},
		},
	}
	machine2 := v1alpha2.Machine{ // 사용할 cluster 설정하기
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:       "update", // 변경 내용 업데이트하기
			Namespace:  "default",
			Finalizers: []string{v1alpha2.MachineFinalizer, metav1.FinalizerDeleteDependents},
		},
		Spec: v1alpha2.MachineSpec{
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha2",
				Kind:       "InfrastructureConfig",
				Name:       "infra-config1",
			},
			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")},
		},
	}
	time := metav1.Now()
	machine3 := v1alpha2.Machine{ // 사용할 cluster 설정하기
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              "delete", // 삭제하기
			Namespace:         "default",
			Finalizers:        []string{v1alpha2.MachineFinalizer, metav1.FinalizerDeleteDependents},
			DeletionTimestamp: &time,
		},
		Spec: v1alpha2.MachineSpec{
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha2",
				Kind:       "InfrastructureConfig",
				Name:       "infra-config1",
			},
			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")},
		},
	}
	clusterList := v1alpha2.ClusterList{ // cluster 목록
		TypeMeta: metav1.TypeMeta{
			Kind: "ClusterList",
		},
		Items: []v1alpha2.Cluster{
			{
				TypeMeta: metav1.TypeMeta{
					Kind: "Cluster",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testcluster",
					Namespace: "default",
				},
			},
			{
				TypeMeta: metav1.TypeMeta{
					Kind: "Cluster",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "rainbow",
					Namespace: "foo",
				},
			},
		},
	}

	type expected struct { //
		result reconcile.Result
		err    bool
	}
	testCases := []struct { // controller에 원하는 상태로 실행하도록 명령을 보내주는 내용을 저장하는 것
		request     reconcile.Request
		existsValue bool
		expected    expected
	}{
		{
			request:     reconcile.Request{NamespacedName: types.NamespacedName{Name: machine1.Name, Namespace: machine1.Namespace}},
			existsValue: false,
			expected: expected{
				result: reconcile.Result{},
				err:    false,
			}, // 생성하는 경우
		},
		{
			request:     reconcile.Request{NamespacedName: types.NamespacedName{Name: machine2.Name, Namespace: machine2.Namespace}},
			existsValue: true,
			expected: expected{
				result: reconcile.Result{},
				err:    false,
			}, // 변경 사항 업데이트하는 경우
		},
		{
			request:     reconcile.Request{NamespacedName: types.NamespacedName{Name: machine3.Name, Namespace: machine3.Namespace}},
			existsValue: true,
			expected: expected{
				result: reconcile.Result{},
				err:    false,
			}, // 삭제하는 경우
		},
	}

	for _, tc := range testCases {
		v1alpha2.AddToScheme(scheme.Scheme)
		r := &MachineReconciler{
			Client: fake.NewFakeClient(&clusterList, &machine1, &machine2, &machine3, &infraConfig),
			Log:    log.Log,
		}

		result, err := r.Reconcile(tc.request)
		if tc.expected.err {
			Expect(err).ToNot(BeNil())
		} else {
			Expect(err).To(BeNil())
		}

		Expect(result).To(Equal(tc.expected.result))
	}
}
