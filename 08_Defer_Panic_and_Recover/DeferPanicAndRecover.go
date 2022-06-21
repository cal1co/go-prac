package main

import "net/http"

func main() {
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// a := "start"
	// defer fmt.Println(a)
	// a = "end"

	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panic("something bad happened")
	// fmt.Println("end")
}

// Defer, Panic, and Recover

// -- Defer --
// Used to delay execution of a statement until function exits
// Useful to group "open" and "close" functions together
// 	* Be careful in loops
// Run in LIFO order
// Arguments evaluated at time defer is execute, not at time of called function execution

// -- Panic --
// Occur when program cannot continue at all
// 	* Don't use when file can't be opened, unless it is critical
// 	* Use for unrecoverable events - cannot obtain TCP port for web server
// Function will stop executing
// 	* Deferred functions will still fire
// If nothing handles panic, program will exit

// -- Recover --
// Used to recover from panics
// Only useful in deferred functions
// Current function will not attempt to continue, but higher functions in call stack will
