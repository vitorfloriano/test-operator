# test-operator - Kubernetes WordPress Operator

test-operator is a Kubernetes operator built with Kubebuilder v4.6.0 that manages WordPress custom resources. It provides automation for deploying and managing WordPress instances on Kubernetes clusters.

**ALWAYS reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Initial Setup and Dependencies
- Prerequisites: Go 1.23.0+, Docker 17.03+, kubectl 1.11.3+, Kubernetes cluster access
- Bootstrap the repository:
  - `go mod tidy` -- downloads dependencies, takes ~40 seconds on first run
  - `go mod download` -- ensures all dependencies are cached, takes ~30 seconds

### Build Commands
- Generate and build the operator:
  - `make manifests` -- generates CRDs and RBAC manifests. FIRST RUN: ~50 seconds (downloads controller-gen). Subsequent runs: ~1 second. NEVER CANCEL.
  - `make generate` -- generates DeepCopy code, takes ~1 second
  - `make fmt` -- formats Go code, takes <1 second
  - `make vet` -- vets Go code, takes ~1 minute. NEVER CANCEL.
  - `make build` -- builds manager binary, takes ~4 seconds (includes all above steps)

### Testing Commands
- Run tests:
  - `make test` -- runs unit tests, takes ~38 seconds (includes envtest setup). NEVER CANCEL. Set timeout to 120+ seconds.
  - `make lint` -- runs golangci-lint. FIRST RUN: ~3+ minutes (downloads linter). Subsequent runs: ~2 seconds. NEVER CANCEL on first run.
  - `make lint-fix` -- automatically fixes linting issues, takes ~1 second
  - `make lint-config` -- verifies linter configuration

### Docker and Container Operations
- **WARNING**: Docker builds fail in sandboxed environments due to certificate verification issues with Go module proxy
- `make docker-build IMG=<registry>/test-operator:tag` -- builds Docker image. EXPECT FAILURE in sandbox environments.
- `make docker-push IMG=<registry>/test-operator:tag` -- pushes Docker image
- Use `make build` instead to create local binary for testing

### Development and Local Testing
- Run the operator locally:
  - `make run` -- runs controller from host (requires cluster access)
  - `./bin/manager --help` -- shows all manager command options
  - `./bin/manager` -- runs manager (will fail without cluster access, this is expected)

### Kubernetes Deployment
- Deploy to cluster:
  - `make install` -- installs CRDs to cluster, takes ~34 seconds (includes kustomize download). NEVER CANCEL.
  - `make deploy IMG=<registry>/test-operator:tag` -- deploys controller to cluster
  - `kubectl apply -k config/samples/` -- creates sample WordPress resources
  - `make uninstall` -- removes CRDs from cluster
  - `make undeploy` -- removes controller from cluster

### End-to-End Testing
- E2E tests with Kind:
  - Requires Kind cluster: `kind create cluster` (takes ~20-30 seconds)
  - `make test-e2e` -- runs e2e tests. **KNOWN ISSUE**: Fails in sandbox due to Docker build dependency
  - Alternative: Use `make install` + `kubectl apply -k config/samples/` to test CRDs manually

## Validation

### Always Test These Workflows
- **Basic build validation**: `make build && make test && make lint`
- **CRD functionality**: `make install && kubectl apply -k config/samples/ && kubectl get wordpresses`
- **Resource cleanup**: `kubectl delete -k config/samples/ && make uninstall`

### Expected Timing (NEVER CANCEL these operations)
- First-time setup (dependencies + tools): ~3-5 minutes total
- Regular build cycle: `make build` (~4s) + `make test` (~38s) + `make lint` (~2s) = ~45 seconds
- CRD installation: ~34 seconds
- Kind cluster creation: ~20-30 seconds

### Manual Validation Steps
After making changes, always:
1. Run `make build` to ensure compilation succeeds
2. Run `make test` to verify unit tests pass  
3. Run `make lint` to check code quality
4. Test CRD functionality with sample resources
5. Verify the manager binary starts with `./bin/manager --help`

## Common Tasks

### Repository Structure
```
.
├── api/v1/                 # WordPress CRD definitions
├── cmd/main.go             # Manager entry point
├── config/                 # Kubernetes manifests and kustomization
│   ├── crd/               # Generated CRDs
│   ├── default/           # Default deployment config
│   └── samples/           # Sample WordPress resources
├── internal/
│   ├── controller/        # WordPress controller logic
│   └── webhook/           # Validation and defaulting webhooks
├── test/
│   ├── e2e/              # End-to-end tests
│   └── utils/            # Test utilities
├── Makefile              # All build and development commands
└── PROJECT               # Kubebuilder project metadata
```

### Key Files for Development
- `api/v1/wordpress_types.go` -- WordPress CRD spec and status definitions
- `internal/controller/wordpress_controller.go` -- Main reconciliation logic
- `config/samples/example.com_v1_wordpress.yaml` -- Sample WordPress resource
- `Makefile` -- All validated commands and their dependencies
- `.golangci.yml` -- Linting configuration (affects CI)

### Working with WordPress Resources
```yaml
# Sample WordPress resource structure
apiVersion: example.com.my.domain/v1
kind: Wordpress
metadata:
  name: wordpress-sample
spec:
  size: 1  # Number of instances (1-3)
```

### CI/CD Integration
- GitHub Actions workflows: `.github/workflows/test.yml`, `.github/workflows/lint.yml`, `.github/workflows/test-e2e.yml`
- Always run `make lint` before committing (CI will fail on lint errors)
- Tests run on every push and pull request
- E2E tests run in CI with Kind cluster setup

## Known Limitations and Workarounds

### Docker Build Issues
- **Problem**: Docker builds fail in sandboxed environments with certificate verification errors
- **Workaround**: Use `make build` to create local binary, test functionality with `make install` + sample resources
- **Production**: Docker builds work in normal environments with internet access

### E2E Test Dependencies
- **Problem**: E2E tests require Docker image build, which fails in sandbox
- **Workaround**: Test operator functionality using:
  1. `make install` (install CRDs)
  2. `kubectl apply -k config/samples/` (create resources)
  3. `kubectl get wordpresses` (verify resources)
  4. Manual controller testing with `./bin/manager --help`

### Environment Requirements
- Requires active Kubernetes cluster for `make deploy` and `make run`
- Use Kind for local testing: `kind create cluster`
- Manager binary can run outside cluster with kubeconfig

## Quick Reference Commands

**Daily Development:**
```bash
# Full validation cycle
make build && make test && make lint

# Test with local cluster
kind create cluster
make install
kubectl apply -k config/samples/
kubectl get wordpresses
```

**Release Preparation:**
```bash
# Generate all artifacts
make manifests generate build-installer

# Validate everything works
make test lint build
```

**Troubleshooting:**
```bash
# Clean and rebuild
make clean  # (if available)
go mod tidy
make build

# Lint fixes
make lint-fix

# Reset cluster state
make uninstall || true
make install
```