#!/bin/sh

acme -f /mnt/font/DejaVuSansMono/14a/font -m /mnt/acme &
# pgrep "acme-lsp" | parallel 'kill -9 {}' && ACME_LSP_CONFIG=$HOME/.config/acme-lsp/config.toml acme-lsp -hidediag &
