// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

// generated by cmd/generatewshclient/main-generatewshclient.go

package wshclient

import (
	"github.com/wavetermdev/thenextwave/pkg/wshutil"
	"github.com/wavetermdev/thenextwave/pkg/wshrpc"
	"github.com/wavetermdev/thenextwave/pkg/waveobj"
)

// command "controller:input", wshserver.BlockInputCommand
func BlockInputCommand(w *wshutil.WshRpc, data wshrpc.CommandBlockInputData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "controller:input", data, opts)
    return err
}

// command "controller:restart", wshserver.BlockRestartCommand
func BlockRestartCommand(w *wshutil.WshRpc, data wshrpc.CommandBlockRestartData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "controller:restart", data, opts)
    return err
}

// command "createblock", wshserver.CreateBlockCommand
func CreateBlockCommand(w *wshutil.WshRpc, data wshrpc.CommandCreateBlockData, opts *wshrpc.WshRpcCommandOpts) (*waveobj.ORef, error) {
    resp, err := sendRpcRequestCallHelper[*waveobj.ORef](w, "createblock", data, opts)
    return resp, err
}

// command "deleteblock", wshserver.DeleteBlockCommand
func DeleteBlockCommand(w *wshutil.WshRpc, data wshrpc.CommandDeleteBlockData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "deleteblock", data, opts)
    return err
}

// command "file:append", wshserver.AppendFileCommand
func AppendFileCommand(w *wshutil.WshRpc, data wshrpc.CommandFileData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "file:append", data, opts)
    return err
}

// command "file:appendijson", wshserver.AppendIJsonCommand
func AppendIJsonCommand(w *wshutil.WshRpc, data wshrpc.CommandAppendIJsonData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "file:appendijson", data, opts)
    return err
}

// command "file:read", wshserver.ReadFile
func ReadFile(w *wshutil.WshRpc, data wshrpc.CommandFileData, opts *wshrpc.WshRpcCommandOpts) (string, error) {
    resp, err := sendRpcRequestCallHelper[string](w, "file:read", data, opts)
    return resp, err
}

// command "file:write", wshserver.WriteFile
func WriteFile(w *wshutil.WshRpc, data wshrpc.CommandFileData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "file:write", data, opts)
    return err
}

// command "getmeta", wshserver.GetMetaCommand
func GetMetaCommand(w *wshutil.WshRpc, data wshrpc.CommandGetMetaData, opts *wshrpc.WshRpcCommandOpts) (map[string]interface {}, error) {
    resp, err := sendRpcRequestCallHelper[map[string]interface {}](w, "getmeta", data, opts)
    return resp, err
}

// command "message", wshserver.MessageCommand
func MessageCommand(w *wshutil.WshRpc, data wshrpc.CommandMessageData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "message", data, opts)
    return err
}

// command "resolveids", wshserver.ResolveIdsCommand
func ResolveIdsCommand(w *wshutil.WshRpc, data wshrpc.CommandResolveIdsData, opts *wshrpc.WshRpcCommandOpts) (wshrpc.CommandResolveIdsRtnData, error) {
    resp, err := sendRpcRequestCallHelper[wshrpc.CommandResolveIdsRtnData](w, "resolveids", data, opts)
    return resp, err
}

// command "setmeta", wshserver.SetMetaCommand
func SetMetaCommand(w *wshutil.WshRpc, data wshrpc.CommandSetMetaData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "setmeta", data, opts)
    return err
}

// command "setview", wshserver.BlockSetViewCommand
func BlockSetViewCommand(w *wshutil.WshRpc, data wshrpc.CommandBlockSetViewData, opts *wshrpc.WshRpcCommandOpts) error {
    _, err := sendRpcRequestCallHelper[any](w, "setview", data, opts)
    return err
}

// command "streamtest", wshserver.RespStreamTest
func RespStreamTest(w *wshutil.WshRpc, opts *wshrpc.WshRpcCommandOpts) chan wshrpc.RespOrErrorUnion[int] {
    return sendRpcRequestResponseStreamHelper[int](w, "streamtest", nil, opts)
}


