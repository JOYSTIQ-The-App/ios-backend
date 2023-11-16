#!/bin/bash

# Fetch all remote branches and tags
git fetch --all

# Prune stale references
git fetch -p

# Delete local branches that correspond to deleted remote branches
git branch -vv | grep ': gone]' | awk '{print $1}' | xargs git branch -d

echo "Branch synchronization complete."
