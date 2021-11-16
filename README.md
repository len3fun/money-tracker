### Deployment
Change or create `config.yml` in `./configs/`.

`config.yml` example:
~~~
port: "8000"

db:
  host: "db"
  port: "5432"
  dbname: "some_db"
  sslmode: "disable"
~~~

Create `.env` and `db.env` files. 

`.env` example:
~~~
DB_PASSWORD=some_password
DB_USERNAME=some_username
DEBUG=false
~~~
`db.env` example:
~~~
POSTGRES_USER=some_user
POSTGRES_PASSWORD=some_password
POSTGRES_DB=some_db
~~~
Build the docker image:
~~~
docker build -t money-tracker .
~~~
Run docker compose:
~~~
docker-compose up -d
~~~
