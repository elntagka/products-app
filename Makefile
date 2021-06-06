RELEASE_NAME:=testdeploy
NAMESPACE:=products-app

helm-install:
	@helm install ${RELEASE_NAME} helm/products-app --create-namespace --namespace ${NAMESPACE}

helm-upgrade:
	@helm upgrade ${RELEASE_NAME} helm/products-app --namespace ${NAMESPACE}

helm-diff:
	@helm diff ${RELEASE_NAME} helm/products-app --namespace ${NAMESPACE}

port-forward:
	@kubectl port-forward service/${RELEASE_NAME}-products-app 8000:8080 -n ${NAMESPACE}

helm-uninstall:
	@helm uninstall ${RELEASE_NAME} -n ${NAMESPACE}
