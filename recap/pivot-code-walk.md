# Golang 스터디 - cluster-api code-walk: pivot

# Cluster API - code walk: pivot

cluster-api의 code walk를 진행해보도록 하겠습니다. 오늘 진행할 부분은 cluster-api > cmd > phases > pivot.go 파일입니다. 
이 때까지 주로 cluster-api의 cluster_controller, machine_controller 등 controllers 디렉토리의 코드를 파악해봤습니다. 
이번에는 cluster-api의 pivot workflow를 살펴보고 해당 부분의 코드를 살펴보도록 하겠습니다.

pivot은 로컬 환경에서 생성한 cluster를 AWS(Amazon Web Service), GCP(Google Cloud Platform)과 같은 클라우드나 이전하는 환경에 
이전하는 것을 말합니다.

[클러스터 API git 저장소](https://github.com/kubernetes-sigs/cluster-api)에 있는 Readme.md에는 아래와 같이 cluster-api를 정의하고 있습니다.

```
The Cluster API is a Kubernetes project to bring declarative, Kubernetes-style APIs to cluster creation, 
configuration, and management. It provides optional, additive functionality on top of core Kubernetes.
```

쿠버네티스 클러스터를 간단하게 생성하고 관리하는 것이 존재이유인 cluster api이므로 이전하고 싶은 환경에 생성한 쿠버네티스 클러스터를 이전하는 기능을 제공하고 
있습니다.

cluster api에서 클러스터를 생성하면 일어나는 진행과정을 잘 보여주는 영상 2가지를 먼저 소개하도록 하겠습니다. cluster api에서 cluster를 생성하는 과정인 
bootstrap workflow를 설명하는 부분만 살펴보도록 하겠습니다.

## 참고영상 - Cluster API Bootstrap workflow

Cluster-api에서 cluster를 생성하는 Bootstrap workflow에 대한 부분만 참고하도록 하겠습니다.

## Deep Dive: Cluster Lifecycle SIG (Cluster API) - Robert Bailey, Google & David E. Watson, Samsung

[영상 링크](https://www.youtube.com/watch?v=LxENsOKq7UA&t=918s)
Bootstrap workflow 설명 부분: 7:36 ~ 10:11

## Deep Dive: Cluster Lifecycle SIG (Cluster API) - Jason DeTiberus, VMware & Hardik Dodiya, SAP

[영상 링크](https://www.youtube.com/watch?v=Mtg8jygK3Hs&t=1183s)
Bootstrap workflow 설명 부분: 14:23 ~ 19:45

진행과정을 정리하면 아래와 같습니다.

진행과정)

```
1. User가 CLI 도구를 사용해 Bootstrap Cluster를 생성한다.
2. Bootstrap Cluster에서 Cluster Controller, Machine Controller와 같은 Controller들을 실행한다.
3. Custom Resource Definition에 정의된 Cluster와 Machine을 실행한다.
4. Target Cluster를 생성한다.
5. To pivot or not to pivot? Target 클러스터로의 이전 여부를 결정한다.
6. Bootstrap Cluster의 Custom Resource Definition들을 Target Cluster로 복사한다.
7. Machine Cluster, Cluster Controller와 Controller들을 Target Cluster에서 실행한다.
8. Bootstrap Cluster를 제거한다.
```

그리고, 영상을 본 다음 생각해볼 수 있는 퀴즈입니다.

Quiz. Bootstrap Cluster와 Target Cluster에서 동시에 Controller들을 운영헀을 때 생기는 문제점은?
Answer: ?

두 영상 모두 cluster-api가 kubernetes의 선언적인 스타일을 cluster의 생성과 관리에 적용하려고 했다는 것을 비유적으로 잘 설명하고 있기 때문에, Q&A 세션 전까지는 모두 보시는 것을 
추천합니다.

# Code walk

참고)

* clusterctl > clusterdeployer > clusterdeployer.go 파일을 보면 Create, Delete 함수에서 호출하고 있는 것을 알 수 있습니다. 
* clusterctl > clusterclient > clusterclient.go 파일을 참고하면 Apply 등의 함수가 구현되어 있습니다.
* StatefulSet이란?
    * [쿠버네티스 문서 StatefulSet 설명](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)
    * [조대협님의 블로그 - StatefulSet을 이용한 상태유지 Pod 관리하기](https://bcho.tistory.com/1306)
* upstream 저장소와 동기화하는 방법
```
git remote add upstream https://github.com/kubernetes-sigs/cluster-api.git
git remote -v
git fetch upstream
git checkout master
git merge upstream/master
```

# Discussion

* 왜 clusterclient > clusterclient.go 에 있는 Client 인터페이스를 그대로 사용하지 않았을까?
* 왜 source cluster의 자원을 검사한 다음 target cluster의 자원도 검사할까?
* source cluster에서 Scale down하는 이유는? line96의 // Scale down the controller managers in the source cluster 주석의 의미는?
* StatefulSet이 등장하는 이유는?

# 맺음말

부족하게나마 Cluster-api의 Bootstrap workflow 중 pivot에 대한 내용을 살펴봤습니다. 선정한 뒤 참고되는 코드가 많고 간단하게 느껴지지 않아 당혹스러웠는데, 
스터디 이후 내용을 정리하여 추가하여 마무리 짓도록 하겠습니다.
