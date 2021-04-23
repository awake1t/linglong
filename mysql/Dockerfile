FROM mysql:5.7

COPY ./init.sql /docker-entrypoint-initdb.d

COPY my.cnf /etc/mysql/conf.d/my.cnf

RUN chmod 0444 /etc/mysql/conf.d/my.cnf
