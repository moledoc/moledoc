export PLAN9=$HOME/plan9port
export PATH=$PATH:/opt/homebrew/bin:/usr/local/go/bin:$HOME/go/bin:$PLAN9/bin

alias acme="acme -f /mnt/font/Menlo-Regular/14a/font"
# alias lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}';ACME_LSP_CONFIG=$HOME/.config/acme-lsp/config.toml acme-lsp -hidediag &"
# alias kill-lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}'"
alias postgres="LC_ALL="C" /opt/homebrew/opt/postgresql@17/bin/postgres -D /opt/homebrew/var/postgresql@17"

function git_branch_name() {
	branch=$(git branch --show-current 2> /dev/null)
	test -n "$branch" && echo "($branch) "
}

PS1='$PWD $(git_branch_name)% '
