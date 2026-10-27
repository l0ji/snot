# snot

static web server to put on top of your tangled knot

## run

```sh
go build
./snot
```

to strap on top of your knot you should probably use a reverse proxy with everything but the XRPC api sent to this
server

## checklist

### git

* show list of repositories per-owner
* have knowledge of all branches in a repo
* show file trees for branches and individual commits
* show commit history of all branches
* show file contents in a commit
* show diffs between commits and branches

### appearance

* can get basic user info
* can validate commit signatures
* shows README + cloning guide on repo summary
* auto-link PRs and issues to

### server

* literally render everything statically
* there should be no server-side dynamic behavior
* client-side dynamic behavior is fine tho
* configure with an env file because i don't feel like downloading a toml parser