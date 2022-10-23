# Project "local_mocker"

Project developed in GO as a local mock

---

### Objectives:

* Local mock for develpment
* Read files to file system
* Rest API do return Json File


---

### Upgrade Plan
| Description                                     | Check?                |
|-------------------------------------------------|-----------------------|
| Include HttpStatus 5XX                          | :black_square_button: |
| Create unit tests                               | :black_square_button: |


---

### Local execution

* Run command: "go run main.go"
* Url call: http://localhost:3000/?code=400

"code" is the http status you want returned.


### References

* https://go.dev/src/net/http/status.go
* https://golangbyexample.com/set-http-status-code-golang/