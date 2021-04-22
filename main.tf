terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  address = "https://api.zoom.us/v2/users"
  token   = "access_token"
}

/*
resource "zoom_user" "user1" {
  email      = "ui17co14@iiitsurat.ac.in"
  first_name = "ekansh"
  last_name  = "singh"
  active = "activate"
}
*/

resource "zoom_user" "user1" {
    email      = "tapendrakmr786@gmail.com"
    first_name = "ekansh"
    last_name  = "rock"
    active = "activate"
}


resource "zoom_user" "user2" {
  #  email      = "ekansh0786@gmail.com"
  #  first_name = "ekansh"
  #  last_name  = "rock"
  #  active = "activate"
}

output "user2" {
  value = zoom_user.user2
}


/*

resource "zoom_user" "user3" {
    email      = "ekansh0786@gmail.com"
    first_name = "ekansh"
    last_name  = "rock"
    active = "activate"
}

output "user3" {
  value = zoom_user.user3
}

*/
/*
data "zoom_user" "user1" {
  id = "ekansh0786@gmail.com"
}


output "user1" {
  value = data.zoom_user.user1
}

*/


/*
data "zoom_user" "user2" {
  id = ""
}


output "user2" {
  value = data.zoom_user.user2
}
*/



