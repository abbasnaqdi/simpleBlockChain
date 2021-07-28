This is a **blockchain** training project that has been programmed using **`Protocol Buffers`** and **`gRPC`**.



[![Donate](https://img.shields.io/badge/Cryptocurrency-Donate-green)](https://idpay.ir/oky2abbas) **BTC**: `1HPZyUP9EJZi2S87QrvCDrE47qRV4i5Fze`

[![Donate](https://img.shields.io/badge/Cryptocurrency-Donate-blue)](https://idpay.ir/oky2abbas) **ETH or USDT**: `0x4a4b0A26Eb31e9152653E4C08bCF10f04a0A02a9`



##### Dependencies
```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
go get -u golang.org/x/net/context
go get -u github.com/golang/protobuf/proto
```

##### Get Started


###### Step One
```
go run server/server.go             //run server block chain
```


###### Step Two
```
go run client/client.go --start     //start connecting the chains
go run client/client.go --add       //added a block to the chain
go run client/client.go --list      //get all of the generated block chains
```

