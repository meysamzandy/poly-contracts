#!/bin/bash

set -e


buf lint

buf generate


echo "Generated successfully"