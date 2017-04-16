package main

/*
#include "postgres.h"
#include "commands/trigger.h"
#include "utils/elog.h"
#include "utils/rel.h"
#include "access/htup_details.h"

static int trigger_fired_by_update(TriggerEvent tg_event) {
	return (TRIGGER_FIRED_BY_UPDATE(tg_event)) != 0;
}

static Datum pointer_get_datum(HeapTuple t) {
	return PointerGetDatum(t);
}

static char *getarg_text(TriggerData *trigdata, HeapTuple rettuple, int idx) {
	bool isnull;
	TupleDesc tupdesc = trigdata->tg_relation->rd_att;
	text * t = DatumGetTextP(heap_getattr(rettuple, idx, tupdesc, &isnull));
	if (isnull || !t) {
		return "";
	}
	return VARDATA(t);
}

static void elog_info(char *s) {
	elog(INFO, "%s", s);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export mytrigger
func mytrigger(fcInfo *C.FunctionCallInfoData) C.Datum {
	trigdata := (*C.TriggerData)(unsafe.Pointer(fcInfo.context))

	var rettuple *C.HeapTupleData
	if C.trigger_fired_by_update(trigdata.tg_event) != 0 {
		rettuple = (*C.HeapTupleData)(trigdata.tg_newtuple)
	} else {
		rettuple = (*C.HeapTupleData)(trigdata.tg_trigtuple)
	}

	url := C.GoString(C.getarg_text(trigdata, rettuple, 1))

	C.elog_info(C.CString(fmt.Sprintf("got url=%s", url)))
	fmt.Println(url)

	return C.pointer_get_datum(rettuple)
}
