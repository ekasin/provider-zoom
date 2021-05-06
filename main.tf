terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  apikey = ""
  apisecret = ""
}


resource "zoom_user" "user1" {
   email      = "tapendrakmr786@gmail.com"
   first_name = "tapendra"
   last_name  = "kumar"
   status = "activate"
}


resource "zoom_user" "user2" {
   email      = "ashishdhodria27@gmail.com"
   first_name = "coding"
   last_name  = "ninza"
   status = "activate"
}

data "zoom_user" "user2" {
  id = "ekansh0786@gmail.com"
}


output "user2" {
  value = data.zoom_user.user2
}


