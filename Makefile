SHA := $(shell git rev-parse --short HEAD)
AMI := $(shell jq -r '.builds[0].artifact_id|split(":")[1]' ./manifest.json)
package:
	packer build -force -machine-readable -var "git_sha=$(SHA)" -var "aws_access_key=${AWS_ACCESS_KEY_ID}" -var "aws_secret_key=${AWS_SECRET_ACCESS_KEY}" hello-image.pkr.hcl

ship:
	cd terraform && \
	terraform apply -auto-approve -var "hello_ami=$(AMI)"

certs:
	certstrap init --passphrase "" --cn vault-ca && \
	certstrap request-cert --passphrase "" --domain vault --ip 127.0.0.1 && \
	certstrap sign vault --CA vault-ca && \
	certstrap request-cert --passphrase "" --cn concourse && \
	certstrap sign concourse --CA vault-ca && \
	mv out vault-certs && \
	chmod +r vault-certs/*