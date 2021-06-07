# products-app
Hi there! Thanks for the opportunity to work on this project I really enjoyed it!
Please see below how to deploy/test it.

## Prerequisites/Assumptions:
- Terraform service account exists and has the required permissions. For this example I assigned the role `Owner` and the json file with the account credentials is available under the directory `terraform/creds/key.json`
- Terraform remote state is configured using a a bucket on GCP
- GCP account exists
- GCP project exists. For this example I'm using `products-app-aw`
- Docker repository for the `product-app` image exists. In this example, I'm using `gcr.io/products-app-aw/products-app`  
- GCP services required are enabled
- Cloud Build service account has admin permissions against the Kubernetes cluster. For this example, I had assigned the role `Owner`
- The Helm cloud builder already exists on GCR. For this example I build and push the community builder from [here](https://github.com/GoogleCloudPlatform/cloud-builders-community/tree/master/helm)
- The JMeter cloud builder already exists on GCR. For this example, I fixed some issues on the community builder which are currently under my fork [here](https://github.com/elntagka/cloud-builders-community/tree/master/jmeter)
- This project repository is stored on a compatible with Cloud Build repository and the trigger is already configured to run on every push to the `main` branch with the below substitution variables:
    ```
    _CLUSTER_NAME = products-app-cluster-dev
    _CLUSTER_REGION = europe-west1
    _HELM_RELEASE_NAME = testdeploy
    _NAMESPACE = products-app
    _SERVICE_NAME = products-app
    ```

## Deployment and Testing

### GKE Cluster Deployment using Terraform
In order to deploy the GKE cluster with the default values for this example, change into the terraform directory and use the makefile to perform and Terraform plan and apply
```
cd terraform
make plan
make apply
```
Once the deployment is succesful use the command below which utilises the `kubeconfig-dev` file that will be available at the end of the deployment to confirm your access to the newly created cluster and check the default namespaces available.
```
kubectl get ns --kubeconfig=kubeconfig-dev
```

###