# GitHub Copilot Instructions

## Project Overview
This is a Kubernetes operator project built with Kubebuilder. The project includes automated workflows to keep the Kubebuilder scaffold up-to-date with the latest releases.

## Key Components

### Kubebuilder Operator
- Built using the Kubebuilder framework
- Follow Kubernetes operator patterns and best practices
- Use controller-runtime for reconciliation logic
- Implement proper error handling and status updates

### Automated Updates
- The `.github/workflows/kb-alpha-update.yml` workflow automatically updates Kubebuilder scaffolding
- Runs every Tuesday at midnight UTC via cron schedule
- Creates draft PRs with AI-generated descriptions when updates are available
- Falls back to creating issues if PR creation fails

## **CRITICAL: Kubebuilder Update Issue Handling**

### For issues with titles matching "update kubebuilder from version vX.Y.Z to version vA.B.C":

**DO NOT perform any kubebuilder commands or generate new scaffolding.**

**ONLY do the following:**
1. Find the most recent branch named `kubebuilder-alpha-update-*` (where * is a number)
2. Switch to that branch: `git checkout kubebuilder-alpha-update-<number>`
3. Create a PR from that existing branch to main
4. If there are conflict markers (`<<<<<<<`, `=======`, `>>>>>>>`) in any files:
   - Resolve the conflicts by choosing the appropriate code sections
   - Remove the conflict markers
   - Ensure the code compiles and follows project conventions
   - Commit the resolved conflicts

**Example workflow:**
```bash
# Find the latest update branch
git branch -r | grep kubebuilder-alpha-update | sort -V | tail -1

# Switch to that branch
git checkout kubebuilder-alpha-update-123

# Check for conflicts
grep -r "<<<<<<< HEAD" . || echo "No conflicts found"

# If conflicts exist, resolve them manually and commit
git add .
git commit -m "resolve: fix merge conflicts in kubebuilder update"

# Create PR (only after conflicts are resolved)
gh pr create --title "chore: update Kubebuilder scaffold" --body "..." --base main
```

## Code Style Guidelines

### Go Code
- Follow standard Go conventions and formatting (use `gofmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Implement proper error handling with wrapped errors
- Use structured logging with the controller-runtime logger

### YAML Files
- Use 2-space indentation for YAML files
- Follow Kubernetes resource naming conventions
- Include proper labels and annotations
- Use meaningful resource names

### GitHub Actions
- Use descriptive step names
- Include proper error handling with `continue-on-error` where appropriate
- Use environment variables for sensitive data
- Add comments for complex shell scripts

## Kubebuilder Patterns
- Controllers should implement the Reconciler interface
- Use finalizers for cleanup logic
- Implement proper status conditions
- Follow the controller pattern: observe, analyze, act
- Use owner references for dependent resources

## Testing
- Write unit tests for business logic
- Use the envtest package for integration tests
- Test error conditions and edge cases
- Mock external dependencies in tests

## Documentation
- Update README.md when adding new features
- Document API changes in godoc format
- Include examples in code comments
- Keep CHANGELOG.md updated for releases

## Security Considerations
- Use RBAC with minimal required permissions
- Validate all input from external sources
- Use secure defaults in configurations
- Follow Kubernetes security best practices

## Common Commands
```bash
# Generate code after modifying APIs
make generate

# Run tests
make test

# Build and push image
make docker-build docker-push IMG=<registry>/test-operator:tag

# Deploy to cluster
make deploy IMG=<registry>/test-operator:tag

# Run locally against cluster
make run
```

## When suggesting code changes:
1. Ensure compatibility with the current Kubebuilder version
2. Follow the existing project structure and patterns
3. Include proper error handling and logging
4. Consider the impact on the automated update workflow
5. Maintain backward compatibility when possible