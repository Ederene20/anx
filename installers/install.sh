#!/bin/sh
# Based on Deno installer: Copyright 2019 the Deno authors. All rights reserved. MIT license.
# TODO(everyone): Keep this script simple and easily auditable.

set -e

version=0.2.2

anx_uri="https://github.com/Ederene20/anx/releases/download/v$version/anx_$version.tar.gz"

anx_install="${ANX_INSTALL:-$HOME/.anx}"

bin_dir="$anx_install/bin"
exe="$bin_dir/anx"

if [ ! -d "$bin_dir" ]; then
    mkdir -p "$bin_dir"
fi

curl -q --fail --location --progress-bar --output "$exe.tar.gz" "$anx_uri"

cd "$bin_dir"
tar xzf "$exe.tar.gz"
chmod +x "$exe"
rm "$exe.tar.gz"

echo "anx was installed successfully to $exe"

if command -v anx >/dev/null; then
    echo "Run 'anx --help' to get started"
else
    case $SHELL in
    /bin/bash) shell_profile=".bashrc" ;;
    /bin/zsh) shell_profile=".zshrc" ;;
    *) shell_profile=".bash_profile" ;;
    esac
    echo "Manually add the directory to your \$HOME/$shell_profile (or similar)"
    echo " export ANX_INSTALL='$anx_install' "
    echo " export PATH=\$ANX_INSTALL/bin:\$PATH "
    echo "Run '$exe --help' to get started"
fi
