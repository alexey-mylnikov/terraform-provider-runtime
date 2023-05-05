---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "runtime Data Source - terraform-provider-runtime"
subcategory: ""
description: |-
  The runtime data source allows determine some information about host operating system.
---

# runtime (Data Source)

The `runtime` data source allows determine some information about host operating system.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `arch` (String) The running program's operating system target: one of darwin, freebsd, linux, and so on.
- `id` (String) The id of the data source. This will always be set to `-`
- `os` (String) The running program's architecture target: one of 386, amd64, arm, s390x, and so on.

