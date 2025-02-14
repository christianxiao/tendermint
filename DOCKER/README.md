# Docker

## Supported tags and respective `Dockerfile` links

- `0.17.1`, `latest` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/208ac32fa266657bd6c304e84ec828aa252bb0b8/DOCKER/Dockerfile)
- `0.15.0` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/170777300ea92dc21a8aec1abc16cb51812513a4/DOCKER/Dockerfile)
- `0.13.0` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/a28b3fff49dce2fb31f90abb2fc693834e0029c2/DOCKER/Dockerfile)
- `0.12.1` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/457c688346b565e90735431619ca3ca597ef9007/DOCKER/Dockerfile)
- `0.12.0` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/70d8afa6e952e24c573ece345560a5971bf2cc0e/DOCKER/Dockerfile)
- `0.11.0` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/9177cc1f64ca88a4a0243c5d1773d10fba67e201/DOCKER/Dockerfile)
- `0.10.0` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/e5342f4054ab784b2cd6150e14f01053d7c8deb2/DOCKER/Dockerfile)
- `0.9.1`, `0.9`, [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/809e0e8c5933604ba8b2d096803ada7c5ec4dfd3/DOCKER/Dockerfile)
- `0.9.0` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/d474baeeea6c22b289e7402449572f7c89ee21da/DOCKER/Dockerfile)
- `0.8.0`, `0.8` [(Dockerfile)](https://github.com/christianxiao/tendermint/blob/bf64dd21fdb193e54d8addaaaa2ecf7ac371de8c/DOCKER/Dockerfile)

## Quick reference

- **Where to get help:**
  [cosmos.network/ecosystem](https://cosmos.network/ecosystem)

- **Where to file issues:**
  [Tendermint Issues](https://github.com/christianxiao/tendermint/issues)

- **Supported Docker versions:**
  [the latest release](https://github.com/moby/moby/releases) (down to 1.6 on a best-effort basis)

## Tendermint

Tendermint Core is Byzantine Fault Tolerant (BFT) middleware that takes a state transition machine, written in any programming language, and securely replicates it on many machines.

For more background, see the [the docs](https://tendermint.com/docs/introduction/#quick-start).

To get started developing applications, see the [application developers guide](https://tendermint.com/docs/introduction/quick-start.html).

## How to use this image

### Start one instance of the Tendermint core with the `kvstore` app

A quick example of a built-in app and Tendermint core in one container.

```
docker run -it --rm -v "/tmp:/tendermint" tendermint/tendermint init
docker run -it --rm -v "/tmp:/tendermint" tendermint/tendermint node --proxy_app=kvstore
```

## Local cluster

To run a 4-node network, see the `Makefile` in the root of [the repo](https://github.com/christianxiao/tendermint/blob/master/Makefile) and run:

```
make build-linux
make build-docker-localnode
make localnet-start
```

Note that this will build and use a different image than the ones provided here.

## License

- Tendermint's license is [Apache 2.0](https://github.com/christianxiao/tendermint/blob/master/LICENSE).

## Contributing

Contributions are most welcome! See the [contributing file](https://github.com/christianxiao/tendermint/blob/master/CONTRIBUTING.md) for more information.
