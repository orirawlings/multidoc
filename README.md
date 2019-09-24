# multidoc
Convert Kubernetes List Kinds into multidoc YAML

The Kubernetes API represents a collection of resources using a [List kind](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds). This program will read a YAML document containing a List at the top-level and convert it into a [YAML multi-document](https://yaml.org/spec/1.2/spec.html#id2760395), where each document is one of the items in the original List.

## Install

```
go get github.com/orirawlings/multidoc
```

## Use

```
$ kubectl get services -o yaml 
apiVersion: v1
items:
- apiVersion: v1
  kind: Service
  metadata:
    creationTimestamp: "2019-09-24T13:32:37Z"
    labels:
      app: example
    name: example
    namespace: default
    resourceVersion: "1180"
    selfLink: /api/v1/namespaces/default/services/example
    uid: c5da20a1-decf-11e9-a0dc-025000000001
  spec:
    externalName: www.example.com
    selector:
      app: example
    sessionAffinity: None
    type: ExternalName
  status:
    loadBalancer: {}
- apiVersion: v1
  kind: Service
  metadata:
    creationTimestamp: "2019-09-24T13:22:34Z"
    labels:
      component: apiserver
      provider: kubernetes
    name: kubernetes
    namespace: default
    resourceVersion: "156"
    selfLink: /api/v1/namespaces/default/services/kubernetes
    uid: 5e9121fc-dece-11e9-a0dc-025000000001
  spec:
    clusterIP: 10.96.0.1
    ports:
    - name: https
      port: 443
      protocol: TCP
      targetPort: 6443
    sessionAffinity: None
    type: ClusterIP
  status:
    loadBalancer: {}
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""


$ kubectl get services -o yaml | multidoc

---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2019-09-24T13:32:37Z"
  labels:
    app: example
  name: example
  namespace: default
  resourceVersion: "1180"
  selfLink: /api/v1/namespaces/default/services/example
  uid: c5da20a1-decf-11e9-a0dc-025000000001
spec:
  externalName: www.example.com
  selector:
    app: example
  sessionAffinity: None
  type: ExternalName
status:
  loadBalancer: {}

---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2019-09-24T13:22:34Z"
  labels:
    component: apiserver
    provider: kubernetes
  name: kubernetes
  namespace: default
  resourceVersion: "156"
  selfLink: /api/v1/namespaces/default/services/kubernetes
  uid: 5e9121fc-dece-11e9-a0dc-025000000001
spec:
  clusterIP: 10.96.0.1
  ports:
  - name: https
    port: 443
    protocol: TCP
    targetPort: 6443
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```
