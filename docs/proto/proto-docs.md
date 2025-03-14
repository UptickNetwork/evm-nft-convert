<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [uptick/erc721/v1/erc721.proto](#uptick/erc721/v1/erc721.proto)
    - [TokenPair](#uptick.erc721.v1.TokenPair)
    - [UIDPair](#uptick.erc721.v1.UIDPair)
  
    - [Owner](#uptick.erc721.v1.Owner)
  
- [uptick/erc721/v1/genesis.proto](#uptick/erc721/v1/genesis.proto)
    - [GenesisState](#uptick.erc721.v1.GenesisState)
    - [Params](#uptick.erc721.v1.Params)
  
- [uptick/erc721/v1/query.proto](#uptick/erc721/v1/query.proto)
    - [QueryEvmAddressRequest](#uptick.erc721.v1.QueryEvmAddressRequest)
    - [QueryEvmAddressResponse](#uptick.erc721.v1.QueryEvmAddressResponse)
    - [QueryParamsRequest](#uptick.erc721.v1.QueryParamsRequest)
    - [QueryParamsResponse](#uptick.erc721.v1.QueryParamsResponse)
    - [QueryTokenPairRequest](#uptick.erc721.v1.QueryTokenPairRequest)
    - [QueryTokenPairResponse](#uptick.erc721.v1.QueryTokenPairResponse)
    - [QueryTokenPairsRequest](#uptick.erc721.v1.QueryTokenPairsRequest)
    - [QueryTokenPairsResponse](#uptick.erc721.v1.QueryTokenPairsResponse)
  
    - [Query](#uptick.erc721.v1.Query)
  
- [uptick/erc721/v1/tx.proto](#uptick/erc721/v1/tx.proto)
    - [MsgConvertERC721](#uptick.erc721.v1.MsgConvertERC721)
    - [MsgConvertERC721Response](#uptick.erc721.v1.MsgConvertERC721Response)
    - [MsgConvertNFT](#uptick.erc721.v1.MsgConvertNFT)
    - [MsgConvertNFTResponse](#uptick.erc721.v1.MsgConvertNFTResponse)
    - [MsgTransferERC721](#uptick.erc721.v1.MsgTransferERC721)
    - [MsgTransferERC721Response](#uptick.erc721.v1.MsgTransferERC721Response)
  
    - [Msg](#uptick.erc721.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="uptick/erc721/v1/erc721.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## uptick/erc721/v1/erc721.proto



<a name="uptick.erc721.v1.TokenPair"></a>

### TokenPair
TokenPair defines an instance that records a pairing consisting of a native
Cosmos Coin and an ERC721 token address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc721_address` | [string](#string) |  | address of ERC721 contract token |
| `class_id` | [string](#string) |  | cosmos nft class ID to be mapped to |






<a name="uptick.erc721.v1.UIDPair"></a>

### UIDPair
defines the unique id of nft asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc721_did` | [string](#string) |  | address of ERC721 contract token + tokenId |
| `class_did` | [string](#string) |  | cosmos nft class ID to be mapped to + nftId |





 <!-- end messages -->


<a name="uptick.erc721.v1.Owner"></a>

### Owner
Owner enumerates the ownership of a ERC721 contract.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OWNER_UNSPECIFIED | 0 | OWNER_UNSPECIFIED defines an invalid/undefined owner. |
| OWNER_MODULE | 1 | OWNER_MODULE erc721 is owned by the erc721 module account. |
| OWNER_EXTERNAL | 2 | EXTERNAL erc721 is owned by an external account. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="uptick/erc721/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## uptick/erc721/v1/genesis.proto



<a name="uptick.erc721.v1.GenesisState"></a>

### GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#uptick.erc721.v1.Params) |  | module parameters |
| `token_pairs` | [TokenPair](#uptick.erc721.v1.TokenPair) | repeated | registered token pairs |






<a name="uptick.erc721.v1.Params"></a>

### Params
Params defines the erc721 module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_erc721` | [bool](#bool) |  | parameter to enable the conversion of Cosmos nft <--> ERC721 tokens. |
| `enable_evm_hook` | [bool](#bool) |  | parameter to enable the EVM hook that converts an ERC721 token to a Cosmos NFT by transferring the Tokens through a MsgEthereumTx to the ModuleAddress Ethereum address. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="uptick/erc721/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## uptick/erc721/v1/query.proto



<a name="uptick.erc721.v1.QueryEvmAddressRequest"></a>

### QueryEvmAddressRequest
QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `port` | [string](#string) |  | token identifier can be either the hex contract address of the ERC721 or the Cosmos nft classID |
| `channel` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |






<a name="uptick.erc721.v1.QueryEvmAddressResponse"></a>

### QueryEvmAddressResponse
QueryEvmAddressResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pair` | [TokenPair](#uptick.erc721.v1.TokenPair) |  |  |






<a name="uptick.erc721.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="uptick.erc721.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#uptick.erc721.v1.Params) |  |  |






<a name="uptick.erc721.v1.QueryTokenPairRequest"></a>

### QueryTokenPairRequest
QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token` | [string](#string) |  | token identifier can be either the hex contract address of the ERC721 or the Cosmos nft classID |






<a name="uptick.erc721.v1.QueryTokenPairResponse"></a>

### QueryTokenPairResponse
QueryTokenPairResponse is the response type for the Query/TokenPair RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pair` | [TokenPair](#uptick.erc721.v1.TokenPair) |  |  |






<a name="uptick.erc721.v1.QueryTokenPairsRequest"></a>

### QueryTokenPairsRequest
QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="uptick.erc721.v1.QueryTokenPairsResponse"></a>

### QueryTokenPairsResponse
QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pairs` | [TokenPair](#uptick.erc721.v1.TokenPair) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="uptick.erc721.v1.Query"></a>

### Query
Query defines the gRPC queried service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `TokenPairs` | [QueryTokenPairsRequest](#uptick.erc721.v1.QueryTokenPairsRequest) | [QueryTokenPairsResponse](#uptick.erc721.v1.QueryTokenPairsResponse) | TokenPairs retrieves registered token pairs | GET|/uptick/erc721/v1/token_pairs|
| `TokenPair` | [QueryTokenPairRequest](#uptick.erc721.v1.QueryTokenPairRequest) | [QueryTokenPairResponse](#uptick.erc721.v1.QueryTokenPairResponse) | TokenPair retrieves a registered token pair | GET|/uptick/erc721/v1/token_pairs/{token}|
| `EvmContract` | [QueryEvmAddressRequest](#uptick.erc721.v1.QueryEvmAddressRequest) | [QueryEvmAddressResponse](#uptick.erc721.v1.QueryEvmAddressResponse) | EvmContract retrieves a registered evm contract | GET|/uptick/erc721/v1/evm_contract|
| `Params` | [QueryParamsRequest](#uptick.erc721.v1.QueryParamsRequest) | [QueryParamsResponse](#uptick.erc721.v1.QueryParamsResponse) | Params retrieves the erc721 module params | GET|/uptick/erc721/v1/params|

 <!-- end services -->



<a name="uptick/erc721/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## uptick/erc721/v1/tx.proto



<a name="uptick.erc721.v1.MsgConvertERC721"></a>

### MsgConvertERC721
MsgConvertERC721 defines a Msg to convert a ERC721 token to a native Cosmos
nft.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evm_contract_address` | [string](#string) |  | ERC721 token contract address registered in a token pair |
| `evm_token_ids` | [string](#string) | repeated | tokenID to convert |
| `cosmos_receiver` | [string](#string) |  | bech32 address to receive native Cosmos coins |
| `evm_sender` | [string](#string) |  | sender hex address from the owner of the given ERC721 tokens |
| `class_id` | [string](#string) |  | nft classID to cnvert to ERC721 |
| `cosmos_token_ids` | [string](#string) | repeated | nftID to cnvert to ERC721 |






<a name="uptick.erc721.v1.MsgConvertERC721Response"></a>

### MsgConvertERC721Response
MsgConvertERC721Response returns no fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evm_contract_address` | [string](#string) |  | ERC721 token contract address registered in a token pair |
| `evm_token_ids` | [string](#string) | repeated | tokenID to convert |
| `cosmos_receiver` | [string](#string) |  | bech32 address to receive native Cosmos coins |
| `evm_sender` | [string](#string) |  | sender hex address from the owner of the given ERC721 tokens |
| `class_id` | [string](#string) |  | nft classID to cnvert to ERC721 |
| `cosmos_token_ids` | [string](#string) | repeated | nftID to cnvert to ERC721 |






<a name="uptick.erc721.v1.MsgConvertNFT"></a>

### MsgConvertNFT
MsgConvertNFT defines a Msg to convert a native Cosmos nft to a ERC721 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  | nft classID to cnvert to ERC721 |
| `cosmos_token_ids` | [string](#string) | repeated | nftID to cnvert to ERC721 |
| `evm_receiver` | [string](#string) |  | recipient hex address to receive ERC721 token |
| `cosmos_sender` | [string](#string) |  | cosmos bech32 address from the owner of the given Cosmos coins |
| `evm_contract_address` | [string](#string) |  | ERC721 token contract address registered in a token pair |
| `evm_token_ids` | [string](#string) | repeated | ERC721 token id registered in a token pair |






<a name="uptick.erc721.v1.MsgConvertNFTResponse"></a>

### MsgConvertNFTResponse
MsgConvertNFTResponse returns no fields






<a name="uptick.erc721.v1.MsgTransferERC721"></a>

### MsgTransferERC721
MsgTransferERC721 defines a message for transferring erc721 tokens through IBC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evm_contract_address` | [string](#string) |  | evm_contract_address is the ERC721 token contract address |
| `evm_token_ids` | [string](#string) | repeated | tokenID to convert |
| `source_port` | [string](#string) |  | the port on which the packet will be sent |
| `source_channel` | [string](#string) |  | the channel by which the packet will be sent |
| `class_id` | [string](#string) |  | the class_id of tokens to be transferred |
| `cosmos_token_ids` | [string](#string) | repeated | the non fungible tokens to be transferred |
| `evm_sender` | [string](#string) |  | the sender address |
| `cosmos_receiver` | [string](#string) |  | the recipient address on the destination chain |
| `timeout_height` | [ibc.core.client.v1.Height](#ibc.core.client.v1.Height) |  | timeout_height is the timeout height relative to the current block height The timeout is disabled when set to 0 |
| `timeout_timestamp` | [uint64](#uint64) |  | timeout_timestamp is the timeout timestamp in absolute nanoseconds since unix epoch The timeout is disabled when set to 0 |
| `memo` | [string](#string) |  | memo is an optional memo field |






<a name="uptick.erc721.v1.MsgTransferERC721Response"></a>

### MsgTransferERC721Response
MsgTransferERC721Response defines the response type for Transfer erc721 RPC





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="uptick.erc721.v1.Msg"></a>

### Msg
Msg defines the erc721 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConvertNFT` | [MsgConvertNFT](#uptick.erc721.v1.MsgConvertNFT) | [MsgConvertNFTResponse](#uptick.erc721.v1.MsgConvertNFTResponse) | ConvertNFT mints a ERC721 representation of the native Cosmos nft that is registered on the token mapping. | GET|/uptick/erc721/v1/tx/convert_nft|
| `ConvertERC721` | [MsgConvertERC721](#uptick.erc721.v1.MsgConvertERC721) | [MsgConvertERC721Response](#uptick.erc721.v1.MsgConvertERC721Response) | ConvertERC721 mints a native Cosmos coin representation of the ERC721 token contract that is registered on the token mapping. | GET|/uptick/erc721/v1/tx/convert_erc721|
| `TransferERC721` | [MsgTransferERC721](#uptick.erc721.v1.MsgTransferERC721) | [MsgTransferERC721Response](#uptick.erc721.v1.MsgTransferERC721Response) | TransferERC721 transfers a erc721 token from one chain to another chain through IBC | GET|/uptick/erc721/v1/transfer_erc721|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

