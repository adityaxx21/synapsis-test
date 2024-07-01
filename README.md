# Synapsis Test
Backend API Using Go Fiber for Transaction


## How To Setup

 - install golang
 - install postgresql
 - create database `synapsis.Db`
 - setup `.env` based on `.env.example`


## How to run
 - go mod tidy
 - go run /config/main.go

## Current Step of Work
  - [x] Customer can view product list by product category
  - [x] Customer can view product list by product category
  - [x] Customer can add product to shopping cart
  - [x] Customers can see a list of products that have been added to the shopping cart
  - [x] Customer can delete product list in shopping cart
  - [x] Customers can checkout and make payment transactions
  - [x] Login and register customers (it separe role by user and admin with capabilty admin can create item)


## Api Docs
you can check on `api-docs.yaml`

## Issue
- docker didn't work yet cause some issue in connect into database from docker itself
- payment gateway not integrated yet, but i already put the code so i can continue to integrate later