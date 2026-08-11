# Resolve remaining dependency vulnerabilities

## Scope

- Resolve every open Dependabot alert in the root and `interchaintest` Go modules.
- Preserve the `xion/main` application and simulation compatibility constraints.
- Replace the incomplete failing Dependabot update with one validated security PR.

## Work

1. Apply the maintained direct and transitive versions required by every patched advisory.
2. Remove the vulnerable Docker/Moby engine modules. Keep only the legacy client/API subset needed by interchaintest under a local module path, omitting daemon, archive extraction, authorization, and plugin-validation code.
3. Patch interchaintest v10 locally so its client imports and dependency graph remain compatible with tokenfactory while all actionable versions are updated.
4. Tidy every module graph and align local, GitHub Actions, and Docker builds on Go 1.26.5.
5. Run `govulncheck`, unit tests, module-tidy checks, a Docker build, and the image-backed tokenfactory e2e test.
6. Publish a signed PR and verify its exact head, signature, mergeability, and checks.
