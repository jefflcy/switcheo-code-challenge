package todolist

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "todolist/api/todolist/todolist"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowItem",
					Use:            "show-item [id]",
					Short:          "Query show-item",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListItem",
					Use:            "list-item",
					Short:          "Query list-item",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateItem",
					Use:            "create-item [desc] [priority]",
					Short:          "Send a create-item tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "desc"}, {ProtoField: "priority"}},
				},
				{
					RpcMethod:      "DeleteItem",
					Use:            "delete-item [id]",
					Short:          "Send a delete-item tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "UpdateItem",
					Use:            "update-item [desc] [priority] [id]",
					Short:          "Send a update-item tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "desc"}, {ProtoField: "priority"}, {ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
