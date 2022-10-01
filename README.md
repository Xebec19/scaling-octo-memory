# scaling-octo-memory

## Git hooks set up
Please write 
```
chmod +x .githooks/*
```
to make git hooks executable

## Prerequisites
1. Docker
2. Go migrate
3. Go 1.19

## Commands 
1. We need an docker image of postgres tag: 12-alpine 
    to set up postgres
    ``` 
    make postgres
    ```
2. To create database
    ```
    make createdb
    ```
3. To delete database
    ```
    make dropdb
    ```
4. To set up schema of database
    ```
    make migrateup
    ```
5. To revert migration
    ```
    make migratedown
    ```
6. To start development server
    ```
    make live
    ```