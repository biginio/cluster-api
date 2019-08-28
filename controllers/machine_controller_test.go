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


// 이 테스트는 machine_controller.go를 테스트하기 위한 테스트 코드 입니다.
import (
<<<<<<< HEAD
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
=======
	"testing" // go 가 제공 해주는 테스팅 도구

	. "github.com/onsi/ginkgo" // go 에서 BDD를 하게 해주는 테스팅 도구
	. "github.com/onsi/gomega" // go 에서 BDD하며 테스팅 결과를 매칭 할수 있게 해주는 operator 제공
	corev1 "k8s.io/api/core/v1" // kubernetes에 코어 API
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // kubernetes에 common 코어 API Types
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured" // kubernetes에 common 코어 API Types
	"k8s.io/apimachinery/pkg/types" // kubernetes에 common 코어 API Types
	"k8s.io/client-go/kubernetes/scheme" // kubernetes scheme
	"k8s.io/utils/pointer" // kubernetes custom pointer types
	"sigs.k8s.io/cluster-api/api/v1alpha2" // 테스트할 코드
	"sigs.k8s.io/controller-runtime/pkg/client/fake" // Kubernetes Mock Client 코드
	"sigs.k8s.io/controller-runtime/pkg/log" // Kubernetes Log 도구
	"sigs.k8s.io/controller-runtime/pkg/reconcile" // Kubernetes reconcile 헬퍼 도구 코드
>>>>>>> 5cfa408b258d71e95da626299c0c8f3cac7d4a61
)

var _ = Describe("Machine Reconciler", func() { // Machine Reconciler 를 테스팅 하기 위한 테스트 Set
	It("Should create a Machine", func() { // Machine Reconciler 테스트 item
		// TODO 해야할 업무 BDD 형식으로 변경할 예정인듯
	})
})

<<<<<<< HEAD

func TestReconcileRequest(t *testing.T) {
	RegisterTestingT(t) // gomega 에서 RegisterTestingT 라는 function을 사용 합니다. RegisterTestingT connects Gomega to Golang's XUnit style Testing


	//test 할 환경들을 설정합니다.
	infraConfig := unstructured.Unstructured{
		Object: map[string]interface{}{
=======
func TestReconcileRequest(t *testing.T) { // 테스트를 위한 function
	RegisterTestingT(t) // 테스트 등록

	infraConfig := unstructured.Unstructured{ // 테스트를 하기 위해 선언하는 소스 코드 선언문
		Object: map[string]interface{}{ // Machine을 테스트하기 위해 Machine이 소속되어야 할 Infra 설정 ----------------------
>>>>>>> 5cfa408b258d71e95da626299c0c8f3cac7d4a61
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
<<<<<<< HEAD
	}

	fmt.Println(infraConfig)

	machine1 := v1alpha2.Machine{
=======
	} // Machine을 테스트하기 위해 Machine이 소속되어야 할 Infra 설정 ----------------------
	machine1 := v1alpha2.Machine{ // Machine 1 설정
>>>>>>> 5cfa408b258d71e95da626299c0c8f3cac7d4a61
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine", // 설정 종류
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:       "create", // 현재 상태 등록
			Namespace:  "default", // 네임스페이스 등록
			Finalizers: []string{v1alpha2.MachineFinalizer, metav1.FinalizerDeleteDependents}, // Finalizer 입니다.
		},
		Spec: v1alpha2.MachineSpec{ // 소속될 infra
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha2",
				Kind:       "InfrastructureConfig",
				Name:       "infra-config1",
			},
<<<<<<< HEAD

			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")},
		},
	}


	machine2 := v1alpha2.Machine{
=======
			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")}, // 임의에 데이터
		},
	}
	machine2 := v1alpha2.Machine{ // Machine 2 등록
>>>>>>> 5cfa408b258d71e95da626299c0c8f3cac7d4a61
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine", // 설정 종류
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:       "update", // 현재 상태 등록
			Namespace:  "default", // 네임스페이스 등록
			Finalizers: []string{v1alpha2.MachineFinalizer, metav1.FinalizerDeleteDependents}, // Finalizer 입니다.
		},
		Spec: v1alpha2.MachineSpec{ // 소속될 infra
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha2",
				Kind:       "InfrastructureConfig",
				Name:       "infra-config1",
			},
			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")}, // 임의에 데이터
		},
	}
	time := metav1.Now() // 만들어지
	machine3 := v1alpha2.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind: "Machine", // 설정 종류
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              "delete", // 현재 상태 등록
			Namespace:         "default", // 네임스페이스 등록
			Finalizers:        []string{v1alpha2.MachineFinalizer, metav1.FinalizerDeleteDependents},  // Finalizer 입니다.
			DeletionTimestamp: &time, // 지워진 시간 임의에 데이터
		},
		Spec: v1alpha2.MachineSpec{ // 소속될 infra
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha2",
				Kind:       "InfrastructureConfig",
				Name:       "infra-config1",
			},
			Bootstrap: v1alpha2.Bootstrap{Data: pointer.StringPtr("data")},
		},
	}
	clusterList := v1alpha2.ClusterList{ // Cluster List입니다.
		TypeMeta: metav1.TypeMeta{
			Kind: "ClusterList",
		},
		Items: []v1alpha2.Cluster{
			{
				TypeMeta: metav1.TypeMeta{
					Kind: "Cluster", // Cluster 1
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testcluster", // 임의에 데이터
					Namespace: "default",
				},
			},
			{
				TypeMeta: metav1.TypeMeta{
					Kind: "Cluster", // Cluster 1
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "rainbow", // 임의에 데이터
					Namespace: "foo",
				},
			},
		},
	}

<<<<<<< HEAD
	// test run 후 나와 야 되는 내용 (답안지)
=======

	// Machine 설정
>>>>>>> 5cfa408b258d71e95da626299c0c8f3cac7d4a61
	type expected struct {
		result reconcile.Result
		err    bool
	}
	testCases := []struct {
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
			},
		},
		{
			request:     reconcile.Request{NamespacedName: types.NamespacedName{Name: machine2.Name, Namespace: machine2.Namespace}},
			existsValue: true,
			expected: expected{
				result: reconcile.Result{},
				err:    false,
			},
		},
		{
			request:     reconcile.Request{NamespacedName: types.NamespacedName{Name: machine3.Name, Namespace: machine3.Namespace}},
			existsValue: true,
			expected: expected{
				result: reconcile.Result{},
				err:    false,
			},
		},
	}

	for _, tc := range testCases {

		v1alpha2.AddToScheme(scheme.Scheme)
		// 설정으로 변경
		r := &MachineReconciler{
<<<<<<< HEAD
			Client: fake.NewFakeClient(&clusterList, &machine1, &machine2, &machine3, &infraConfig), //위에 선언된 상테를 테스트합니다.
=======
			Client: fake.NewFakeClient(&clusterList, &machine1, &machine2, &machine3, &infraConfig), // Mock Client 만들기
>>>>>>> 5cfa408b258d71e95da626299c0c8f3cac7d4a61
			Log:    log.Log,
		}
		fmt.Println(r)


		result, err := r.Reconcile(tc.request) // 결과가 반영
		if tc.expected.err {
			Expect(err).ToNot(BeNil()) // Nil 이 아니어야 한다.
		} else {
			Expect(err).To(BeNil()) // Nil 이어야 한다.
		}

		Expect(result).To(Equal(tc.expected.result)) // 다 에러없이 걍 뱉어야 함.
	}

	fmt.Println("dsdd")
}


