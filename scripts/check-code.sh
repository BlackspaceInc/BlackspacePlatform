#!/bin/bash
set -euo pipefail

# check-code.sh
#
# SUMMARY
#
#   Checks all Blackspacepace code

cargo check --workspace --all-targets
