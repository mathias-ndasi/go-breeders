package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.Parse()

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Server is running on port ", port)

	pattern := regexp.MustCompile("[a-zA-Z0-9]+")
	matched := pattern.FindAll([]byte(`hello`), -1)
	fmt.Println(matched)

	re := regexp.MustCompile(`f.*`)
	if mat := re.Find([]byte(`seafood fool`)); mat != nil {
		fmt.Println("Match found at index", mat[0])
	} else {
		fmt.Println("No match found")
	}
	// fmt.Printf("%q\n", )

	// content := []byte("London")
	// re := regexp.MustCompile(`o.`)
	// fmt.Println(re.FindAllIndex(content, 1))
	// fmt.Println(re.FindAllIndex(content, -1))
	// a := []string{"banana", "orange"}
	// for index, value := range a {
	// 	fmt.Printf("Index: %d, Value: %s\n", index, value)
	// }

	fmt.Println("Hello.....1")
	doneChannel := make(chan bool)
	go slowGreeting(doneChannel)
	select {
	case <-doneChannel:
		fmt.Println("slowGreeting completed")
		break
	case <-time.After(time.Second * 5):
		fmt.Println("slowGreeting timed out")
		break
	default:
		fmt.Println("slowGreeting still running")
		break
	}
	fmt.Println("Hello.....2")

	jobs := make(chan int, 1000)
	resultChannel := make(chan int, 1000)

	go worker(jobs, resultChannel)
	go worker(jobs, resultChannel)
	go worker(jobs, resultChannel)
	go worker(jobs, resultChannel)

	for i := 0; i < 1000; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 1000; j++ {
		fmt.Printf("Result: %d\n", <-resultChannel)
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Panic(error)
	}
}

func slowGreeting(doneChannel chan bool) {
	time.Sleep(time.Second * 7)
	fmt.Println("Hello from slowGreeting!")
	doneChannel <- true
}

func worker(jobs <-chan int, resultChannel chan int) {
	for number := range jobs {
		resultChannel <- fib(number)
	}
}

func fib(number int) int {
	if number <= 1 {
		return number
	}
	return fib(number-1) + fib(number-2)
}
