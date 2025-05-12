package ToDoListStore

import (
	"context"
	"log"
	"testing"
)

func setupTest(tb testing.TB) func(tb testing.TB) {
	go ProcessDataJobs()
	return func(tb testing.TB) {
		log.Println("done")
	}
}

func TestLoadEntries(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", LoadData, "", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
}

func TestListEntries(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", FetchData, "", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
}

func TestAddEntries(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", AddData, "My test", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
}

func TestDeleteAll(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", DeleteData, "*", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
}

func TestDeleteEntry(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", AddData, "My test", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
	data = DataStoreJob{context.Background(), "simon", DeleteData, "My test", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal = <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
}

func TestUpdateEntry(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", AddData, "My test", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
	data = DataStoreJob{context.Background(), "simon", UpdateData, "My test", "My updated test", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal = <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v", err)
	}
}

func TestPersistEntriess(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	data := DataStoreJob{context.Background(), "simon", StoreData, "", "", make(chan ReturnChannelData)}
	DataJobQueue <- data
	returnVal := <-data.ReturnChannel
	if returnVal.Err != nil {
		t.Errorf("Expected nil got %v adding entry", returnVal.Err)
		return
	}
}
