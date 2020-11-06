# Sync Fork

## Fork to a repo

- For example want to for with `todo-api` repo into your local.
- Open `https://github.com/pk80/todo-api.git` and click `fork` at the top righ corner.
- This forks `https://github.com/pk80/todo-api.git` into your local repo.
- Clone `https://github.com/pk80/todo-api.git` to your local by cloning.

## Configure Git to sync your fork with the original todo-api repository

- After cloning the forked `todo-api` repo into your local, navigate to `todo-api` folder.
- Follow below steps:

```cmd
$ git remote -v
origin  https://github.com/pk80/todo-api.git (fetch)
origin  https://github.com/pk80/todo-api.git (push)
```

- Now, add upstream

```cmd
$ git remote add upstream https://github.com/muly/todo-api.git
$ git remote -v
origin  https://github.com/pk80/todo-api.git (fetch)
origin  https://github.com/pk80/todo-api.git (push)
upstream        https://github.com/muly/todo-api.git (fetch)
upstream        https://github.com/muly/todo-api.git (push)
```

- Now, you can keep your fork synced with the upstream repository with a few Git commands.

## Syncing a fork

Resource: <https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/syncing-a-fork>  

```cmd
$ git fetch upstream
From https://github.com/muly/todo-api
 * [new branch]      httprouter-test -> upstream/httprouter-test
 * [new branch]      master          -> upstream/master
```

- Check out your fork's local default branch

```cmd
$ git checkout master
> Switched to branch 'master'
```

- Now, merge the changes from the upstream default branch upstream/master - into - your local default branch.

```cmd
$ git merge upstream/master
```

## Create a new branch

```cmd
$ git checkout -b pk
> Switched to a new branch 'pk'
$ git push origin pk:master
> Everything up-to-date
```
