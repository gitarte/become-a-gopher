FROM postgres
RUN  localedef -i pl_PL -c -f UTF-8 -A /usr/share/locale/locale.alias pl_PL.UTF-8
ENV  LANG              pl_PL.UTF-8
ENV  LANGUAGE          pl_PL.UTF-8
ENV  POSTGRES_PASSWORD mydbpass
ENV  POSTGRES_USER     mydbuser
ENV  POSTGRES_DB       mydb
ENV  PGDATA            /var/lib/postgresql/data/pgdata
COPY 0_schema.sql      /docker-entrypoint-initdb.d/
COPY 1_data.sql        /docker-entrypoint-initdb.d/