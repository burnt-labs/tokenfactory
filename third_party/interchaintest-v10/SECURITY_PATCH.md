# Tokenfactory security patch

This directory contains `github.com/cosmos/interchaintest/v10` v10.0.1 under
its original Apache-2.0 license.

The local patch updates vulnerable dependencies and redirects Docker/Moby
client imports to tokenfactory's client-only compatibility module. No Docker
daemon code is included in the resulting module graph.

Only packages required by tokenfactory's compiled interchain test graph and
their package-local tests are retained; unrelated upstream examples, tooling,
documentation, and workflows are omitted.
