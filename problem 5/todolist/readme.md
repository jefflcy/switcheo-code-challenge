# todolist

**todolist** is a basic practical CRUD blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

**Make sure your go/bin folder is on our PATH or the shell may not recognize the command.**

For bash shell,

```
nano .bash_profile
#add this line of code to your profile
export PATH="$HOME/go/bin:$PATH"
```

For mac zsh,

```
nano .bash_profile
#add this line of code to your profile
export PATH="$HOME/go/bin:$PATH"
```

## CRUD functionalities using CLI

### Create a todolist item by Alice

- Structure:
  `[blockchain daemon] [tx type] [module name] [msg] [item desc] [item priority] [from flag] [chain id flag]`

```
todolistd tx todolist create-item "Code up solution" 3 --from alice --chain-id todolist
```

### Show the item (using item id)

- Structure:
  `[blockchain daemon] [tx type] [module name] [msg] [item id]`

```
todolistd q todolist show-item 0 // show the first item in the store
```

### Update the item (rewrite all params, using item id)

- Structure:
  `[blockchain daemon] [tx type] [module name] [msg] [new item desc] [new item priority] [item id] [from flag] [chain id flag]`

```
todolistd tx todolist update-item "Wash the red car" 2 0 --from alice --chain-id todolist
```

### Delete the item (using item id)

- Structure:
  `[blockchain daemon] [tx type] [module name] [msg] [item id] [from flag] [chain id flag]`

```
todolistd tx todolist delete-item 0 --from alice --chain-id todolist
```

### List the stored todolist items

- Structure:
  `[blockchain daemon] [tx type] [module name] [msg]`

To show all items:

```
todolistd q todolist list-all-items
```

To show just the **highest priority** item (largest priority int):

```
todolistd q todolist list-item
```
