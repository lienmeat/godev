#!/bin/bash

docker-compose -f docker-compose-test.yml rm -f && docker-compose -f docker-compose-test.yml up --force-recreate