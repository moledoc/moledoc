# CLI (Command line interface) Git

Here are some more common command line commands for git.

To clone existing repository, we use the command clone.

```sh
# clone existing remote repository to local computer
git clone <remote repository location>
```

Another option is to create an local git repository and then set the remote repository manually.

```sh
# initialize local git repository
git init
# add remote repository to initialized local git repository
git remote add origin <path to remote repository> # eg git@github.com:moledoc/moledoc.git
# remove remote, if the path was inserted incorrectly or it does not exist anymore
git remote rm origin
# pull/push data and set main/master branch as the upstream (where files/directories are pulled/pushed from/to)
git pull/push --set-upstream origin main
```

When having multiple branches, then it is recommended to use git worktree instead. That is because in that case each branch gets its own dedicated directory and switching between branches is less confusing. Easiest to set up git worktree is with clone command.

```sh
# git worktree using clone command
git clone --bare <remote repository location> <directory name where the repository is cloned>
cd <directory name where the repository is cloned>
# add necessary branches
git worktree add master
git worktree add feature1
git worktree add feature2
```

We must describe to git, which files we want to commit. For that we stage the files with the command add.

```sh
# stage file in the directory where the terminal is currently located
git add <file>
# stage file in the given path
git add <path to file>
# stage everything in current directory
git add .
# stage everything in parent directory
git add ..
```

To check the status of local repository, we use the command status

```sh
# check local repository status
git status
```

To commit the changes in git, we use the command commit.

```sh
# commit staged changes with a message
git commit -m "<here goes commit message>"
```

To pull files or changes from remote to local repository, we use the command pull. To push the changes to remote repository, we use the command push.

```sh
# pull remote repository changes to local repository
git pull
# push local committed changes to remote repository
git push
```

Let's say we have changed something in a file, but we want to have the remote repository version of the file in our local repository. Then we can use the command checkout to get the remote repository version of the file to local repository.

```sh
# get remote repository version of a file to local repository
git checkout -- <filename>
```

To make a new branch, view branches, switch to that branch, remove/delete that branch in local and remote repository

```sh
# IN CASE OF NON GIT WORKTREE
# make new branch
git branch <new branch>
# view all existing branches
git branch -a
# view existing branches in local repository
git branch -l
# switch branches
git checkout <branch name>
# delete local branch 
git branch -d <branch name>
# delete local branch with commits
git branch -D <branch name>
# delete remote branch
git push origin --delete <remote-branch-name>
```

```sh
# IN CASE OF GIT WORKTREE
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

To merge branches, we use the command merge. It is highly recommended to first merge master/main branch to other branch, before merging the other branch into master. That is because then we get the latest master branch files into our other branch, avoiding possible merge conflicts. Furthermore, if any merge conflict does occur, then it happens in the other branch, not in the master branch.

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

To see git log, we use the command log. To see what changes have been made in the local repository, we use the command diff.

```sh
# show commit log
git log # q to quit out of the log
# show changes in the local repository
git diff # q to quit out of the diff
```