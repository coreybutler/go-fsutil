#!/bin/sh
rm -rf ./.data
richgo test ./... -v
rm -rf ./.data
