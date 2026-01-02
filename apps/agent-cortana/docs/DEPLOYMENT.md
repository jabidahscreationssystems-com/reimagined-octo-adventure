# Deployment Guide for Agent Cortana

This guide provides instructions for deploying Agent Cortana GitHub App to various platforms.

## Prerequisites

Before deploying Agent Cortana, you need:

1. A GitHub App created in your GitHub account or organization
2. GitHub App credentials (App ID, Private Key)
3. A server or platform to host the application

## Creating a GitHub App

1. Navigate to GitHub Settings → Developer settings → GitHub Apps
2. Click "New GitHub App"
3. Fill in the required information:
   - **App name**: Agent Cortana
   - **Homepage URL**: Your application URL
   - **Webhook URL**: `https://your-domain.com/webhooks`
   - **Webhook secret**: Generate a secure secret
4. Set permissions:
   - Repository permissions:
     - Contents: Read & Write
     - Issues: Read & Write
     - Pull requests: Read & Write
     - Workflows: Read & Write
   - Subscribe to events:
     - Issues
     - Issue comment
     - Pull request
     - Pull request review
     - Pull request review comment
     - Push
     - Workflow run
5. Click "Create GitHub App"
6. Generate and download a private key
7. Note your App ID

## Environment Configuration

Create a `.env` file based on `.env.example`:

```bash
PORT=8080
GITHUB_APP_ID=123456
GITHUB_APP_PRIVATE_KEY="-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----"
WEBHOOK_SECRET=your_webhook_secret
```

## Deployment Options

### Option 1: Local Development

```bash
# Install dependencies
go mod download

# Run the server
make run
```

### Option 2: Docker

```bash
# Build the Docker image
make docker-build

# Run with docker-compose
make docker-run
```

### Option 3: Kubernetes

Create a Kubernetes deployment:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: agent-cortana
spec:
  replicas: 2
  selector:
    matchLabels:
      app: agent-cortana
  template:
    metadata:
      labels:
        app: agent-cortana
    spec:
      containers:
      - name: agent-cortana
        image: agent-cortana:latest
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        - name: GITHUB_APP_ID
          valueFrom:
            secretKeyRef:
              name: github-app-secrets
              key: app-id
        - name: GITHUB_APP_PRIVATE_KEY
          valueFrom:
            secretKeyRef:
              name: github-app-secrets
              key: private-key
        - name: WEBHOOK_SECRET
          valueFrom:
            secretKeyRef:
              name: github-app-secrets
              key: webhook-secret
---
apiVersion: v1
kind: Service
metadata:
  name: agent-cortana
spec:
  selector:
    app: agent-cortana
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

Apply the deployment:

```bash
kubectl apply -f deployment.yaml
```

### Option 4: Cloud Platforms

#### AWS Elastic Beanstalk

1. Create a `Dockerrun.aws.json`:
```json
{
  "AWSEBDockerrunVersion": "1",
  "Image": {
    "Name": "agent-cortana:latest"
  },
  "Ports": [
    {
      "ContainerPort": 8080
    }
  ]
}
```

2. Deploy using EB CLI:
```bash
eb init -p docker agent-cortana
eb create agent-cortana-env
eb deploy
```

#### Google Cloud Run

```bash
# Build and push to Google Container Registry
gcloud builds submit --tag gcr.io/PROJECT_ID/agent-cortana

# Deploy to Cloud Run
gcloud run deploy agent-cortana \
  --image gcr.io/PROJECT_ID/agent-cortana \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars GITHUB_APP_ID=123456
```

#### Azure Container Instances

```bash
# Build and push to Azure Container Registry
az acr build --registry myregistry --image agent-cortana:latest .

# Deploy to Container Instances
az container create \
  --resource-group myResourceGroup \
  --name agent-cortana \
  --image myregistry.azurecr.io/agent-cortana:latest \
  --dns-name-label agent-cortana \
  --ports 8080
```

## Post-Deployment

1. Verify the deployment:
   ```bash
   curl https://your-domain.com/health
   ```

2. Update your GitHub App webhook URL to point to your deployment

3. Install the GitHub App on your repositories

4. Test webhook delivery by creating an issue or pull request

## Monitoring

Monitor your Agent Cortana deployment:

- Check application logs for webhook events
- Monitor the `/health` endpoint for uptime
- Set up alerts for failures

## Troubleshooting

### Webhook not receiving events

- Verify webhook URL is publicly accessible
- Check webhook secret matches configuration
- Review GitHub App webhook delivery logs

### Authentication errors

- Verify GitHub App ID is correct
- Ensure private key is properly formatted
- Check App permissions are sufficient

### Connection issues

- Verify firewall rules allow incoming HTTPS traffic
- Check SSL/TLS certificates are valid
- Ensure DNS is properly configured

## Security Considerations

1. Always use HTTPS for webhook URLs
2. Validate webhook signatures
3. Store secrets securely (use environment variables or secret managers)
4. Limit GitHub App permissions to minimum required
5. Regularly rotate private keys
6. Monitor for suspicious activity

## Scaling

For high-traffic deployments:

1. Run multiple instances behind a load balancer
2. Use a message queue for webhook processing
3. Implement rate limiting
4. Cache GitHub API responses
5. Use webhooks selectively (subscribe only to needed events)
