terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" { 
  #set API_SECRET environmnemt variable
  #set API_KEY environmnemt variable
}

resource "zoom_user" "user1" {
   email      = ""
   first_name = ""
   last_name  = ""
   status = "active"
   license_type =   1
   department = ""
   job_title = ""
   location = ""
}
