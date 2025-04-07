# CI/CD workshop repository

## Prerequisites

### Install Argo CD in k8s

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### Install Argo CD Argo Rollouts

```bash
kubectl create namespace argoproj
kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml
```

## Add a repository to Argo CD

Get the admin password and use it to login with argo cd:
```shell
kubectl get secret argocd-initial-admin-secret -n argocd   -o jsonpath="{.data.password}" | base64 -d && echo
argocd login localhost:8080 --username admin --password <password>
```

Then add the git repo to Argo CD:
```bash
export USERNAME=<your-github-repository>
export REPO=go-simple-webserver
argocd app create go-simple-webserver \
  --repo https://github.com/$USERNAME/$REPO.git \
  --path manifests \
  --dest-server https://kubernetes.default.svc \
  --dest-namespace default
```

## Argo UI Guideline

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

To login use the following credentials: `admin/<password>`

`<password>` can be retrieved from `kubectl get secret argocd-initial-admin-secret -n argocd   -o jsonpath="{.data.password}" | base64 -d && echo`
