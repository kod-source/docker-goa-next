#!/usr/bin/env bash

set -euo pipefail

if [ $# -ne 2 ]; then
    echo "実行するに2個の引数が必要です。" 1>&2
    exit 1
fi

EMAIL=$1
PASSWORD=$2

# JWTのトークン取得
TOKEN=$(curl -X POST -H "Content-Type: application/json" -d '{"email":"'$EMAIL'", "password":"'$PASSWORD'"}' http://localhost:3000/login | jq -r '.token')

if [ -z "$TOKEN" ] || [ "$TOKEN" == null ]; then
    echo "emailかpasswordが間違っております。" 1>&2
    exit 1
fi

echo -n "$TOKEN"
