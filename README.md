# Clean Architecture Go gRPC
This is example of implementation clean architecure in golang and gRPC.

## Installation
```bash
make all
```

## Set the config
Open config.json Change to your own database, port, log, app config
```json
{
  "app": {
    "name": "clean-arch-go"
  },
  "web": {
    "port": 3000
  },
  "log": {
    "level": 6
  },
  "database": {
    "username": "product",
    "password": "mysql123",
    "host": "localhost",
    "port": "5432",
    "name": "product-db",
    "pool": {
      "idle": 10,
      "max": 100,
      "lifetime": 300
    }
  }
}
```

## Usage
```bash
make run
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)