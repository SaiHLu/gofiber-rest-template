**Setup Guide**

- Install golang from [Golang](https://go.dev/doc/install)
- Setup your local/remote database and redis that match with **.env.example**
- RUN
  - go mod tidy
  - make migrate # create table
  - air (if not installed yet, reference [Air](https://github.com/air-verse/air))
