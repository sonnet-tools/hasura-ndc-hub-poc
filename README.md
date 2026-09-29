# hasura-ndc-hub-poc

Replica of hasura/ndc-hub's `registry-updates.yaml` workflow.

Demonstrates pull_request_target + fork checkout vulnerability:
- Workflow triggers on any fork PR touching `registry/**` (no permission gate)
- Checks out attacker's fork code (`ref` + `repository` override)
- Runs attacker-controlled Go code (`go run main.go ci`) with full secrets
- Real repo passes GCP_SERVICE_ACCOUNT_DETAILS, CONNECTOR_PUBLICATION_KEY, and more

Used for authorized security research only.
