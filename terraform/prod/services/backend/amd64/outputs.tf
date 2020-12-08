output "dns_name" {
  value       = module.backend.public_dns
  description = "The domain name of the server"
}