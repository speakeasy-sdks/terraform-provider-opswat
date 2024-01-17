terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.11.3"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}