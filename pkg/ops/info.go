package ops

import (
	"fmt"
	"io"

	"github.com/openark/golib/log"

	"github.com/cohenjo/dbender/pkg/config"
	"github.com/cohenjo/dbender/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CheckLocks(dbhost string) ([]types.InnoLock, error) {
	cmd := "SELECT waiting_trx_id, waiting_pid, waiting_query, blocking_trx_id, blocking_pid, blocking_query FROM sys.innodb_lock_waits;"
	locks := []types.InnoLock{}
	err := RunCmd(&locks, dbhost, cmd)
	return locks, err
}

func RunCmd(dest interface{}, dbhost string, cmd string) error {
	port := 3306
	user := config.Config.User
	passwd := config.Config.Password
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/test", user, passwd, dbhost, port)
	db := sqlx.MustOpen("mysql", connURL)
	defer db.Close()

	err := db.Select(dest, cmd)
	return err
}

func StatementAnalysis(dbhost string) ([]types.InnoLock, error) {
	cmd := "select digest,query, db, full_scan,exec_count, total_latency, max_latency, avg_latency, lock_latency, rows_sent, rows_examined, rows_sorted, last_seen from sys.x$statement_analysis limit 20;"
	locks := []types.InnoLock{}
	err := RunCmd(&locks, dbhost, cmd)
	return locks, err
}

func Report(dbhost string, w *io.PipeWriter) {
	defer w.Close()
	// 	-- top 95
	// select SQL_TEXT  from sys.statements_with_runtimes_in_95th_percentile fts,performance_schema.events_statements_history shist where fts.digest = shist.digest\G
	cmd := "select * from sys.statements_with_runtimes_in_95th_percentile where db is not null and db not in ('sys','performance_schema') and exec_count > 10 and last_seen > date_sub(now() , interval 7 DAY);"
	// cmd := "select * from statements_with_runtimes_in_95th_percentile ;"
	srp := []types.StatementsWithRuntimesIn95thPercentile{}
	_ = RunCmd(&srp, dbhost, cmd)

	n, err := fmt.Fprintf(w, "################################################\n")
	if err != nil {
		log.Error("problem writeing to w, \n", err)
	}
	log.Debugf("wrote %d chars", n)
	n, err = fmt.Fprintf(w, "Statments in the 95%% seen in the last week:\n")
	if err != nil {
		log.Error("problem writeing to w, \n", err)
	}
	log.Debugf("wrote %d chars", n)
	for _, s := range srp {
		fmt.Fprintf(w, "%s\n", s)
	}
	fmt.Fprintf(w, "################################################\n")

	// -- FTS
	// select SQL_TEXT  from sys.statements_with_full_table_scans fts,performance_schema.events_statements_history shist where  fts.digest = shist.digest\G
	cmd = "select * from sys.statements_with_full_table_scans ;"
	sft := []types.StatementsWithFullTableScans{}
	_ = RunCmd(&sft, dbhost, cmd)

	fmt.Fprintf(w, "################################################\n")
	fmt.Fprintf(w, "## Statments with FTS:\n")
	for _, s := range sft {
		fmt.Fprintf(w, "%s\n", s)
	}
	n, err = fmt.Fprintf(w, "################################################\n")
	log.Debugf("wrote %d chars", n)

	// -- Not used indexes
	// select * from sys.schema_unused_indexes;

	cmd = "select * from sys.schema_unused_indexes;"
	sui := []types.SchemaUnusedIndexes{}
	_ = RunCmd(&sui, dbhost, cmd)
	fmt.Fprintf(w, "################################################\n")
	fmt.Fprintf(w, "Unused Indexes:\n")
	for _, unusedIndex := range sui {
		fmt.Fprintf(w, "%s\n", unusedIndex)
	}
	fmt.Fprintf(w, "################################################\n")

	// -- table stat
	// select * from sys.schema_table_statistics;
	sts := []types.SchemaTableStatistics{}
	cmd = "select * from sys.schema_table_statistics;"
	_ = RunCmd(&sts, dbhost, cmd)
	fmt.Fprintf(w, "################################################\n")
	fmt.Fprintf(w, "Table_schema.Table_name.Index_name: s.Insert_latency | s.Select_latency | s.Rows_selected:\n")
	for _, unusedIndex := range sts {
		fmt.Fprintf(w, "%s\n", unusedIndex)
	}
	fmt.Printf("################################################\n")
	// -- index stat
	// select * from sys.x$schema_index_statistics;
	sis := []types.SchemaIndexStatistics{}
	cmd = "select * from sys.x$schema_index_statistics;"
	_ = RunCmd(&sis, dbhost, cmd)
	fmt.Fprintf(w, "################################################\n")
	fmt.Fprintf(w, "Table_schema.Table_name.Index_name: s.Insert_latency | s.Select_latency | s.Rows_selected:\n")
	for _, unusedIndex := range sis {
		fmt.Fprintf(w, "%s\n", unusedIndex)
	}
	fmt.Fprintf(w, "################################################\n")

	// -- query using temp tables on disk
	// select * from sys.statements_with_temp_tables where disk_tmp_tables > 1;
	swt := []types.StatementsWithTempTables{}
	cmd = "select * from sys.statements_with_temp_tables where disk_tmp_tables > 1;"
	_ = RunCmd(&swt, dbhost, cmd)
	fmt.Fprintf(w, "################################################\n")
	fmt.Fprintf(w, "Table_schema.Table_name.Index_name: s.Insert_latency | s.Select_latency | s.Rows_selected:\n")
	for _, unusedIndex := range swt {
		fmt.Fprintf(w, "%s\n", unusedIndex)
	}
	fmt.Fprintf(w, "################################################\n")

	// -- query with errors
	// select * from sys.statements_with_errors_or_warnings order by exec_count desc,errors desc;

	swtew := []types.StatementsWithErrorsWarnings{}
	cmd = "select * from sys.statements_with_temp_tables where disk_tmp_tables > 1;"
	_ = RunCmd(&swtew, dbhost, cmd)
	fmt.Fprintf(w, "################################################\n")
	fmt.Fprintf(w, "Table_schema.Table_name.Index_name: s.Insert_latency | s.Select_latency | s.Rows_selected:\n")
	for _, unusedIndex := range swtew {
		fmt.Fprintf(w, "%s\n", unusedIndex)
	}
	fmt.Fprintf(w, "################################################\n")

}
