#!/bin/bash

./app_env.sh

$image_name = "$HUB_REPO/$APP_NAME:$APP_BRANCH"
echo "Deploying $image_name"

docker push $HUB_REPO/$APP_NAME:$APP_BRANCH
