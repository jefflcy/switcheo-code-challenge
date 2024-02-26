# todolist

**todolist** is a basic practical CRUD blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Consensus-Breaking Change

**In** `x/todolist/types/message_create_item.go` & `x/todolist/types/message_update_item.go`:

```
func NewMsgCreateItem(creator string, desc string, priority uint64) *MsgCreateItem {
	return &MsgCreateItem{
		Creator:  creator,
		Desc:     desc, // removed the "priority" field
	}
}

func NewMsgUpdateItem(creator string, desc string, priority uint64, id uint64) *MsgUpdateItem {
	return &MsgUpdateItem{
		Creator:  creator,
		Desc:     desc,
		Id:       id, // removed the "priority" field
	}
}
```

**In** `x/todolist/keeper/msg_server_create_item.go` & `x/todolist/keeper/msg_server_update_item.go`:

```
// In CreateItem
	var item = types.Item{
		Creator: msg.Creator,
		Desc:    msg.Desc, // removed the "priority" field
	}

// In UpdateItem
var item = types.Item{
		Creator: msg.Creator,
		Id:      msg.Id,
		Desc:    msg.Desc, // removed the "priority" field
	}

```

**In** `proto/todolist/tx.proto`:

```
message MsgCreateItem {
  option (cosmos.msg.v1.signer) = "creator";
  string creator  = 1;
  string desc     = 2; //removed the "priority" field
}

message MsgUpdateItem {
  option (cosmos.msg.v1.signer) = "creator";
  string creator  = 1;
  string desc     = 2;
  uint64 id       = 3; //removed the "priority" field
}
```

- What does it mean to break consensus?

Breaking consensus in a blockchain context means making a change that is not backwards compatible as it would require all nodes to upgrade to the new version to continue participating in the network. Consensus is the process in which all nodes in the network agree on the validity of transactions and the state of the blockchain.

That is why if there is a change introduced that affects this agreement process, like altering the transaction structure, and if some nodes adopt this change while others do not, those nodes will no longer be able to agree on a unified state of the ledger.

This disagreement can then lead to forks in a blockchain, where different nodes have different views of the transaction history and state.

- Why does my change break consensus?

The removal of the "priority" field from the create-item/update-item transaction would break consensus because the past transactions that were considered valid would now be invalid according to this new code.

Nodes running the old version expect the "priority" field to be present and may reject transactions without it, while nodes that have updated would consider those same transactions valid.

As a result, the network would split into two incompatible versions, with each set of nodes having a different view of the blockchain's state. This is a consensus-breaking change because it requires all nodes to update synchronously to maintain a unified network. If this coordinated update does not happen, it could lead to a network fork.
