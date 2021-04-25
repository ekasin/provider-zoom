# Terraform_Zoom_Provider

This providre allows to perform Create ,Read ,Update and Delete user for zoom channel

## Requirements

terraform zoom provider is based on Terraform, this means that you need

1. Go lang >=1.11 <br>
2. Terraform>=v0.13.0 <br/>
3. Zoom Pro account (Token)

## Installation

1. Download required binaries<br>
2. move binary ~/.terraform.d/plugins/[architecture name]/

## Comands to Run the Provider

1. terraform init <br/>
2. terraform plan <br>
3. terrafrom apply (To create or update the user)<br>
4. terraform destroy (To destroy the created user)<br>

## Steps to run import command

1. Write manually a resource configuration block for the resource, to which the imported object will be mapped.
2. RUN terraform import zoom_user.sample <user_id> 