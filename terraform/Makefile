TF_OUT_PLAN := ./tf.out
ENV := dev
TF_CONFIG := config/$(ENV)/env.tfvars
TF_BACKEND := config/$(ENV)/backend.tfvars


clean:
	@if [ -d .terraform ] ; then \
		rm -fr .terraform ; \
	fi
	@if [ -f "$(TF_OUT_PLAN)" ] ; then \
		rm -f "$(TF_OUT_PLAN)"; \
	fi

init:clean
	@terraform init \
		-backend=true \
		-input=false \
		-upgrade \
		-backend-config=$(TF_BACKEND)

plan:init
	@terraform plan \
		-input=false \
		-lock=true \
		-refresh=true \
		-out=$(TF_OUT_PLAN) \
		-var-file=$(TF_CONFIG)

apply:
	@terraform apply \
		-input=false \
		-auto-approve=true \
		-lock=true \
		-refresh=true \
		-state=$(TF_OUT_PLAN) \
		-var-file=$(TF_CONFIG)
	@if [ -f "$(TF_OUT_PLAN)" ] ; then \
		rm -f "$(TF_OUT_PLAN)"; \
	fi

destroy: 
	@terraform destroy \
		-input=false \
		-lock=true \
		-state=$(TF_OUT_PLAN) \
		-var-file $(TF_CONFIG)