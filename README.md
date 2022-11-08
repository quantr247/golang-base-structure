# golang-base-structure
1. Firstly, plz check Makefile to generate folder structure for project
    - cmd               # Main application for this project
        - server        # Contain main.go file for server side
    - config            # Configuration file
    - internal          # Private package of application
        - adapter       # Service for external/third-party api
            - model     # Data struct for external/third-party api
        - api           # Define http api service of server side
        - common        # Constant 
        - dto           # Data struct for http api service of server side
        - helper        # Dependency library need for this project
        - registry      # Dependency injection container
        - repository    # Database implement for function interface
        - usecase       # Handle business logic
        - util          # Util for this project    

2. Run make install to install all dependency
    make depends
    make install

3. Run make run to test
    make run    

# Explain of this project
1. Base structure for Golang project with:
- HTTP gateway
- DB connection to: Oracle, Postgres, SQLServer
- Memcache
- Redis cache
- Dependency injection
- Logging with zap

2. OracleDB
- To implement with Oracle db, need add Golang Oracle database driver: https://github.com/mattn/go-oci8

3. Check file /registry/di_container.go to remove comment for connect to DB