helm upgrade --install --namespace default -f values.yaml blue .
helm upgrade --install --namespace default -f values.yaml green --set image.tag=0.2.0 --set ingress.enabled=false .
