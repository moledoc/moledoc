# git

Here are some more common command line commands for git.

To get additional information about git command or flags, we use the command **help**.

```sh
# To see available git commands
git help
# To see detailed information about given command
git help <command>
```

To clone existing repository, we use the command **clone**.

```sh
# clone existing remote repository to local computer
git clone <remote repository location>
```

Another option is to create an local git repository and then set the remote repository manually.

```sh
# initialize local git repository
git init
# add remote repository to initialized local git repository
git remote add origin <path to remote repository> 
# Example 
# * https -- https://github.com/moledoc/moledoc.git
# * ssh -- git@github.com:moledoc/moledoc.git
# pull/push data and set main/master branch as the upstream (where files/directories are pulled/pushed from/to)
git pull/push --set-upstream origin main

# remove remote, if the path was inserted incorrectly or it does not exist anymore
git remote rm origin
```

Remote name does not have to be *origin*, it could be something else, but typically it is set as *origin*.
Furthermore, one local repository can be linked to multiple remote repositories by defining multiple remotes. For example

```sh
git remote add origin <path to first git repo>
git remote add secondary <path to second git repo>
```

**NB!** This will make the two repos mirrored.

You can change the remote with

```sh
git remote set-url origin git@github.com:<username>/<repo>.git # to ssh
git remote set-url origin https://github.com/<username>/<repo>.git # to https
```

To configure git behaviour, we use the command **config**.
This can either be `--local` (in that specific repo) or `--global` (applies to all local repos).
Useful examples:

```sh
# username used in git commits
git config --global user.name "<username>"

# email used in git commits
git config --global user.email "<email>"

# changing git commit editor, where option is a text editor (ie vim, nano, but can be graphical as well)
git config --global core.editor "<option>"
git config --local core.editor "vim"

# setting git command aliases.
git config --local alias.co checkout # usages: git co
```

When having multiple branches, then it is recommended to use git **worktree** instead.
In that case each branch gets its own dedicated directory and switching between branches is less confusing.
Only negative point is that git worktree branch doesn't show how many commits ahead/behind your local repository is compared to the remote repository.

Easiest way to set up git worktree is with clone command.

```sh
# git worktree using clone command
git clone --bare <remote repository location> <directory name where the repository is cloned>
cd <directory name where the repository is cloned>
# add necessary branches
git worktree add master
git worktree add feature1
git worktree add feature2
```

When we want to add files to remote repository, we need to describe, which files we want to commit.
For that we stage the files with the command **add**.

```sh
# stage file in the directory where the terminal is currently located
git add <file>
# stage file in the given path
git add <path to file>
# stage everything in current directory
git add .
# stage everything in parent directory
git add ..

# add specific chunks from given files
git add -p <file(s)>
```

To check the status of local repository, we use the command **status**.

```sh
# check local repository status
git status
# check local repository status from a specific directory
git status -- <path/to/dir>
```

To commit the changes in git, we use the command **commit**.

```sh
# commit staged changes (this opens up a commit message buffer in set text editor, see git config core.editor)
git commit

# commit staged changes with a message
git commit -m "<here goes commit message>"

# commit all unstaged files with a message
git commit -am "<here goes commit message>"
```

To pull files or changes from remote to local repository, we use the command **pull**.
To push the changes to remote repository, we use the command **push**.

```sh
# pull remote repository changes to local repository
git pull
# push local committed changes to remote repository
git push
```

Let's say we have changed something in a file, but we want to have the remote repository version of the file in our local repository.
Then we can use the command **checkout** to get the remote repository version of the file to local repository.

```sh
# get remote repository version of a file to local repository
git checkout -- <filename>
```

To bring the latest metadata from remote to local repository, we use the command **fetch**.

```sh
git fetch
```

To make a new branch, view branches, switch to that branch, remove/delete that branch in local and remote repository, we use the command **branch**.
Before switching branch (not relevant for git-worktree), then it is recommended to **stash** the changes of the current branch, so the changes would not get committed in the wrong branch.

```sh
# make new branch
git branch <new branch>
# view all existing branches
git branch -a
# view existing branches in local repository
git branch -l
# rename branch
git branch -m <old name> <new name>

# switch branches
git checkout <branch name>
## make a new branch and switch to it
git checkout -b <new branch>
# stash current branch changes before checkouting another branch
git stash
# stash current branch changes with message/name
git stash push -m "message"
# list stash
git stash list
# show changes in stash
git stash show
# unstash (pop) latest stashed changes
git stash pop
# unstash n'th stash (by list index - indeces start from 0)
git stash pop n
git stash pop stash@{n}
# unstash a stash by name
git stash pop stash^{/my_stash_name}


# delete local branch 
git branch -d <branch name>
# delete local branch with commits
git branch -D <branch name>
# delete remote branch
git push origin --delete <remote-branch-name>
```

When using git-worktree, we can use the following commands to work with branches:

```sh
# make new branch
git worktree add <new branch>
# view all worktree branches
git worktree list
# switch branches
* change path to corresponding branch directory
# delete worktree branch
git worktree remove <branch name>
git worktree prune # to clean up any stale administrative files
```

To merge branches, we use the command **merge**.
It is highly recommended to first merge master/main branch to other branch, before merging the other branch into master.
That is because then we get the latest master branch files into our other branch, avoiding possible merge conflicts.
Furthermore, if any merge conflict does occur, then it happens in the other branch, not in the master branch.

```sh
# switch/navigate to the other branch

## in case of git-worktree, just open the corresponding branch directory
git checkout <other branch>
# merge master into the other branch
git merge master
# just in case push changes in other branch
git push
# switch/navigate to the master/main branch
## in case of git-worktree, just open the master/main branch directory
git checkout master
# merge other branch to master/main and push the changes to remote repository
git merge <other branch>
git push
```

To see changes and history, we use the commands **log** or **diff**.

```sh
# show current branch commit log
git log # q to quit out of the log
# show current branch changes
git diff # q to quit out of the diff

# create patch text (diff) for current branch (shows diff's for previous commits)
git log -p
# create patch text (diff) for specific file in the current branch
git log -p -- <filename>
# show changes for specific file in the current branch 
git diff -- <filename>

# show given branch commit log
git log <branch name>
# show given branch changes
git diff <branch name>

# create patch text (diff) for given branch (shows diff's for previous commits)
git log -p <branch name>
# create patch text (diff) for specific file in the given branch
git log -p <branch name> -- <filename>
# show changes for specific file in the given branch 
git diff <branch name> -- <filename>
```

To see, which remote branch is being tracked by local branch, we use the command **remote**.

```sh
git remote show <remote name> # remote name is usually 'origin', but one repository can have multiple remotes

# to see which remotes are connected to the local repository
git remote  show
```

## Author

Written by
Meelis Utt
