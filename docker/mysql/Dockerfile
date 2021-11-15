FROM mysql:5.7

COPY ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./docker/mysql/init.sql /docker-entrypoint-initdb.d
RUN chmod 644 /etc/mysql/conf.d/my.cnf
RUN chmod 644 /docker-entrypoint-initdb.d
