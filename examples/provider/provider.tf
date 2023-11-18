terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.9.0"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}