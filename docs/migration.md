# Migration

When upgrading to a new major version, follow these steps:

1. Run `buf breaking` to detect breaking changes.
2. Update dependent services accordingly.
3. Regenerate code with `buf generate`.
