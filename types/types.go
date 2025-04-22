package types

import (
	"context"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

type Plugin struct {
	MCPType    string `json:"mcp_type"`
	Name       string `json:"name"`
	VarName    string `json:"var_name"`
	PluginFile string `json:"plugin_file"`
}

// Tool is mcp tool for the MCP server,
// it can be loadded as a plugin and add to server dynamically
type MCPTool struct {
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Option  []mcp.ToolOption                                                                    `json:"-"`
	Handler func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) `json:"-"`
}
