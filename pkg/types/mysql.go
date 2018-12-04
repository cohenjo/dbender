package types

import (
	"database/sql"
	"fmt"
	"time"
)

type InnoLock struct {
	WaitingTrxId   string         `db:"waiting_trx_id"`
	WaitingPid     int64          `db:"waiting_pid"`
	WaitingQuery   sql.NullString `db:"waiting_query"`
	BlockingTrx_id string         `db:"blocking_trx_id"`
	BlockingPid    int64          `db:"blocking_pid"`
	BlockingQuery  sql.NullString `db:"blocking_query"`
}

//digest,query, db, full_scan,exec_count, total_latency, max_latency, avg_latency, lock_latency, rows_sent, rows_examined, rows_sorted, last_seen
type StmntAnalysis struct {
	digest        string         `db:"waiting_trx_id"`
	query         int64          `db:"waiting_pid"`
	db            sql.NullString `db:"waiting_query"`
	full_scan     string         `db:"blocking_trx_id"`
	exec_count    int64          `db:"blocking_pid"`
	total_latency sql.NullString `db:"blocking_query"`
	max_latency   sql.NullString `db:"waiting_query"`
	avg_latency   string         `db:"blocking_trx_id"`
	lock_latency  int64          `db:"blocking_pid"`
	rows_sent     sql.NullString `db:"blocking_query"`
	rows_examined string         `db:"blocking_trx_id"`
	rows_sorted   int64          `db:"blocking_pid"`
	last_seen     sql.NullString `db:"blocking_query"`
}

// Schema_unused_indexes is the structure of the home table
type SchemaUnusedIndexes struct {
	Object_schema string `column:"object_schema" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Object_name   string `column:"object_name" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Index_name    string `column:"index_name" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
}

type StatementsWithRuntimesIn95thPercentile struct {
	Query             string         `column:"query" default:"" type:"longtext" key:"" null:"YES" extra:""`
	Db                sql.NullString `column:"db" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Full_Scan         string         `column:"full_scan" default:"" type:"varchar(1)" key:"" null:"NO" extra:""`
	Exec_Count        string         `column:"exec_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Err_Count         string         `column:"err_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Warn_Count        string         `column:"warn_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Total_Latency     string         `column:"total_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Max_Latency       string         `column:"max_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Avg_Latency       string         `column:"avg_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Rows_Sent         string         `column:"rows_sent" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_Sent_Avg     float64        `column:"rows_sent_avg" default:"0" type:"decimal(21,0)" key:"" null:"NO" extra:""`
	Rows_Examined     string         `column:"rows_examined" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_Examined_Avg float64        `column:"rows_examined_avg" default:"0" type:"decimal(21,0)" key:"" null:"NO" extra:""`
	First_Seen        time.Time      `column:"first_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Last_Seen         time.Time      `column:"last_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Digest            string         `column:"digest" default:"" type:"varchar(32)" key:"" null:"YES" extra:""`
}

// Statements_with_full_table_scans is the structure of the home table
type StatementsWithFullTableScans struct {
	Query                    *string   `column:"query" default:"" type:"longtext" key:"" null:"YES" extra:""`
	Db                       *string   `column:"db" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Exec_count               string    `column:"exec_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Total_latency            *string   `column:"total_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	No_index_used_count      string    `column:"no_index_used_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	No_good_index_used_count string    `column:"no_good_index_used_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	No_index_used_pct        float64   `column:"no_index_used_pct" default:"0" type:"decimal(24,0)" key:"" null:"NO" extra:""`
	Rows_sent                string    `column:"rows_sent" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_examined            string    `column:"rows_examined" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_sent_avg            *float64  `column:"rows_sent_avg" default:"" type:"decimal(21,0) unsigned" key:"" null:"YES" extra:""`
	Rows_examined_avg        *float64  `column:"rows_examined_avg" default:"" type:"decimal(21,0) unsigned" key:"" null:"YES" extra:""`
	First_seen               time.Time `column:"first_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Last_seen                time.Time `column:"last_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Digest                   *string   `column:"digest" default:"" type:"varchar(32)" key:"" null:"YES" extra:""`
}

