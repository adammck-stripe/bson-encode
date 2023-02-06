# bson-encode

For when you need to generate some b64-encoded BSON.  
See also: [bson-decode](https://github.com/adammck-stripe/bson-decode).

## Installation

```console
$ mkdir -p $HOME/adammck-stripe
$ cd $HOME/adammck-stripe
$ git clone https://github.com/adammck-stripe/bson-encode.git
$ cd bson-encode
$ go install
$ which bson-encode
```

## Usage

```console
$ bson-encode '[{"name": "pi"}, {"value": 3.14159}]'
MQAAAAMwABIAAAACbmFtZQADAAAAcGkAAAMxABQAAAABdmFsdWUAboYb8PkhCUAAAA==
```

## License

MIT
