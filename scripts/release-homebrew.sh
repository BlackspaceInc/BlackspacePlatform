#!/usr/bin/env bash
set -euo pipefail

# release-homebrew.sh
#
# SUMMARY
#
#   Releases latest version to the timberio homebrew tap

td="$(mktemp -d)"
pushd "$td"

git config --global user.email "bradybot@timber.io"
git config --global user.name "bradybot"

git clone "https://$GITHUB_TOKEN:x-oauth-basic@github.com/timberio/homebrew-brew"
cd homebrew-brew

PACKAGE_URL="https://packages.timber.io/Blackspace/$VERSION/Blackspace-x86_64-apple-darwin.tar.gz"
PACKAGE_SHA256=$(curl -s "$PACKAGE_URL" | sha256sum | cut -d " " -f 1)

update-content() {
  sed "s|url \".*\"|url \"$PACKAGE_URL\"|" \
    | sed "s|sha256 \".*\"|sha256 \"$PACKAGE_SHA256\"|" \
    | sed "s|version \".*\"|version \"$VERSION\"|"
}

NEW_CONTENT="$(update-content < Formula/Blackspace.rb)"

echo "$NEW_CONTENT" > Formula/Blackspace.rb

git commit -am "Release Blackspace $VERSION"
git push

popd
rm -rf "$td"
