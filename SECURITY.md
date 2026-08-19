# Security Policy

This repository is XION's fork of
[`strangelove-ventures/tokenfactory`](https://github.com/strangelove-ventures/tokenfactory),
which the chain node builds against. It is an asset in the
[Blockchain / DLT bug bounty program](https://github.com/burnt-labs/bug-bounty/blob/main/programs/blockchain.md)
— **Burnt Labs' patches only**. This policy is built from that program's terms.
[`burnt-labs/bug-bounty`](https://github.com/burnt-labs/bug-bounty) remains the
canonical source — where this file and the program documents differ, the
program documents govern.

## Reporting a Vulnerability

**Do not open a public GitHub issue for a security vulnerability.** Public
disclosure before a patch is available increases the harm to users.

| Type of finding                  | How to report                                                       |
| -------------------------------- | ------------------------------------------------------------------- |
| Security vulnerability           | **Security → Report a vulnerability** on this repository, or email [security@burnt.com](mailto:security@burnt.com) |
| Non-sensitive or operational bug | Open a GitHub issue on this repository                              |

Prefer GitHub private vulnerability reporting: the fix is developed against the
report, and you are credited on the published advisory and in any CVE we
request.

We acknowledge receipt within **5 business days** and provide a triage decision
within **14 days**. Active exploitation, or confirmed attacker awareness of an
unpatched vulnerability, escalates the issue to Critical response handling —
prioritization, coordination, and disclosure timeline — regardless of its
original classification. That escalation does not change the severity
assessment of the finding or its reward eligibility.

## Fork Scope

This fork ships on mainnet under `-xion.N` version tags, consumed by
[`burnt-labs/xion`](https://github.com/burnt-labs/xion) through a `replace`
directive in its `go.mod`.

**Only the delta between the fork and its upstream base is in scope.** This
fork's version tags do not correspond to upstream release tags — determine the
upstream base as the merge base between the fork tag and upstream `main`
(`git merge-base <fork-tag> upstream/main`). For the current `v0.53.4-xion.N`
tags that base is upstream `v0.50.7-wasmvm2` (commit `dacc993`); diff against
it. A finding that reproduces on the unmodified upstream base belongs to the
upstream project, not to this program, and is not eligible here regardless of
its impact on XION.

Scope applies to the current mainnet release. Findings affecting only
deprecated or end-of-life versions, or already remediated in the currently
deployed mainnet version, are not eligible regardless of whether the fix was
publicly announced. Verify exploitability against the currently deployed
version before submitting. The currently deployed version is recorded in
[`burnt-labs/xion-mainnet-1`](https://github.com/burnt-labs/xion-mainnet-1),
the network's reference configuration.

## Severity

| Severity     | Description                                                                                                                                                                                                                                    |
| ------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **CRITICAL** | Direct, permanent, irrecoverable theft or loss of user funds at protocol scale. Unauthorized minting. Chain halt or consensus failure requiring a hard fork to resolve                                                                          |
| **HIGH**     | Theft or freezing of user funds affecting individual accounts. Significant authentication bypass with demonstrated exploitability                                                                                                              |
| **MEDIUM**   | Limited fund loss or temporary disruption requiring specific preconditions. Attacks requiring privileged-party cooperation                                                                                                                     |
| **LOW**      | Valid, reproducible code-level issue with no direct risk to funds or chain safety, representing a meaningful hardening opportunity. Must include a specific code reference                                                                      |

Only **High** and **Critical** findings are reward eligible. Collecting a
bounty requires completing a KYC process; we cannot pay reporters in
sanctioned jurisdictions. Where several reports describe the same underlying
issue, the first complete report with a working proof of concept is the one
considered. We assess reports as submitted; we do not reclassify a report to a
different severity on a reporter's behalf.

## Proof of Concept

**An end-to-end proof of concept is required.**

Unit tests using `setupKeeper(t)` or similar harnesses bypass transaction
encoding, routing, and the ante handler chain, and do not demonstrate on-chain
exploitability on their own.

The proof of concept should run against a **locally running XION node
configured with mainnet parameters** — with the XION ante handler chain,
module set, and governance configuration matching mainnet — and execute the
attack via standard transaction broadcast (`BroadcastTxSync` or equivalent)
against that node. Broadcast acceptance alone is not sufficient: demonstrate
inclusion in a block and the resulting state change or impact. Simulated
environments that model chain state without running a full node do not
demonstrate exploitability.

## Permissioned Chain Policy

XION mainnet operates with `code_upload_access: Nobody`. Contract deployment
requires a governance proposal. Any attack vector requiring an attacker to
deploy a malicious contract on mainnet is out of scope, regardless of
technical validity.

## Privileged Actor Policy

Attacks requiring a privileged party — governance, a module authority, or a
validator — to take self-destructive or colluding action are classified at
**Medium at most**, regardless of downstream impact. The threat model assumes
privileged actors operate within the specified protocol parameters.

## Out of Scope

**Assets**

- Code that reproduces on the unmodified upstream
  [`strangelove-ventures/tokenfactory`](https://github.com/strangelove-ventures/tokenfactory)
  base — report it upstream
- The chain node and the other modules it builds against — see the
  [Blockchain / DLT program](https://github.com/burnt-labs/bug-bounty/blob/main/programs/blockchain.md)
  for the full asset list
- Smart contracts — see the
  [Core Protocol Contracts program](https://github.com/burnt-labs/bug-bounty/blob/main/programs/contracts.md)
- Third-party infrastructure, RPC providers, and external dependencies
- Public blockchain RPC, REST, gRPC, and Tendermint RPC endpoints

**Vulnerability classes**

- Attacks requiring malicious contract deployment on mainnet
- Denial of service of any form, including single-transaction resource
  exhaustion, node crashes, and chain halts recoverable via a software patch,
  coordinated validator restart, or governance parameter update. Chain halts
  requiring a hard fork to resolve remain in scope under Critical
- Governance attacks requiring a malicious proposal to pass
- Theoretical vulnerabilities without a working end-to-end proof of concept
- Attacks where the attacker's cost to execute exceeds the demonstrable harm to
  the protocol or its users
- Best practices, gas optimizations, missing events, and informational findings

## Responsible Disclosure

- Do not exploit a vulnerability beyond what is necessary to confirm it exists
- **Do not test against production systems.** This includes XION mainnet.
  Testing production disqualifies the report
- Use only a local environment or infrastructure you control. Other XION
  testnets, staging, and development deployments are not authorized by
  implication
- Do not access, modify, or exfiltrate user data
- Do not disrupt or degrade our networks, data, or services
- Do not disclose publicly before a fix is confirmed and deployed
- Allow us reasonable time to address the issue

## Safe Harbor

Burnt Labs will not pursue legal action against researchers who report
vulnerabilities in good faith under this policy, do not exploit beyond what is
necessary to confirm the finding, do not access or disclose user data, and do
not disrupt production systems.

Authorization to actively test extends only to local environments and
infrastructure you control. Testing against production systems — including
XION mainnet — is not authorized under this policy. Reporting a vulnerability
you encountered incidentally is always welcome.
