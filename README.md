# Concourse with Vault for Secrets

## Step By Step

- Install [certstrap](https://github.com/square/certstrap)

```
wget https://github.com/square/certstrap/releases/download/v1.2.0/certstrap-1.2.0-linux-amd64
sudo mv certstrap-1.2.0-linux-amd64 /usr/local/bin/certstrap
sudo chmod +x /usr/local/bin/certstrap
```

- Create Certs

```
certstrap init --cn vault-ca
certstrap request-cert --domain vault --ip 127.0.0.1
certstrap sign vault --CA vault-ca
certstrap request-cert --cn concourse
certstrap sign concourse --CA vault-ca
mv out vault-certs
chmod +r vault-certs/*
```

- Login to Vault

```
export VAULT_CACERT=$PWD/vault-certs/vault-ca.crt
vault operator init
vault operator unseal
vault login
```

- Initialize Concourse Secrets

```
vault policy write concourse ./concourse-policy.hcl
vault auth enable cert
vault write auth/cert/certs/concourse policies=concourse certificate=@vault-certs/vault-ca.crt ttl=1h
vault secrets enable -path=concourse kv
```

- Add New Secrets
```
vault kv put concourse/main/hello value=world
```

Note: all secrets reference `value` by default, use (`github.pub` for `concourse/main/github pub=somesecret`)

- Create TF State (only once). Will need to be a unique bucket. Refereneced throughout the project.
```
cd terraform/global/s3
terraform init
terraform apply
```

## Running Go Example

- `vault kv put concourse/main/github private-key="$(cat ~/.ssh/YOUR_PRIVATE_KEY)" access-token="$GITHUB_KEY"`
- `vault kv put concourse/main/aws access-key="$AWS_ACCESS_KEY_ID" secret-access-key="$AWS_SECRET_ACCESS_KEY"`
- `vault kv put concourse/main/docker username="holmes89" access-token="$DOCKER_ACCESS_TOKEN"`
- `fly --target hello login --concourse-url http://127.0.0.1:8080 -u admin -p admin`
- `fly -t hello set-pipeline -c hello_pipeline.yml -p hello`


## Sources
- [Hello JSON](https://github.com/novellac/multilanguage-hello-json)