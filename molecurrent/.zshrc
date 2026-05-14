export PLAN9=$HOME/plan9port
export PATH=$PATH:/opt/homebrew/bin:/usr/local/go/bin:$HOME/go/bin:$PLAN9/bin

alias acme="acme -f /mnt/font/Menlo-Regular/14a/font"
# alias lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}';ACME_LSP_CONFIG=$HOME/.config/acme-lsp/config.toml acme-lsp -hidediag &"
# alias kill-lsp="pgrep \"acme-lsp\" | parallel 'kill -9 {}'"

function git_branch_name() {
	branch=$(git branch --show-current 2> /dev/null)
	test -n "$branch" && echo "($branch) "
}

update_prompt() {
    PS1="%d $(git_branch_name)%% "
}
precmd_functions+=(update_prompt)
update_prompt

