# employees-directory-simple

This demonstration repository will execute the following steps to deploy the Employees Directory API to Kong Gateway Enterprise:
  
- Installs Inso CLI (https://insomnia.rest/products/inso)
- Installs decK (https://docs.konghq.com/deck/latest/)
- Lints the OpenAPI Spec (OAS) using Inso CLI
- Generate Kong declarative configuration from OAS using decK
- Uploads Kong declarative config to Artifact Repository
- Validates Kong declarative config using decK
- Diffs declarative config using decK
- Backup existing Kong configuration using decK
- Uploads Kong config backup to Artifact Repository
- Deploys declarative config to development environment using decK
- Runs Unit Tests using Inso CLI
