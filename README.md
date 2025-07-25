---
title: Non-Fungible Token Convert between evm and cosmos native nft
category: IBC/APP
requires: 25, 26,721
kind: instantiation
author: uptsmart <uptsmart@163.com>
created: 2022-03-04
modified: 2022-03-06
---

## Abstract

EVM and Cosmos are two different blockchain technologies, and conversion between them involves different protocols and bridging mechanisms. In this case, the conversion from EVM to Cosmos native NFT can realize the transfer of NFT issued on the Ethereum blockchain to the Cosmos blockchain for circulation and transactions. This conversion function can help NFT holders transfer their assets from one blockchain to another for better market liquidity and wider usage scenarios.

When performing this conversion, a corresponding bridge protocol or cross-chain interaction protocol needs to be used to establish the connection between the EVM and the Cosmos blockchain. Once this connection is established, NFT holders can transfer their assets between the two blockchains through specific operations. This conversion function can help expand the coverage of the NFT market, promote cross-chain transactions and interoperability, and improve the liquidity and use value of NFT.


### Desired Properties

- Preservation of non-fungibility (i.e., only one instance of any token is *live* across all the IBC-connected blockchains).
- Permissionless token transfers, no need to whitelist connections, modules, or `classId`s.
- Symmetric (all chains implement the same logic, no in-protocol differentiation of hubs & zones).

## Technical Specification

### Data Structures

#### Convert from cosmos NFT to evm ERC721 asset

```typescript
// MsgConvertNFT defines a Msg to convert a native Cosmos nft to a ERC721 token
message MsgConvertNFT {
  // nft classID to cnvert to ERC721
  string class_id = 1;
  // nftID to cnvert to ERC721
  repeated string cosmos_token_ids = 2;
  // recipient hex address to receive ERC721 token
  string evm_receiver = 3;
  // cosmos bech32 address from the owner of the given Cosmos coins
  string cosmos_sender = 4;
  // ERC721 token contract address registered in a token pair
  string evm_contract_address = 5;
  // ERC721 token id registered in a token pair
  repeated string evm_token_ids = 6;
}
```
`classId` is a required field that MUST never be empty, it uniquely identifies the class/collection/contract which the tokens being transferred belong to in the sending chain. In the case of an ERC-1155 compliant smart contract, for example, this could be a string representation of the top 128 bits of the token ID.

`cosmos_token_ids` array is a required field that MUST have a size greater than zero and hold non-empty entries that uniquely identify tokens (of the given class) that are being transferred. In the case of an ERC-1155 compliant smart contract, for example, a `cosmos_token_ids` could be a string representation of the bottom 128 bits of the token ID.

`evm_receiver` recipient hex address to receive ERC721 token.

`evm_contract_address` ERC721 token contract address registered in a token pair.

`evm_token_ids` ERC721 token contract address registered in a token pair.


#### Convert from  evm ERC721 to cosmos NFT asset

```typescript
// MsgConvertERC721 defines a Msg to convert a ERC721 token to a native Cosmos
// nft.
message MsgConvertERC721 {
  // ERC721 token contract address registered in a token pair
  string evm_contract_address = 1;
  // tokenID to convert
  repeated string evm_token_ids = 2;
  // bech32 address to receive native Cosmos coins
  string cosmos_receiver = 3;
  // sender hex address from the owner of the given ERC721 tokens
  string evm_sender = 4;
  // nft classID to cnvert to ERC721
  string class_id = 5;
  // nftID to cnvert to ERC721
  repeated string cosmos_token_ids = 6;
}
```
`evm_contract_address` ERC721 token contract address registered in a token pair.

`evm_token_ids` ERC721 token contract address registered in a token pair.

`cosmos_receiver` bech32 address to receive native Cosmos coins.

`evm_sender` sender hex address from the owner of the given ERC721 tokens.

`classId` is a required field that MUST never be empty, it uniquely identifies the class/collection/contract which the tokens being transferred belong to in the sending chain. In the case of an ERC-1155 compliant smart contract, for example, this could be a string representation of the top 128 bits of the token ID.

`cosmos_token_ids` array is a required field that MUST have a size greater than zero and hold non-empty entries that uniquely identify tokens (of the given class) that are being transferred. In the case of an ERC-1155 compliant smart contract, for example, a `cosmos_token_ids` could be a string representation of the bottom 128 bits of the token ID.


#### related SDK
- Implementation of SDK in nodejs can be found in [uptick-chain-sdk.js](https://github.com/UptickNetwork/uptick-chain-sdk.js).


#### technical documentation
- Related technical documentation can be found in [NFT Conversion Document](https://app.gitbook.com/o/uPfC9w7sfZt6S3IXypqI/s/6zlFzpQT9NAx43Tcz0mE/).

## Backwards Compatibility
Not applicable.


## Copyright

All content herein is licensed under [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0).



