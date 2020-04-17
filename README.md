###### This is a blockchain training project that has been programmed using Protocol Buffers and gRPC.



[![Donate](https://img.shields.io/badge/Donate-IRAN-green)](https://idpay.ir/oky2abbas)
<br>**Funding for faster development**  (`only works with IRAN bank cards`)
<br>**Payment gateway**: https://idpay.ir/oky2abbas



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



##### License
    MIT License
    
    Copyright (c) 2018 Abbas Naghdi (@dfmabbas)
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.

