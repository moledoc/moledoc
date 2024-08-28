export PLAN9=/usr/local/plan9
export PATH=$PATH:$PLAN9/bin:/usr/local/go/bin:$HOME/go/bin:/opt/local/bin

alias acme="acme -f /mnt/font/Menlo-Regular/14a/font"
alias lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}';ACME_LSP_CONFIG=$HOME/.config/acme-lsp/config.toml acme-lsp -hidediag &"
alias kill-lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}'"

function git_branch_name() {
	branch=$(git symbolic-ref HEAD 2> /dev/null | awk 'BEGIN{FS="/"} {print $NF}')
	test -n "$branch" && echo "($branch) "
}

PS1='$PWD $(git_branch_name)% '
