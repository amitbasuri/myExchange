
# x/tokenFreezer

### Abstract

This document specifies the token freezer module of the Cosmos SDK.

The module is responsible for freezing a given amount of tokens on any account. Frozen tokens can’t be moved outside of
account in any way unless unfrozen.

### Freeze

freeze XYZ tokens on any account.

### Unfreeze

unfreeze XYZ tokens on any account.

## State

The x/tokenFreezer module keeps state of amount token is freezed. 

### FreezableAccount interface
```go
// FreezableAccount defines an interface that any vesting account type must implement.
type FreezableAccount interface {
  Account

  FreezeTokens(Coins)
  UnFreezeTokens(Coins)
}
```

## Keepers

The module provides these FreezableAccountKeeper keeper interfaces that can be passed to other modules that read or update
account freezed token.

### FreezableAccount Keeper
```go
type FreezableAccountKeeper interface {
    AccountKeeper

    // Retrieve an FreezableAccount from the store.
    GetFreezableAccount (sdk.Context, sdk.AccAddress) FreezableAccount

    // freeze amount of token
    FreezeTokens (sdk.Context, sdk.AccAddress, Coin) error

    // unfreeze amount of token
    UnFreezeTokens (sdk.Context, sdk.AccAddress, Coin) error
}

```
