# export PLAN9=/usr/local/plan9
# export PATH=$PLAN9/bin:/usr/local/go/bin:$HOME/go/bin:/opt/local/bin:/opt/homebrew/bin:$PATH
export PATH=/usr/local/bin:/usr/local/go/bin:$HOME/go/bin:/opt/local/bin:/opt/homebrew/bin:$PATH

alias acme="acme -f /mnt/font/Menlo-Regular/14a/font"
alias lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}';ACME_LSP_CONFIG=$HOME/.config/acme-lsp/config.toml acme-lsp -hidediag &"
alias kill-lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}'"
alias postgres="LC_ALL="C" /opt/homebrew/opt/postgresql@17/bin/postgres -D /opt/homebrew/var/postgresql@17"

function git_branch_name() {
	branch=$(git symbolic-ref HEAD 2> /dev/null | awk 'BEGIN{FS="/"} {print $NF}')
	test -n "$branch" && echo "($branch) "
}

PS1='$PWD $(git_branch_name)% '
