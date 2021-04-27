# STEPS TO PUBLISH A TERRAFORM PROVIDER

Reference Link: https://www.terraform.io/docs/registry/providers/publishing.html

## Documenting your Provider
1. Create a index.md
2. Create document for each resource and data source
https://www.terraform.io/docs/registry/providers/docs.html

3.  Generating Documentation 
https://github.com/hashicorp/terraform-plugin-docs

* Automatically generate documentation for your provider in the format necessary for the Terraform Registry - https://github.com/hashicorp/terraform-plugin-docs

* The terraform-provider-scaffolding template repository includes example usage of the tfplugindocs command via go generate:

```GO
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
```

For more details on the Document format refer : https://www.terraform.io/docs/registry/providers/docs.html#format

## Creating a GitHub Release 
1. Publishing a provider requires at least one version be available on GitHub Releases. The tag must be a valid [Semantic Version](https://semver.org/) preceded with a v (for example, v1.2.3).

<br> List of [recommend OS / architecture combinations](https://www.terraform.io/docs/registry/providers/os-arch.html) for which we suggest most providers create binaries.
 
 ### GitHub Actions (Preferred)
 
<br><b> [A] Create a Signing Key </b>
 1. Generate a GPG key which will be used for signing releases (https://docs.github.com/en/github/authenticating-to-github/generating-a-new-gpg-key)
 2. Export the public key in ASCII-armor format using the following command:
```bash
gpg --armor --export "[Key ID or email address]"
```

<br> [B] Copy [GoReleaser configuration](https://github.com/hashicorp/terraform-provider-scaffolding/blob/master/.goreleaser.yml) from the `terraform-provider-scaffolding repository` to `.goreleaser.yml` file in root directory of your repository.

<br> [C] Copy GitHub Actions workflow from the terraform-provider-scaffolding repository to `.github/workflows/release.yml` in your repository.

<br> [D] Go to *Settings > Secrets* in your repository, and add the following secrets : 
     * `GPG_PRIVATE_KEY`  -  It is the ASCII-armored GPG private key. It can be obtained through this command `gpg --armor --export-secret-keys [key ID or email]`
     * `PASSPHRASE`  -  It is the passphrase for your GPG private key.


 
 
 
 
