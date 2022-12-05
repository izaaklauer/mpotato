project = "mpotato"

app "mpotato" {
  config {
    env = {
      "CONFIG_PATH" = "/opt/config/config.hcl"
    }
  }

  build {
    use "docker" {}
    registry {
      // Replace with your docker registry of choice - ttl.sh is not for production
      use "docker" {
        image     = "ttl.sh/mpotato"
        tag        = "1h" // recommend gitrefpretty() for most registries
      }
    }
  }

  deploy {
    use "kubernetes" {
      replicas = 3
      namespace = "default"
      service_port = 8080
    }
  }

  release {
    use "kubernetes" {
      namespace = "default"
      load_balancer = false
    }
  }
}
