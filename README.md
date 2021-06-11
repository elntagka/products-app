# products-app
This project creates and deploys a Go service that accepts requests in the path `/api/products/` and returns some sample data. In addition to that it supports pagination and will return a subset of the test data when a specific page is requested.
This project is deployed on a Kubernetes cluster on GCP using Terraform, Helm and the Google Cloud Build service. Below there are more information about how to deploy and test this project.

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
- This project repository is stored on a compatible with Cloud Build repository and the trigger is already configured to run on every push to the `main` branch. For this example, I used a private repository under my personal GitHub account and the trigger was configured with the below substitution variables:
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

### Application deployment using Helm3 and the Cloud Build service
The steps to deploy the `products-app` service are outlined in the `cloudbuild.yaml` file which defines the steps Cloud Build will follow on every push to the `main` branch, which include:
- Go unit test execution
- Docker build and push
- A Helm upgrade dry run to review the Kubernetes resources prior to deployment
- A Helm upgrade with option to install if a release by this name doesn't already exist
- JMeter load tests to prove the autoscaling configuration

*Notes:* 
- There's a limitation currently on the JMeter step due to time constaints on this work. The public IP assigned to the ingress should be updated in the JMeter step in `cloudbuild.yaml` in the parameter `-JDomain=`. So the limitation of the current example is that user the first time the pipeline runs should deploy the Helm release, get the IP address of the ingress using the command below, update the pipeline `cloudbuild.yaml` and trigger a new build for JMeter load tests to run against the correct IP.
    ```
    kubectl get ingress -n products-app
    ```
- In addition to this pipeline, there's also a Makefile with the helm commands that I used for local testing and I'm leaving here as reference or for future work*

### Testing and Validating the solution
#### Testing the products-app
In addition to the unit tests that run as part of the pipeline, once the above deployments are completed and the ingress IP is available, the application can be tested by simply running curl commands like the below:
```
curl "http://${IP}/api/products?page=5" -i

curl "http://${IP}/api/products?page=five" -i

curl "http://${IP}/api/products?page=100" -i

curl "http://${IP}/api/products?animal=5" -i

curl "http://${IP}/api/products" -i
```

#### Validating the autoscaling 
The jmx file provided has a test configured to prove the scaling configuration that this cluster was tested with. To validate this while the JMeter step of the pipeline runs you can monitor the workload on the GCP console or via using kubectl. For reference, during my tests the deployment scaled from 1 pod to 19.
