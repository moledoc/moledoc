export PLAN9=$HOME/plan9port
export PATH=$HOME/go/bin:$HOME/venv/bin:/usr/local/bin:/opt/homebrew/bin:/usr/local/go/bin:$PLAN9/bin:$PATH

alias acme="open -a acme.app"
function git_branch_name() {
	branch=$(git branch --show-current 2> /dev/null)
	test -n "$branch" && echo "($branch) "
}

update_prompt() {
    PS1="%d $(git_branch_name)%% "
}
precmd_functions+=(update_prompt)
update_prompt

