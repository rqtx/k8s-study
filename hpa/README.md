```bash
    kind create cluster --config kind-config.yml
```

```bash
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo update
```

```bash
    helm install prometheus prometheus-community/kube-prometheus-stack \
    -n monitoring \
    --create-namespace
```

```bash
    helm install prometheus-adapter prometheus-community/prometheus-adapter \
    -n monitoring \
    -f values-adapter.yml
```