**This is a non-functional work in progress**

Simple web-based kanban board. Jira was too slow and complicated and expensive, Trello was too limited, so I'm making my own.

* Written in Go and AngularJS.
* Only supports MySQL.
* Cards get JIRA-style issue numbers (ie. `ABCD-123`) because I like them.
* No access control yet.


## Installation

1. Set up a MySQL database and load the contents of `schema/mysql.sql` into it.
2. Put your database's Data Source Name (as required by [go-sql-driver](https://github.com/go-sql-driver/mysql)) into the environment variable `DATABASE`. It will look something like `username:password@tcp(hostname:3306)/dbname`.
3. Run the `mtti-board` binary.


## Building from source

Unfortunately, the build process is a bit complicated at the moment.

Requirements:

* [NodeJS and npm](https://nodejs.org/en/)
* [govendor](https://github.com/kardianos/govendor)
* [go-bindata](https://github.com/jteeuwen/go-bindata)

After requirements are installed and in PATH, run `build.sh` and the project will be built into `build/mtti-board` like this:

1. Frontend dependencies are installed with `npm install`, including gulp and browserify.
2. The locally installed gulp is executed to build frontend assets.
3. `go-bindata` embeds the generated frontend assets into the Go build by generating the source file `server/bindata/bindata.go`, which will be ignored by Git.
4. `govendor` builds the final executable.

The temporary asset files in `build/static` can be discarded, their content is embedded in the `mtti-board` binary.
