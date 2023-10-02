terraform {
  required_providers {
    Metadefender = {
      source  = "gerbil/Metadefender"
      version = "0.3.1"
    }
  }
}

provider "Metadefender" {
  # Configuration options
}