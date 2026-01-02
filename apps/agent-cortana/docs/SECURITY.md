# Security Best Practices for Agent Cortana

This document outlines security considerations and best practices when deploying and operating Agent Cortana.

## Environment Variables and Secrets

### Never Commit Secrets

- **NEVER** commit `.env` files or any files containing secrets to version control
- Use `.env.example` as a template only
- Secrets should be stored in:
  - Environment variables
  - Secret management systems (AWS Secrets Manager, Azure Key Vault, Google Secret Manager)
  - Kubernetes Secrets
  - CI/CD secret stores

### Required Secrets

The following secrets must be protected:

1. `GITHUB_APP_ID` - While not strictly secret, should be controlled
2. `GITHUB_APP_PRIVATE_KEY` - **Critical**: This authenticates your app
3. `GITHUB_APP_SECRET` - Used for OAuth flows (optional)
4. `WEBHOOK_SECRET` - Used to validate webhook signatures

### Private Key Handling

- Store private keys in secure secret management systems
- Use key rotation policies
- Never log or print private keys
- Use environment variables that are encrypted at rest

## Webhook Security

### Signature Validation (TODO - Implementation Required)

**IMPORTANT**: The current webhook handler does NOT validate signatures. Before production use, implement signature validation:

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

func validateSignature(payload []byte, signature, secret string) bool {
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write(payload)
    expectedMAC := hex.EncodeToString(mac.Sum(nil))
    return hmac.Equal([]byte("sha256="+expectedMAC), []byte(signature))
}
```

### Webhook Best Practices

1. **Always validate webhook signatures** using `X-Hub-Signature-256` header
2. **Use HTTPS only** for webhook URLs
3. **Implement rate limiting** to prevent DoS attacks
4. **Log webhook delivery IDs** for audit trails
5. **Validate event payloads** before processing
6. **Use webhook secrets** that are long and random

## Network Security

### HTTPS/TLS

- **Always use HTTPS** in production
- Use valid TLS certificates (Let's Encrypt, commercial CA)
- Enforce TLS 1.2 or higher
- Configure proper certificate validation

### Firewall Rules

- Restrict inbound traffic to necessary ports only
- Allow GitHub webhook IPs (see [GitHub's Meta API](https://api.github.com/meta))
- Use security groups/network policies to limit access

## Authentication and Authorization

### GitHub App Permissions

Follow the **principle of least privilege**:

1. Only request permissions your app needs
2. Use read-only permissions when possible
3. Regularly audit and review permissions
4. Document why each permission is needed

Current permissions in `app.yml`:
- `contents: write` - For modifying repository files
- `issues: write` - For creating/updating issues
- `pull_requests: write` - For managing PRs
- `metadata: read` - Basic repository metadata
- `workflows: write` - For triggering workflows

### API Rate Limiting

- Implement exponential backoff for API calls
- Cache responses when appropriate
- Use conditional requests with ETags
- Monitor rate limit headers

## Input Validation

### Webhook Payloads

- Validate all input from webhook payloads
- Sanitize user-provided content before using in API calls
- Prevent injection attacks (SQL, command, etc.)
- Validate event types before processing

### Example Validation

```go
func validateIssueEvent(payload map[string]interface{}) error {
    action, ok := payload["action"].(string)
    if !ok {
        return fmt.Errorf("invalid action field")
    }
    
    validActions := map[string]bool{
        "opened": true,
        "closed": true,
        "edited": true,
    }
    
    if !validActions[action] {
        return fmt.Errorf("unsupported action: %s", action)
    }
    
    return nil
}
```

## Logging and Monitoring

### What to Log

- Webhook delivery IDs
- Event types and actions
- API call outcomes
- Errors and exceptions
- Security events (failed auth, invalid signatures)

### What NOT to Log

- Private keys
- Webhook secrets
- User tokens
- Personal information (unless necessary and encrypted)
- Full payloads (may contain sensitive data)

### Log Security

- Store logs securely
- Implement log rotation
- Encrypt logs at rest
- Use centralized logging (ELK, Splunk, CloudWatch)
- Set up alerts for security events

## Deployment Security

### Container Security

- Use minimal base images (Alpine, Distroless)
- Scan images for vulnerabilities
- Don't run as root user
- Use read-only filesystems when possible
- Keep base images updated

### Kubernetes Security

- Use NetworkPolicies to restrict traffic
- Implement PodSecurityPolicies
- Use RBAC for access control
- Store secrets in Kubernetes Secrets (encrypted)
- Use service mesh for mTLS

## Incident Response

### Prepare for Security Incidents

1. **Have a response plan**:
   - Who to contact
   - How to revoke access
   - How to rotate credentials

2. **Enable audit logging**:
   - GitHub audit log
   - Application logs
   - Infrastructure logs

3. **Regular backups**:
   - Configuration
   - State data
   - Secrets (encrypted)

### In Case of Compromise

1. **Immediately revoke** compromised credentials
2. **Rotate all secrets** including:
   - GitHub App private key
   - Webhook secrets
   - Any other credentials
3. **Review audit logs** for unauthorized access
4. **Notify stakeholders**
5. **Document the incident**

## Regular Security Practices

### Ongoing Maintenance

- **Update dependencies** regularly
- **Scan for vulnerabilities** using tools like:
  - `go mod` vulnerability scanning
  - Snyk
  - Dependabot
- **Review access logs** periodically
- **Conduct security audits**
- **Test disaster recovery** procedures

### Code Security

- Use static analysis tools (golint, go vet, staticcheck)
- Implement code review processes
- Use branch protection rules
- Require signed commits
- Enable GitHub security features:
  - Dependabot alerts
  - Code scanning
  - Secret scanning

## Compliance Considerations

Depending on your use case, you may need to comply with:

- **GDPR** - For EU user data
- **SOC 2** - For service organizations
- **HIPAA** - For healthcare data
- **PCI DSS** - For payment data

Ensure you understand your compliance requirements and implement necessary controls.

## Security Checklist

Before deploying to production:

- [ ] Implement webhook signature validation
- [ ] Use HTTPS with valid certificates
- [ ] Store secrets in secure secret management
- [ ] Enable logging and monitoring
- [ ] Configure firewall rules
- [ ] Implement rate limiting
- [ ] Set up alerts for security events
- [ ] Review and minimize GitHub App permissions
- [ ] Scan container images for vulnerabilities
- [ ] Set up incident response procedures
- [ ] Enable GitHub security features
- [ ] Document security procedures

## Resources

- [GitHub Security Best Practices](https://docs.github.com/en/code-security)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)

## Reporting Security Issues

If you discover a security vulnerability in Agent Cortana:

1. **DO NOT** open a public issue
2. Report to the security team privately
3. Include details about the vulnerability
4. Allow time for a fix before public disclosure
