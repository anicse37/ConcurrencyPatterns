#!/bin/bash

echo "Hey Lazzzzzzzzy"
sleep 1

echo "Running Fan In/Fan Out"
go run fan_in_out/main.go
sleep 3

clear
echo "Running Pipeline"
echo "V1"
go run pipeline/V1/main.go
sleep 2

echo "V2"
go run pipeline/V2/main.go
sleep 2

clear 
echo "Running Worker Pool"
go run worker\ pool/main.go
sleep 3

echo "Running Atomic"
go run atomic/main.go
sleep 1

echo " "
echo "Nikl"
