#!/bin/bash
# chmod +x ./update.sh
# ./update.sh $1

if [ $# -eq 1 ]
then
    commitMsg=$1
else
    echo $#
    warning="Errorly exited.\nPlease input 1 parameters:\n\$1-commit message. \n"
    echo -e $warning
    exit 0
fi


git_usrname="mashkyt"
git_email="ashikay@outlook.com"

git config --global user.name $git_usrname
git config --global user.email $git_email

git pull
git add .
git commit -m "Added the latest. ${commitMsg}."

git push
