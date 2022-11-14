# PostgreSQL

If configured, the PostgreSQL integration will write device data into an
[PostgreSQL](https://www.postgresql.org/) database.

## Creating database

You must create a separate database for using the PostgreSQL integration.
To enter the command line utility for PostgreSQL:

```bash
sudo -u postgres psql
```

Inside this prompt, execute the following queries to setup the
`chirpstack_integration` database. It is recommended to use a different
username (role) and password.

```sql
-- create role for authentication
create role chirpstack_integration with login password 'chirpstack_integration';

-- create database
create database chirpstack_integration with owner chirpstack_integration;

-- exit psql
\q
```

## Activating integration

This integration must be configured in the [Configuration](../configuration.md)
file.

### Enable integration

In the configuration file:

```toml
[integration]
  enabled = [
    "mqtt",
  ]
```

Your `enabled` line may look slightly different. Add `"postgresql"` to the list.
In this case, the modified line should appear as enabled=["mqtt", "postgresql"].

### Configuration

You must also add / update the `[integration.postgresql]` section with the
hostname, username, password and the name of the database. Example:

```toml
[integration.postgresql]
  dsn="postgres://<username>:<password>@<host>/<database>?sslmode=disable"
```

In the dns= line, modify `<username>`, `<password>`, `<host>`, and `<database>`
with your appropriate credentials and targets. If you followed the example above,
you would use chirpstack_integration as your username and target database. If
your target Postgres database is on the same machine as ChirpStack, use
localhost as your host.
