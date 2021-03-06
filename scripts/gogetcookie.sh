#!/bin/bash
eval 'set +o history' 2>/dev/null || setopt HIST_IGNORE_SPACE 2>/dev/null
touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
go.googlesource.com,FALSE,/,TRUE,2147483647,o,git-pdecat.gmail.com=1//03yv4I3o2PgEvCgYIARAAGAMSNwF-L9IrRC8k3SSpICt8eGnOaQu48g5hUG4zBSLLMJW0hEg-y_6UkI4T05JPhiONx5u5X_oeVaI
go-review.googlesource.com,FALSE,/,TRUE,2147483647,o,git-pdecat.gmail.com=1//03yv4I3o2PgEvCgYIARAAGAMSNwF-L9IrRC8k3SSpICt8eGnOaQu48g5hUG4zBSLLMJW0hEg-y_6UkI4T05JPhiONx5u5X_oeVaI
__END__
eval 'set -o history' 2>/dev/null || unsetopt HIST_IGNORE_SPACE 2>/dev/null
