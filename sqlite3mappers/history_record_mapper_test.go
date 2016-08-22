package sqlite3mappers

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/mappers"
)

const testTransactions = `
INSERT INTO inflow (document_guid, unixtimestamp, name, amount, source)
VALUES
("203481d2-8026-4613-ae27-c102fb6685df", 1,  "test transaction 1",  10.20, "src"),
("4033f520-dee6-43a0-94d1-4c24dccf3ea3", 3,  "test transaction 3",  5.20,  "src"),
("be8de6ee-c44d-4e0b-95d5-9c6d3f47b7ea", 5,  "test transaction 5",  2.20,  "src"),
("55e899aa-d259-4aa4-9485-715bc43cc63d", 7,  "test transaction 7",  7.50,  "src"),
("b5755d96-7eb3-4310-a780-a3a137b0ec8c", 9,  "test transaction 9",  9.20,  "src"),
("ab56893c-0015-4c3a-bbfc-27d5c22be899", 11, "test transaction 11", 11.20, "src"),
("980e8a00-e543-4668-be4c-cbb7ea01aea0", 13, "test transaction 13", 2.20,  "src"),
("9ec4e2fd-bac5-431e-a3ba-c552a43180de", 15, "test transaction 15", 20.20, "src"),
("c8f81266-ebe3-414a-aeb2-16f35048c994", 17, "test transaction 17", 5.20,  "src"),
("4041b245-c626-4db8-90c8-b761d73a3902", 19, "test transaction 19", 10.20, "src");

INSERT INTO outflow (document_guid, unixtimestamp, name, amount, destination)
VALUES
("11e1553f-5e33-4e9f-a69a-4f645e24ec88", 2,  "test transaction 2",  1.00, "dst"),
("d5689382-2ee9-47f7-bd43-6443be5463bd", 4,  "test transaction 4",  1.00, "dst"),
("ec0462ad-e119-460b-92af-d6ba710057c4", 6,  "test transaction 6",  0.50, "dst"),
("c21c17d2-b5fe-47cd-a3ec-4f01d9ade765", 8,  "test transaction 8",  5.00, "dst"),
("c6e51010-8218-40dc-93bb-c65507228e86", 10, "test transaction 10", 2.00, "dst"),
("14d2e901-4c4b-4594-90ca-df5da8351ffc", 12, "test transaction 12", 1.50, "dst"),
("725f05a5-70b8-4c4e-aadd-032a71fca961", 14, "test transaction 14", 3.00, "dst"),
("cdca1120-fb0a-4086-8d10-22e8418ab65d", 16, "test transaction 16", 5.58, "dst"),
("8ece76aa-1381-4389-b79f-5e992cea5f76", 18, "test transaction 18", 1.10, "dst"),
("a359b83f-a720-44d8-af17-b7234e2c6556", 20, "test transaction 20", 8.10, "dst");
`

const (
	transactionCount = 20
	finishBalance    = 54.52
)

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}

func TestHistoryRecordMapper(t *testing.T) {
	os.Remove(testDBName)
	db, err := InitSQLiteDB(testDBName)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(testDBName)
	}()

	_, err = db.Exec(testTransactions)
	require.Nil(t, err)

	historyRecordMapper, err := NewHistoryRecordMapper(db)
	require.Nil(t, err)

	_, err = historyRecordMapper.GetHistoryRecordsByTimeRange(time.Unix(0, 0), time.Unix(0, 0))
	require.NotNil(t, err)
	require.Equal(t, mappers.ErrGetHistoryRecordsByTimeRangeInvalidTimeRange, err)

	records, err := historyRecordMapper.GetHistoryRecordsByTimeRange(time.Unix(0, 0), time.Unix(21, 0))
	require.Nil(t, err)
	require.Equal(t, len(records), transactionCount)

	var balance float64
	for _, record := range records {
		balance += record.Amount
	}

	balance = round(balance, 2)
	require.Equal(t, finishBalance, balance)
	require.Equal(t, finishBalance, records[0].Balance)
}
