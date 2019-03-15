#!/bin/bash
export POSTGRES_LANG=pl_PL.UTF-8
export POSTGRES_LANGUAGE=pl_PL.UTF-8
export POSTGRES_PASSWORD=mydbpass
export POSTGRES_USER=mydbuser
export POSTGRES_DB=mydb
export POSTGRES_PGDATA=/var/lib/postgresql/data/pgdata
export POSTGRES_HOST=db
export POSTGRES_PORT=5432
docker-compose -f docker-compose.yml build

