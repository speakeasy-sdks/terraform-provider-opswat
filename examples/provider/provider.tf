terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.4.1"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}