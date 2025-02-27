---
page_title: "Resource octopusdeploy_aws_elastic_container_registry - terraform-provider-octopusdeploy"
subcategory: "Feeds"
description: |-
  This resource manages a AWS Elastic Container Registry in Octopus Deploy.
---

# Resource (octopusdeploy_aws_elastic_container_registry)

This resource manages a AWS Elastic Container Registry in Octopus Deploy.

## Example Usage

```terraform
resource "octopusdeploy_aws_elastic_container_registry" "example" {
  access_key = "access-key"
  name       = "Test AWS Elastic Container Registry (OK to Delete)"
  region     = "us-east-1"
  secret_key = "secret-key"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **access_key** (String) The AWS access key to use when authenticating against Amazon Web Services.
- **name** (String) A short, memorable, unique name for this feed. Example: ACME Builds.
- **region** (String) The AWS region where the registry resides.
- **secret_key** (String, Sensitive) The AWS secret key to use when authenticating against Amazon Web Services.

### Optional

- **id** (String) The unique ID for this feed.
- **package_acquisition_location_options** (List of String)
- **space_id** (String) The space ID associated with this feed.


