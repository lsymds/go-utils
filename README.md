# go-utils

Numerous utilities for building go projects. Each folder is a runnable sub-project with tests to allow this repository to grow without ever exploding when you try and use it.

# Utilities

* http_api - A selection of utilities for building HTTP based APIs in Go. Not designed to be
             imported.
* database_wrapper - A database wrapper (for SQLite in this case, though it could be used for any provider) that can run
                     in-process migrations and seeders without the need for an external tool. Not
                     designed to be imported.
* slices - A selection of utility slice functions. Designed to be imported.
