# Install Tendermint

The fastest and easiest way to install the `tendermint` binary
is to run [this script](https://github.com/christianxiao/tendermint/blob/master/scripts/install/install_tendermint_ubuntu.sh) on
a fresh Ubuntu instance,
or [this script](https://github.com/christianxiao/tendermint/blob/master/scripts/install/install_tendermint_bsd.sh)
on a fresh FreeBSD instance. Read the comments / instructions carefully (i.e., reset your terminal after running the script,
make sure you are okay with the network connections being made).

## From Binary

To download pre-built binaries, see the [releases page](https://github.com/christianxiao/tendermint/releases).

## From Source

You'll need `go` [installed](https://golang.org/doc/install) and the required
[environment variables set](https://github.com/christianxiao/tendermint/wiki/Setting-GOPATH)

### Get Source Code

```
mkdir -p $GOPATH/src/github.com/tendermint
cd $GOPATH/src/github.com/tendermint
git clone https://github.com/christianxiao/tendermint.git
cd tendermint
```

### Get Tools & Dependencies

```
make tools
```

### Compile

```
make install
```

to put the binary in `$GOPATH/bin` or use:

```
make build
```

to put the binary in `./build`.

_DISCLAIMER_ The binary of tendermint is build/installed without the DWARF symbol table. If you would like to build/install tendermint with the DWARF symbol and debug information, remove `-s -w` from `BUILD_FLAGS` in the make file.

The latest `tendermint version` is now installed.

## Run

To start a one-node blockchain with a simple in-process application:

```
tendermint init
tendermint node --proxy_app=kvstore
```

## Reinstall

If you already have Tendermint installed, and you make updates, simply

```
cd $GOPATH/src/github.com/christianxiao/tendermint
make install
```

To upgrade, run

```
cd $GOPATH/src/github.com/christianxiao/tendermint
git pull origin master
make install
```

## Compile with CLevelDB support

Install [LevelDB](https://github.com/google/leveldb) (minimum version is 1.7).

Install LevelDB with snappy (optionally). Below are commands for Ubuntu:

```
sudo apt-get update
sudo apt install build-essential

sudo apt-get install libsnappy-dev

wget https://github.com/google/leveldb/archive/v1.20.tar.gz && \
  tar -zxvf v1.20.tar.gz && \
  cd leveldb-1.20/ && \
  make && \
  sudo cp -r out-static/lib* out-shared/lib* /usr/local/lib/ && \
  cd include/ && \
  sudo cp -r leveldb /usr/local/include/ && \
  sudo ldconfig && \
  rm -f v1.20.tar.gz
```

Set a database backend to `cleveldb`:

```
# config/config.toml
db_backend = "cleveldb"
```

To install Tendermint, run:

```
CGO_LDFLAGS="-lsnappy" make install_c
```

or run:

```
CGO_LDFLAGS="-lsnappy" make build_c
```

which puts the binary in `./build`.
