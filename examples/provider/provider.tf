terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.2.0"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}