// X$schema_index_statistics is the structure of the home table
type SchemaIndexStatistics struct {
	Table_schema   string `column:"table_schema" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Table_name     string `column:"table_name" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Index_name     string `column:"index_name" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Rows_selected  string `column:"rows_selected" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Select_latency string `column:"select_latency" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_inserted  string `column:"rows_inserted" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Insert_latency string `column:"insert_latency" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_updated   string `column:"rows_updated" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Update_latency string `column:"update_latency" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Rows_deleted   string `column:"rows_deleted" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Delete_latency string `column:"delete_latency" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
}

type SchemaTableStatistics struct {
	Table_schema      string  `column:"table_schema" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Table_name        string  `column:"table_name" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Total_latency     string  `column:"total_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Rows_fetched      string  `column:"rows_fetched" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Fetch_latency     string  `column:"fetch_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Rows_inserted     string  `column:"rows_inserted" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Insert_latency    string  `column:"insert_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Rows_updated      string  `column:"rows_updated" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Update_latency    string  `column:"update_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Rows_deleted      string  `column:"rows_deleted" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Delete_latency    string  `column:"delete_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Io_read_requests  float64 `column:"io_read_requests" default:"" type:"decimal(42,0)" key:"" null:"YES" extra:""`
	Io_read           string  `column:"io_read" default:"" type:"text" key:"" null:"YES" extra:""`
	Io_read_latency   string  `column:"io_read_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Io_write_requests float64 `column:"io_write_requests" default:"" type:"decimal(42,0)" key:"" null:"YES" extra:""`
	Io_write          string  `column:"io_write" default:"" type:"text" key:"" null:"YES" extra:""`
	Io_write_latency  string  `column:"io_write_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Io_misc_requests  float64 `column:"io_misc_requests" default:"" type:"decimal(42,0)" key:"" null:"YES" extra:""`
	Io_misc_latency   string  `column:"io_misc_latency" default:"" type:"text" key:"" null:"YES" extra:""`
}

type StatementsWithTempTables struct {
	Query                    string    `column:"query" default:"" type:"longtext" key:"" null:"YES" extra:""`
	Db                       string    `column:"db" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Exec_count               string    `column:"exec_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Total_latency            string    `column:"total_latency" default:"" type:"text" key:"" null:"YES" extra:""`
	Memory_tmp_tables        string    `column:"memory_tmp_tables" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Disk_tmp_tables          string    `column:"disk_tmp_tables" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Avg_tmp_tables_per_query float64   `column:"avg_tmp_tables_per_query" default:"0" type:"decimal(21,0)" key:"" null:"NO" extra:""`
	Tmp_tables_to_disk_pct   float64   `column:"tmp_tables_to_disk_pct" default:"0" type:"decimal(24,0)" key:"" null:"NO" extra:""`
	First_seen               time.Time `column:"first_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Last_seen                time.Time `column:"last_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Digest                   string    `column:"digest" default:"" type:"varchar(32)" key:"" null:"YES" extra:""`
}

type StatementsWithErrorsWarnings struct {
	Query       string    `column:"query" default:"" type:"longtext" key:"" null:"YES" extra:""`
	Db          string    `column:"db" default:"" type:"varchar(64)" key:"" null:"YES" extra:""`
	Exec_count  string    `column:"exec_count" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Errors      string    `column:"errors" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Error_pct   float64   `column:"error_pct" default:"0.0000" type:"decimal(27,4)" key:"" null:"NO" extra:""`
	Warnings    string    `column:"warnings" default:"" type:"bigint(20) unsigned" key:"" null:"NO" extra:""`
	Warning_pct float64   `column:"warning_pct" default:"0.0000" type:"decimal(27,4)" key:"" null:"NO" extra:""`
	First_seen  time.Time `column:"first_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Last_seen   time.Time `column:"last_seen" default:"0000-00-00 00:00:00" type:"timestamp" key:"" null:"NO" extra:""`
	Digest      string    `column:"digest" default:"" type:"varchar(32)" key:"" null:"YES" extra:""`
}

func (l InnoLock) String() string {
	bq := "N/A"
	if l.BlockingQuery.Valid {
		if len(l.BlockingQuery.String) > 30 {
			bq = l.BlockingQuery.String[0:30] + "..."
		}
	}
	return fmt.Sprintf("trx: %s(%d) is locked by %s(%d), blocking query: %s", l.WaitingTrxId, l.WaitingPid, l.BlockingTrx_id, l.BlockingPid, bq)
}

func (u SchemaUnusedIndexes) String() string {
	return fmt.Sprintf("%s.%s.%s", u.Object_schema, u.Object_name, u.Index_name)
}

func (s StatementsWithErrorsWarnings) String() string {
	return fmt.Sprintf("%s | %s | %s", s.Query, s.Exec_count, s.Digest)
}
func (s StatementsWithTempTables) String() string {
	return fmt.Sprintf("%s | %s | %s", s.Query, s.Exec_count, s.Digest)
}

func (s StatementsWithRuntimesIn95thPercentile) String() string {
	return fmt.Sprintf("%s | %s | %s", s.Query, s.Exec_Count, s.Digest)
}

func (s StatementsWithFullTableScans) String() string {
	return fmt.Sprintf("%s | %s | %s", s.Query, s.Exec_count, s.Digest)
}

func (s SchemaIndexStatistics) String() string {
	return fmt.Sprintf("%s.%s.%s: %s | %s | %s", s.Table_schema, s.Table_name, s.Index_name, s.Insert_latency, s.Select_latency, s.Rows_selected)
}

func (s SchemaTableStatistics) String() string {
	return fmt.Sprintf("%s.%s: %s | %s | %s | %s", s.Table_schema, s.Table_name, s.Total_latency, s.Rows_updated, s.Rows_deleted, s.Io_read_latency)
}
