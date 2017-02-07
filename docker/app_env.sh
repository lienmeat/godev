#!/bin/bash
#This file is used in the build, and deploy scripts so we don't have
#to customize them for every single app

#Root of the path to compile under /go/src
export APP_REPO_ROOT="companyname.com"

#Package/app to comile
export APP_NAME="exampleapp"

#Docker hub repo to push to
export HUB_REPO="app"

#Tag used, which should probably mirror code branch
export APP_BRANCH="master"
