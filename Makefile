RELEASE_NAME:=testdeploy
NAMESPACE:=products-app
IMAGE_TAG:=5e699f7

helm-template:
	@helm template helm/products-app

helm-upgrade-dry:
	@helm upgrade ${RELEASE_NAME} helm/products-app --install --create-namespace --namespace ${NAMESPACE} --dry-run

helm-upgrade:
	@helm upgrade ${RELEASE_NAME} helm/products-app --install --create-namespace --namespace ${NAMESPACE} --set image.tag=${IMAGE_TAG}

helm-uninstall:
	@helm uninstall ${RELEASE_NAME} -n ${NAMESPACE}

port-forward:
	@kubectl port-forward service/${RELEASE_NAME}-products-app 8000:8080 -n ${NAMESPACE}