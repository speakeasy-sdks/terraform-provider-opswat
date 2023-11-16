terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.8.3"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}