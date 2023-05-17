# Project "local_mocker"

Project developed in GO as a local mock

---

### Objectives:

* Local mock for develpment
* Read files to file system
* Rest API to return Json File


---

### Upgrade Plan
| Description                                     | Done?                 | Date       |
|-------------------------------------------------|-----------------------|------------|
| Include HttpStatus 5XX                          | :heavy_check_mark:    | 17/05/2023 |
| Create unit tests                               | :black_square_button: |            |


---

### Local execution

* Run command: "go run main.go"
* Url call: http://localhost:3000/?code=400

**"code" is the http status you want returned.**


### References
* https://go.dev/dl/ (local go installation)
* https://go.dev/src/net/http/status.go
* https://golangbyexample.com/set-http-status-code-golang/
