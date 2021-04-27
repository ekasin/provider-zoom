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
 
<br> [A] Create and export a signing key that you plan on using to sign your provider releases. See [Preparing and Adding a Signing Key](#add_signing_key)

Create a Signing Key**
 1. Generate a GPG key which will be used for signing releases (https://docs.github.com/en/github/authenticating-to-github/generating-a-new-gpg-key)
 2. Export the public key in ASCII-armor format using the following command:
```bash
gpg --armor --export "[Key ID or email address]"
```

<br> [B] Copy [GoReleaser configuration](https://github.com/hashicorp/terraform-provider-scaffolding/blob/master/.goreleaser.yml) from the `terraform-provider-scaffolding repository` to `.goreleaser.yml` file in root directory of your repository.

<br> [C] Copy GitHub Actions workflow from the terraform-provider-scaffolding repository to `.github/workflows/release.yml` in your repository.

<br> [D] Go to *Settings > Secrets* in your repository, and add the following secrets : 

- `GPG_PRIVATE_KEY`  -  It is the ASCII-armored GPG private key. It can be obtained through this command `gpg --armor --export-secret-keys [key ID or email]`

- `PASSPHRASE`  -  It is the passphrase for your GPG private key.

<br> [E] Push a new valid version tag (e.g. v1.2.3) to test that the GitHub Actions releaser is working.

### Manually Preparing a Release
- There are 1 or more zip files containing the built provider binary for a single architecture
  - The binary name is `terraform-provider-[NAME]_v[VERSION]`
  - The archive name is `terraform-provider-{NAME}_{VERSION}_{OS}_{ARCH}.zip`

- There is a `terraform-provider-{NAME}_{VERSION}_SHA256SUMS` file, which contains a sha256 sum for each zip file in the release.
  - `shasum -a 256 *.zip > terraform-provider-{NAME}_{VERSION}_SHA256SUMS`

- There is a `terraform-provider-{NAME}_{VERSION}_SHA256SUMS.sig` file, which is a valid GPG binary (not ASCII armored) signature of the `terraform-provider-{NAME}_{VERSION}_SHA256SUMS` file using the keypair.
  - `gpg --detach-sign terraform-provider-{NAME}_{VERSION}_SHA256SUMS`

- Release is finalized
 
 ## Publishing to the Registry
 
 ### Signing in
 1. Sign in to the [Terraform Registry](https://www.terraform.io/docs/registry/index.html#user-account) with a GitHub account. 
 2. Following permissions should be given by the GitHub account used to the provider repository you wish to publish. <br>
 ![alt text](https://www.terraform.io/docs/registry/providers/images/github-oauth-permissions-8f791b2d.png)
 
 ### <a name="add_signing_key"></a>Preparing and Adding a Signing Key
 1. Generate a GPG key which will be used for signing releases (https://docs.github.com/en/github/authenticating-to-github/generating-a-new-gpg-key)
 2. Export the public key in ASCII-armor format using the following command:
```bash
gpg --armor --export "[Key ID or email address]"
```
 
 
 
