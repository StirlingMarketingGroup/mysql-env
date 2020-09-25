# MySQL env

A small MySQL UDF library for setting and getting environment variables written in Golang.

---

### `setenv`

Sets an environment variable.

```sql
`setenv` ( `key` , `value` )
```
- `` `key` ``
  - The name of the environment variable to set
- `` `value` ``
  - The value to store
---
  ### `getenv`

Gets an environment variable. Returns `NULL` if the name given is `NULL`. Case sensitive.

```sql
`getenv` ( `key` )
```
- `` `key` ``
  - The name of the environment variable to get
---
  ### `unsetenv`

Unsets an environment variable. Case sensitive.

```sql
`unsetenv` ( `key` )
```
- `` `key` ``
  - The name of the environment variable to unset
---
## Examples

```sql
select`setenv`('FIZZ','buzz');
-- NULL

select`getenv`('FIZZ');
-- 'buzz'

select`unsetenv`('FIZZ');
-- NULL

select`getenv`('FIZZ');
-- ''

select`getenv`('PATH');
-- '/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin'

select`getenv`('path'); -- NOTICE the lowercase here
-- ''
```
---

## Dependencies

You will need Golang, which you can get from here https://golang.org/doc/install. You will also need the MySQL dev library.

Debian / Ubuntu
```shell
sudo apt update
sudo apt install libmysqlclient-dev
```
## Installing

You can find your MySQL plugin directory by running this MySQL query

```sql
select @@plugin_dir;
```

then replace `/usr/lib/mysql/plugin` below with your MySQL plugin directory.

```shell
cd ~ # or wherever you store your git projects
git clone https://github.com/StirlingMarketingGroup/mysql-env.git
cd mysql-env
go get -d ./...
go build -buildmode=c-shared -o env.so
sudo cp env.so /usr/lib/mysql/plugin/ # replace plugin dir here if needed
```

Enable the functions in MySQL by running this MySQL query

```sql
create function`Setenv`returns int soname'env.so';
create function`Getenv`returns string soname'env.so';
create function`Unsetenv`returns int soname'env.so';
```