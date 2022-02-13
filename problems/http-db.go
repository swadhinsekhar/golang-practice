
type DeviceDetails struct {
    mac         string
    deviceType  string
}

func (s *server) apiHandler(w http.ResponseWriter, r *http.Request) {

    device := &DeviceDetails{}

    jsonResp := "success"

    mac := r.FormValue("mac")
    if mac == "" {
        w.WriteHeader(404)
        return
    }
    deviceType := r.FormValue("type")
    if deviceType == "" {
        w.WriteHeader(404)
        return
    }

    device.mac = mac
    device.deviceType = deviceType
    s.runChan <- device

    w.Write(jsonResp)

}

// - The object from the request body from API handler
// should be passed to the run goroutine
// - The run goroutine is a long lived background process
// - The run goroutine calls Write when it has to save the
// data to the DB
// - The Write call should be done only when there are more
// 100 objects in the buffer or more than 10 seconds has passed
func (db *DB) Run(runChan chan *DeviceDetails) {

    insertBuf := make([] DeviceDetails, 100)

    cleanUpBufTicker := time.NewTicker(time.Second * time.Duration(10))

    for {
        select {
        case recvBuf, ok <- runChan:

            if !ok { exitFlag = true}

            copy(insertBuf, recvBuf)

            if (len(inserBuf) > 100) {
                db.Write(insertBuf)
            }

        case cleanUpBufTicker:
            if (len(insertBuf)) {
                db.Write(insertBuf)
            }
        }

        if exitFlag == true{
            break
        }
    }
}

// Don't implement this
func (db *DB) Write() {}

type server struct {
    runChan chan *DeviceDetails
}

func main() {

    http.HandleFunc("/", s.apiHandler)
    http.ListenAndServe(":8080", nil)

    s.runChan := make(chan *DeviceDetails, 100)
    go db.Run(s.runChan)
}
