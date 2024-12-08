#!/bin/bash

# Base URL of the API
BASE_URL="http://localhost:9999"

# Test data
KEY1="testKey1"
VALUE1="testValue1"
KEY2="testKey2"
VALUE2="testValue2"

echo "Starting tests..."

# 1. Set operation: Add key-value pairs
echo "Testing Set operation..."
curl -X POST "$BASE_URL/set" \
     -H "Content-Type: application/json" \
     -d "{\"key\": \"$KEY1\", \"value\": \"$VALUE1\"}" -s -o /dev/null -w "HTTP Code: %{http_code}\n"

curl -X POST "$BASE_URL/set" \
     -H "Content-Type: application/json" \
     -d "{\"key\": \"$KEY2\", \"value\": \"$VALUE2\"}" -s -o /dev/null -w "HTTP Code: %{http_code}\n"

# 2. Get operation: Retrieve a value by key
echo "Testing Get operation..."
RESPONSE=$(curl -X GET "$BASE_URL/get?key=$KEY1" -s)
echo "Get response for $KEY1: $RESPONSE"

RESPONSE=$(curl -X GET "$BASE_URL/get?key=$KEY2" -s)
echo "Get response for $KEY2: $RESPONSE"

# 3. List operation: Retrieve all key-value pairs
echo "Testing List operation..."
RESPONSE=$(curl -X GET "$BASE_URL/list" -s)
echo "List response: $RESPONSE"

# 4. Delete operation: Delete a key-value pair
echo "Testing Delete operation..."
curl -X DELETE "$BASE_URL/delete?key=$KEY1" -s -o /dev/null -w "HTTP Code: %{http_code}\n"

# 5. Verify Deletion
echo "Testing Get operation after deletion..."
RESPONSE=$(curl -X GET "$BASE_URL/get?key=$KEY1" -s -o /dev/null -w "HTTP Code: %{http_code}\n")
echo "Get response for $KEY1 after deletion: $RESPONSE"

# Final List
echo "Testing List operation after deletion..."
RESPONSE=$(curl -X GET "$BASE_URL/list" -s)
echo "List response: $RESPONSE"

echo "Tests completed."
