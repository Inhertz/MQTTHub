FROM mcr.microsoft.com/mssql/server

RUN mkdir /src
WORKDIR /src

COPY . /src
COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /src/create-db.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 1433

ENTRYPOINT ["entrypoint.sh"]

