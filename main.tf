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
  token   = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E"
}


resource "zoom_user" "user1" {
  email      = "ui17co14@iiitsurat.ac.in"
  first_name = "ekansh"
  last_name  = "rock"

}




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



