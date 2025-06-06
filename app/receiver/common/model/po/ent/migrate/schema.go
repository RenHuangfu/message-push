// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BusinessAppsColumns holds the columns for the "business_apps" table.
	BusinessAppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "app_key", Type: field.TypeString, Size: 255},
		{Name: "app_id", Type: field.TypeString, Size: 255},
		{Name: "app_type", Type: field.TypeString, Size: 255},
		{Name: "is_del", Type: field.TypeInt},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// BusinessAppsTable holds the schema information for the "business_apps" table.
	BusinessAppsTable = &schema.Table{
		Name:       "business_apps",
		Columns:    BusinessAppsColumns,
		PrimaryKey: []*schema.Column{BusinessAppsColumns[0]},
	}
	// BusinessClientsColumns holds the columns for the "business_clients" table.
	BusinessClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "client_key", Type: field.TypeString, Size: 255},
		{Name: "app_id", Type: field.TypeString, Size: 255},
		{Name: "client_id", Type: field.TypeString, Size: 255},
		{Name: "client_type", Type: field.TypeString, Size: 255},
		{Name: "is_del", Type: field.TypeInt},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// BusinessClientsTable holds the schema information for the "business_clients" table.
	BusinessClientsTable = &schema.Table{
		Name:       "business_clients",
		Columns:    BusinessClientsColumns,
		PrimaryKey: []*schema.Column{BusinessClientsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "app_id", Type: field.TypeInt},
		{Name: "client_ids", Type: field.TypeJSON},
		{Name: "content", Type: field.TypeString},
		{Name: "delay", Type: field.TypeInt},
		{Name: "send_time", Type: field.TypeTime},
		{Name: "send_count", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt},
		{Name: "is_del", Type: field.TypeInt},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BusinessAppsTable,
		BusinessClientsTable,
		MessagesTable,
	}
)

func init() {
}
