plan:
	terraform plan -out=plans/plan.tfplan

apply:
	terraform apply plans/plan.tfplan

init:
	terraform init

# destroy:
# 	terraform destroy
# ENV_FILE := ../../../envs/s3.env

# env:
# 	@echo "CLIENT_UI_ACCESS_KEY_ID=$$(terraform output -raw client_ui_access_key_id)" >  $(ENV_FILE)
# 	@echo "CLIENT_UI_SECRET_KEY=$$(terraform output -raw client_ui_secret_key)"       >> $(ENV_FILE)
# 	@echo "ADMIN_UI_ACCESS_KEY_ID=$$(terraform output -raw admin_ui_access_key_id)" >> $(ENV_FILE)
# 	@echo "ADMIN_UI_SECRET_KEY=$$(terraform output -raw admin_ui_secret_key)"       >> $(ENV_FILE)
# 	@echo "BUCKET_NAME=$$(terraform output -raw bucket_name)"                       >> $(ENV_FILE)
# 	@echo "generated $(ENV_FILE)"
#
#
