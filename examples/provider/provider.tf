terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.11.0"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}