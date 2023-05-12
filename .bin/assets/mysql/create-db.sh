#!/bin/env bash
for f in /docker-entrypoint-initdb.d/*.sql
  do
   echo "Processing $f file..."
   mysql -u root "-p123456" < "$f" 2>&1 | grep -v "Using a password" || true
  done