# chowkidar
A temporary watch man for your resources

## Manifest file :
```yaml
apiVersion: crd.devops.tools/v1
kind: Chowki
metadata:
  labels:
    app.kubernetes.io/name: chowkidar
    app.kubernetes.io/managed-by: kustomize
  name: chowki-sample
#  namespace: nginx-namespace
spec:
  # TODO(user): Add fields here
    resoureType: "pod"
    resourceName: "nginx-pod" 
```

## Reference
[Go Operator](https://github.com/ManojChandran/go-operator)