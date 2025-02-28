package datamanager

type Datahandler interface {
	DataWriter
	DataViwer
}

type DataWriter interface {
	WriteToFile() error
}

type DataViwer interface {
	Display() error
}
