---
page_title: "Resource octopusdeploy_feed - terraform-provider-octopusdeploy"
subcategory: "Feeds"
description: |-
  
---

# Resource (octopusdeploy_feed)



## Example Usage

```terraform
resource "octopusdeploy_feed" "example" {
  download_attempts                    = 10
  download_retry_backoff_seconds       = 60
  feed_uri                             = "https://repo.maven.apache.org/maven2/"
  name                                 = "Test Maven Feed (OK to Delete)"
  package_acquisition_location_options = ["Server", "ExecutionTarget"]
  password                             = "password123"
  space_id                             = "Spaces-123"
  username                             = "john.smith@example.com"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **feed_uri** (String)
- **name** (String) A short, memorable, unique name for this feed. Example: ACME Builds.

### Optional

- **access_key** (String)
- **api_version** (String)
- **delete_unreleased_packages_after_days** (Number)
- **download_attempts** (Number) The number of times a deployment should attempt to download a package from this feed before failing.
- **download_retry_backoff_seconds** (Number) The number of seconds to apply as a linear back off between download attempts.
- **feed_type** (String)
- **id** (String) The unique ID for this resource.
- **is_enhanced_mode** (Boolean)
- **package_acquisition_location_options** (List of String)
- **password** (String, Sensitive) The password associated with this resource.
- **registry_path** (String)
- **secret_key** (String, Sensitive)
- **space_id** (String) The space ID associated with this resource.
- **username** (String, Sensitive) The username associated with this resource.

### Read-Only

- **region** (String)


