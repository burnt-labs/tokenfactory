# Client-only Docker compatibility module

This local module contains the Docker v28.5.2 client, API types, error types,
and stream helpers under the repository's own module path. The original source
is Apache-2.0 licensed; `LICENSE`, `NOTICE`, and `AUTHORS` are retained here.

Tokenfactory's interchain tests need the v28 client interface, but the upstream
`github.com/docker/docker` and `github.com/moby/moby` modules also include the
Docker daemon. Current advisories affect daemon-side archive extraction, AuthZ,
and plugin validation. This compatibility module intentionally omits all daemon,
archive extraction, authorization plugin, and plugin validation code. It is not
a Docker engine fork.

Imports are renamed to this local module path so the vulnerable full-engine
modules are absent from the build graph. This module can be removed when
interchaintest supports Moby's v29 split client API.
