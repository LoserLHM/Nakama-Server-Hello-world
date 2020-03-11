package main

import (
  "context"
  "database/sql"
  "encoding/json"
  "github.com/heroiclabs/nakama-common/runtime"
)

// All Go modules must have a InitModule function with this exact signature.
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
  // Register the RPC function.
  if err := initializer.RegisterRpc("HelloWorld_Rpc_id", HelloWorld_Func); err != nil {
    logger.Error("Unable to register: %v", err)
    return err
  }
  return nil
}

func HelloWorld_Func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
  logger.Info("Payload: %s", payload)

  // "payload" is bytes sent by the client we'll JSON decode it.
  var value interface{}
  if err := json.Unmarshal([]byte(payload), &value); err != nil {
    return "", err
  }

  response, err := json.Marshal("Hello-World")
  if err != nil {
    return "", err
  }

    return string(response), nil
}
