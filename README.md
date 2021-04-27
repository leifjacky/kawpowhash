# Kawpowhash

For details on this project, please see the Ethereum wiki:
https://github.com/ethereum/wiki/wiki/Ethash

### Coding Style for C++ code:

Follow the same exact style as in [cpp-ethereum](https://github.com/ethereum/cpp-ethereum/blob/develop/CodingStandards.txt)

### Coding Style for C code:

The main thing above all is code consistency.

- Tabs for indentation. A tab is 4 spaces
- Try to stick to the [K&R](http://en.wikipedia.org/wiki/Indent_style#K.26R_style),
  especially for the C code.
- Keep the line lengths reasonable. No hard limit on 80 characters but don't go further
  than 110. Some people work with multiple buffers next to each other.
  Make them like you :)



ProgPoW can be tuned using the following parameters.  The proposed settings have been tuned for a range of existing, commodity GPUs:

* `PROGPOW_PERIOD`: Number of blocks before changing the random program
* `PROGPOW_LANES`: The number of parallel lanes that coordinate to calculate a single hash instance
* `PROGPOW_REGS`: The register file usage size
* `PROGPOW_DAG_LOADS`: Number of uint32 loads from the DAG per lane
* `PROGPOW_CACHE_BYTES`: The size of the cache
* `PROGPOW_CNT_DAG`: The number of DAG accesses, defined as the outer loop of the algorithm (64 is the same as Ethash)
* `PROGPOW_CNT_CACHE`: The number of cache accesses per loop
* `PROGPOW_CNT_MATH`: The number of math operations per loop

The value of these parameters has been tweaked between version 0.9.2 (live on the Gangnam testnet) and 0.9.3 (proposed for Ethereum adoption).  See [this medium post](https://medium.com/@ifdefelse/progpow-progress-da5bb31a651b) for details.

| Parameter             | 0.9.2     | 0.9.3     | KAWPOW    |
| --------------------- | --------- | --------- | --------- |
| `PROGPOW_PERIOD`      | `50`      | `10`      | `3`       |
| `PROGPOW_LANES`       | `16`      | `16`      | `16`      |
| `PROGPOW_REGS`        | `32`      | `32`      | `32`      |
| `PROGPOW_DAG_LOADS`   | `4`       | `4`       | `4`       |
| `PROGPOW_CACHE_BYTES` | `16x1024` | `16x1024` | `16x1024` |
| `PROGPOW_CNT_DAG`     | `64`      | `64`      | `64`      |
| `PROGPOW_CNT_CACHE`   | `12`      | `11`      | `11`      |
| `PROGPOW_CNT_MATH`    | `20`      | `18`      | `18`      |

KAWPOW_EPOCH_LENGTH = 7500

## Usage

```bash
go get -v -u github.com/leifjacky/kawpowhash
cd ${GOPATH}/src/github.com/leifjacky/kawpowhash
go test -v ./...
```



