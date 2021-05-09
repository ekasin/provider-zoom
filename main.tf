terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" { 
  apisecret="Qk2vHwnMdT0K1dXqWoFyYkMwt2CDkxOwYixV"
  apikey="lNGJBHjuROOFKCM68LjH0g"
}
