# go-utils/database_wrapper

A database wrapper (for SQLite in this case, though it could be used for any provider) that can run
in-process migrations and seeders without the need for an external tool.

This is heavily influenced and based on Ben B Johnson's article and WTF Dial example codebase:
https://github.com/benbjohnson/wtf/blob/main/sqlite/sqlite.go
