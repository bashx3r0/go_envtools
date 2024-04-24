## ENVTOOLS
Very simple go package to load ENV and get your specified variables

## Features
---

### Get all variables 

`envtools.GetAllEnv(envpath)`

Accepts single string parameter `envpath` - your ENV file location. Provide empty string `""` to default to `.env`
Returns a map of all env [ENV_NAME]ENV_VALUE


### Get single variable 

`envtools.GetAllEnv(envpath, envvar)`

Accepts 2 string parameters:
1.  `envpath` - your ENV file location. Provide empty string `""` to default to `.env`
2.  `envvar` - your environment variable name

Returns a value of specified `envvar`
