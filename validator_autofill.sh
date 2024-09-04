#!/bin/bash

set -e

echo 'populating validator settings from environment values...';

if [ -z "$SIGNER_ACCOUNT_ADDRESS" ]; then
    echo "SIGNER_ACCOUNT_ADDRESS not found, please set this in your .env!";
    exit 1;
fi

if [ -z "$SIGNER_ACCOUNT_PRIVATE_KEY" ]; then
    echo "SIGNER_ACCOUNT_PRIVATE_KEY not found, please set this in your .env!";
    exit 1;
fi


if [ -z "$PROST_RPC_URL" ]; then
    echo "$PROST_RPC_URL not found, please set this in your .env!";
    exit 1;
fi
if [ -z "$PROTOCOL_STATE_CONTRACT" ]; then
    echo "PROTOCOL_STATE_CONTRACT not found, please set this in your .env!";
    exit 1;
fi

# Assuming default values for each variable if not provided
export REDIS_HOST="${REDIS_HOST:-redis}"
export REDIS_PORT="${REDIS_PORT:-6379}"
export IPFS_URL="${IPFS_URL:-/dns/ipfs/tcp/5001}"
export IPFS_API_KEY="${IPFS_API_KEY:-}"
export IPFS_API_SECRET="${IPFS_API_SECRET:-}"
export BATCH_SUBMISSION_LIMIT="${BATCH_SUBMISSION_LIMIT:-60}"
export BLOCK_TIME="${BLOCK_TIME:-1}"

cd config

# Template to actual settings.json manipulation
cp settings.example.json settings.json

# Replace placeholders in settings.json with actual values from environment variables
sed -i'.backup' -e "s#PROST_RPC_URL#$PROST_RPC_URL#" \
                -e "s#PROTOCOL_STATE_CONTRACT#$PROTOCOL_STATE_CONTRACT#" \
                -e "s#REDIS_HOST#$REDIS_HOST#" \
                -e "s#REDIS_PORT#$REDIS_PORT#" \
                -e "s#IPFS_URL#$IPFS_URL#" \
                -e "s#SIGNER_ACCOUNT_ADDRESS#$SIGNER_ACCOUNT_ADDRESS#" \
                -e "s#SIGNER_ACCOUNT_PRIVATE_KEY#$SIGNER_ACCOUNT_PRIVATE_KEY#" \
                -e "s#PROST_CHAIN_ID#$PROST_CHAIN_ID#" settings.json \
                -e "s#BLOCK_TIME#$BLOCK_TIME#" \
                -e "s#BATCH_SUBMISSION_LIMIT#$BATCH_SUBMISSION_LIMIT#" settings.json

# Cleanup backup file
rm settings.json.backup
