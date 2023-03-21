package webhookLog

var Log *DefaultLogger

func main() {
	// test logging
	Log = NewDefaultLogger("test", "")
	Log.Infof("test")
	Log.Errorf("yo")
}
