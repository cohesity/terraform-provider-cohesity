// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// SnapshotDeltaProto Message that encapsulates information about delta change to snapshot for
// any of the environments we support. Environment specific snapshot infos are
// defined as extensions to this proto.
//
// Each available extension is listed below along with the location of the
// proto file (relative to magneto/connectors) where it is defined.
// Each adapter should implement a utility function to merge delta snapshot
// to an existing snapshot.
// SnapshotDeltaProto extension             Location                       No.
// =============================================================================
// sql::SnapshotDelta
// ::sql_snapshot_delta                     sql/sql.proto                  101
//
// swagger:model SnapshotDeltaProto
type SnapshotDeltaProto interface{}
