terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.6.0"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}