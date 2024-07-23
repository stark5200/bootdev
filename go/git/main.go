/*

PORCELAIN AND PLUMBING
In Git, commands are divided into high-level ("porcelain") commands and low-level ("plumbing") commands. The porcelain commands are the ones that you will use most often as a developer to interact with your code. Some porcelain commands are:

git status
git add
git commit
git push
git pull
git log
Don't worry about what they do yet, we'll cover them in detail soon. Some examples of plumbing commands are:

git apply
git commit-tree
git hash-object
We'll focus on the high-level commands because that's what you use 99% of the time as a developer, but we'll dip down into the low-level commands occasionally to really understand how Git works.

// how to view git config file 
cat ~/.gitconfig

//
git cat-file -p 6ee62cb5ceef467b5a3f0929bc2b760f89efa4e9
tree 5b21d4f16a4b07a6cde5a3242187f6a5a68b060f
author stark5200 <a.taha5200@gmail.com> 1721674649 +0300
committer stark5200 <a.taha5200@gmail.com> 1721674683 +0300

A: add contents.md
stark@Ammars-Laptop webflyx % 
//

*/