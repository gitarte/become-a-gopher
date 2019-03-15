FROM postgres
RUN  localedef -i pl_PL -c -f UTF-8 -A /usr/share/locale/locale.alias pl_PL.UTF-8
COPY 0_schema.sql      /docker-entrypoint-initdb.d/
COPY 1_data.sql        /docker-entrypoint-initdb.d/