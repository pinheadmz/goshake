# goshake

## Handshake (HNS) tools written in go

### Install go

Go 1.11 is required at minimum for Go Modules. This package was built with Go 1.14.
Installation assumes your `$GOPATH` is in your environment's `$PATH`.
If you just installed go, you may just need to add these line to `~/.profile`:

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

then `source ~/.profile`


### Install goshake

To install all tools on your system:


```
git clone https://github.com/pinheadmz/goshake
cd goshake
GO111MODULE=on go install -v ./cmd/...
```

### Usage

There is only one command at this time: `hsdtx`

This command will decode a raw Handshake transaction into human readble elements:

```
$ hsdtx 00000000018e9524fcf5b2f54539bc48b1db0d3e23bf1e4c1e1f54a9ca2662eb0641819f6f01000000ffffffff02408af701000000000014162283266a81d5c960cf44f0d4f0f54c69d9e32803042087d48e21f6aeae6f068535fe90c423a21471cb66f60af34a7280bce9e9aff4e3041065000013636b687574636869736f6e686f6c64696e6773204dc5859aad8b5b24cbb5b18c91979b3d3445a838030521b131e76af4bc15b7f34cb513520200000000149cf180f590bf328fe6003c8c3f77b211de16515b0000000000000241fa654bac97e8aba0a1fe9ee8a8ff6260fa396b88a49d092213194bff3f0cb31c5e77d68b260a62d11d7d15bfba5065cf5ab7b6ca712eda1c112358bc76ff8f86012102d942e3f658abaeb88f2cf1d410f13cc4ea98fca26459ff32103387cd8453c630

version: 00000000 
input count: 01 
 input #0: 
  txid: 8e9524fcf5b2f54539bc48b1db0d3e23bf1e4c1e1f54a9ca2662eb0641819f6f 
  index: 00000001 
  sequence: ffffffff 
output count: 02 
 output #0: 
  value: 33000000 
  address version: 00 
  address hash size: 14 (20 bytes) 
  address hash: 162283266a81d5c960cf44f0d4f0f54c69d9e328 
  covenant type: 03 
  covenant item count: 04 
   item #0: 
    item size: 20 (32 bytes) 
    item: 87d48e21f6aeae6f068535fe90c423a21471cb66f60af34a7280bce9e9aff4e3 
   item #1: 
    item size: 04 (4 bytes) 
    item: 10650000 
   item #2: 
    item size: 13 (19 bytes) 
    item: 636b687574636869736f6e686f6c64696e6773 
   item #3: 
    item size: 20 (32 bytes) 
    item: 4dc5859aad8b5b24cbb5b18c91979b3d3445a838030521b131e76af4bc15b7f3 
 output #1: 
  value: 1377023308 
  address version: 00 
  address hash size: 14 (20 bytes) 
  address hash: 9cf180f590bf328fe6003c8c3f77b211de16515b 
  covenant type: 00 
  covenant item count: 00 
locktime: 00000000 
witness: 
 witness for input #0: 
   stack size: 02 
    item #0: 
     item size: 41 (65 bytes) 
     item: fa654bac97e8aba0a1fe9ee8a8ff6260fa396b88a49d092213194bff3f0cb31c5e77d68b260a62d11d7d15bfba5065cf5ab7b6ca712eda1c112358bc76ff8f8601 
    item #1: 
     item size: 21 (33 bytes) 
     item: 02d942e3f658abaeb88f2cf1d410f13cc4ea98fca26459ff32103387cd8453c630 

```

