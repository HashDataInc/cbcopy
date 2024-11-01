package builtin

import (
	"github.com/cloudberrydb/cbcopy/internal/dbconn"
	"github.com/cloudberrydb/cbcopy/meta/builtin/toc"
	"github.com/cloudberrydb/cbcopy/options"
	"github.com/spf13/pflag"
	"strings"
)

// gpbackup, backup/global_variables.go

/*
 * Empty struct type used for value with 0 bytes
 */
type Empty struct{}

var (
	gpdbVersion    dbconn.GPDBVersion
	gpdbVersionDst dbconn.GPDBVersion
	hdwVersion     dbconn.GPDBVersion
	hdwVersionDst  dbconn.GPDBVersion

	// --> automation test purpose, start
	// current code many place has connectionPool, which is not very clear how it's used, let's comment it out to see effect.
	// connectionPool *dbconn.DBConn
	connectionPoolDst *dbconn.DBConn
	cmdFlags          *pflag.FlagSet
	// <-- automation test purpose, end

	globalTOC            *toc.TOC
	quotedRoleNames      map[string]string
	filterRelationClause string
	excludeRelations     []string
	includeRelations     []string
	includeSchemas       []string
	excludeSchemas       []string
	objectCounts         map[string]int

	errorTablesMetadata map[string]Empty
	errorTablesData     map[string]Empty
	needConvert         bool
	relevantDeps        DependencyMap
	redirectSchema      map[string]string
	inclDestSchema      string
	ownerMap            map[string]string
)

// --> automation test purpose, start
func SetCmdFlags(flagSet *pflag.FlagSet) {
	cmdFlags = flagSet
}

func SetConnection(conn *dbconn.DBConn) {
	// current code many place has connectionPool, which is not very clear how it's used, let's comment it out to see effect.
	// connectionPool = conn
	gpdbVersion = conn.Version
}

func SetQuotedRoleNames(quotedRoles map[string]string) {
	quotedRoleNames = quotedRoles
}

func SetFilterRelationClause(filterClause string) {
	filterRelationClause = filterClause
}

func SetSchemaFilter(optionName string, optionValue string) {
	if optionName == options.SCHEMA {
		if strings.TrimSpace(optionValue) != "" {
			includeSchemas = append(includeSchemas, optionValue)
		} else {
			includeSchemas = includeSchemas[0:0]
		}
	} else if optionName == options.EXCLUDE_SCHEMA {
		if strings.TrimSpace(optionValue) != "" {
			excludeSchemas = append(excludeSchemas, optionValue)
		} else {
			excludeSchemas = excludeSchemas[0:0]
		}
	}
}

func SetRelationFilter(optionName string, optionValue string) {
	if optionName == options.INCLUDE_TABLE {
		if strings.TrimSpace(optionValue) != "" {
			includeRelations = append(includeRelations, optionValue)
		} else {
			includeRelations = includeRelations[0:0]
		}
	} else if optionName == options.EXCLUDE_TABLE {
		if strings.TrimSpace(optionValue) != "" {
			excludeRelations = append(excludeRelations, optionValue)
		} else {
			excludeRelations = excludeRelations[0:0]
		}
	}
}

// <-- automation test purpose, end
