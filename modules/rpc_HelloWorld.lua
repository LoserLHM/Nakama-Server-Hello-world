local nk = require("nakama")

local function get_HelloWorld(context, payload)

  local json = nk.json_decode(payload)

  nk.logger_info(("Payload: HelloWorld"):format(json))

  return nk.json_encode({"HelloWorld"})
end

nk.register_rpc(get_HelloWorld, "SayHelloWorld_rpc_func_ID")
