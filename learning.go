/* package main
hello go

import "fmt"

func main(){

	fmt.Println("Hello world")
} */


// control shift a yoruma alma
/*
Values
package main

import "fmt"

 func main(){

	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0 =",7.0/3.0)

	fmt.Println(true && false)

	fmt.Println(true || false)

	fmt.Println(!true)
}
*/

/*
 Variables
package main

import "fmt"


func main(){

var a= "initial"
fmt.Println(a)

var b, c int =1,2

fmt.Println(b,c)

var d =true
fmt.Println(d)

var e int
fmt.Println(e)

f:= "apple"
fmt.Println(f)
}

*/
/*
Constants
package main

import (
	"fmt"
	"math"
)

// we can declare const variables out side of main we can declare that in every where

const s string = "Constants"

func main() {

	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
*/
/*
package main
// for loops

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {

		fmt.Println("Loop")
		break
	}

	for n := 0; n <= 5; n++ {

		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
*/

/* package main
// if else
import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println(" 7 is   odd")
	}

	if 8%4 == 0 {

		fmt.Println(" 8 divisible by 4")
	}
	if num := 9; num < 0 {

		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, "has mutiple digits")
	}
} */
/*
package main

// switch case kullanımı ayrıca time kullanımı

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as")
	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")

	default:
		fmt.Println("It s a week day")

	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println(" biz salak ingiliz/amerikalılar buna before noon diyoz ")
	default:
		fmt.Println("biz salak ingiliz/amerikalılar bunada after noon diyoz")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf(" I dont know my type its should be %T", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey\n")

}
*/
/*
package main
arrays

import (
	"fmt"
)

func main() {

	var a [5]int
	fmt.Println("emp :", a)
	a[4] = 100
	fmt.Println("set : ", a[4])

	fmt.Println("Len :", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl : ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
/*
package main

import (
	"fmt"
)

func main() {

	s := make([]string, 3)
	fmt.Println("emp", 5)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len of s", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apended s: ", s)
	fmt.Println("len of s", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Coppied s c slices:", c)

	l := s[2:5]
	fmt.Println(" l is 2:5 of s what is output  sonuncuyu görmüyor 5 diyince: ", l)
	// ikinci kere değer atmaswı yaparken bunu ":=" ekleme  "=" bunu ekle sadece
	l = s[:5]
	fmt.Println(" l is :5 of s what is output  :  ", l)

	l = s[2:]
	fmt.Println(" l is 2: of s what is output  :  ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("DCL :", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(" 2 dimensional \n:", twoD)

}
*/

/*
package main
// map yapısı

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("Map : ", m)
	v1 := m["k1"]
	fmt.Println("v1 : ", v1)
	fmt.Println("Len : ", len(m))
	delete(m, "k2")

	fmt.Println("map :", m)

	_, prs := m["k2"]
	fmt.Println("prs : ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n map :", n)

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum : ", sum)

	for i, num := range nums {

		if num == 3 {
			fmt.Println("İndex :", i)
		}

	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {

		fmt.Printf("%s -> %s \n", k, v)

	}

	for k := range kvs {
		fmt.Println("Key :", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

} */

/* package main

import (
	"fmt"
)


func plus(a int, b int)int {
	return a+b
}

func plusPlus(a,b,c int)int  {

	return a+b+c
}

// functions
func main() {
	res:= plus(1,2)
	fmt.Println("1+2 := ",res)
	res= plusPlus(1,2,3)
	fmt.Println("1+2+3 := ",res)
} */

/* package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

// multiple functions return dual
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, v := vals()
	fmt.Println(v)
	d, _ := vals()
	fmt.Println(d)
	//d := vals()  hata tekli alınmıyor alcaksan blank space koy
	fmt.Println(d)
}
*/
/*
package main
// variadic değişken functions slice range of array gez der gibi dizide
import (
	"fmt"
)

func sum(numbers ...int) {
	fmt.Print(numbers, " =: ")
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum(nums...)
}
*/
/*
package main

// Closures
import (
	"fmt"
)

func intSeq() func() int {

	i := 0

	return func() int {
		i++
		return i
	}

}

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt)

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts)

	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts)

} */

/* package main
// recurisive function recursive
import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1

	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(9))

} */
/*
package main

// pointers geldi yine kanser C olayları  eski referans değerleri pointerlar ile alabilir ve işleyebiliriiz Go pointerları destekliyor

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial  value : ", i)
	zeroval(i)
	// zeroval foknsiyon olarak  main icinden i' yi değiştirmedi fakaatttt
	fmt.Println("Zero value : ", i)
	// zeroptr değiştirdi cünkü bu değer için memoryde adresi var
	// bunu kullanarak main icindeki değerleri pointerlar yardımı ile değiştirebiliriz
	zeroptr(&i)
	fmt.Println("Zeroptr : ", i)
	fmt.Println("Pointer : ", &i)
}
*/
/* package main
// structs
import "fmt"

type person struct {
	name string
	age  int
}
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))
	x := newPerson("Şeyma")
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(x.age)
	x.age = 56
	fmt.Println(x.age)

}
*/
/*
package main

// METHODS

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area : ", r.area())

	fmt.Println("Area : ", r.perim())

	rp := &r
	fmt.Println("Area : ", rp.area())
	fmt.Println("Perim : ", rp.perim())
}
*/

/* package main

// method paketleri koleksiyon methodlar uno ekmek bu da şok poşeti :)

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {

	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
} */
/*
package main
// errors
import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cant work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {

	return fmt.Sprintf("%d -%s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)

	}
}
*/
/*
package main
// go routines

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {

		fmt.Println(from, ":", i)
	}

}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {

		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done .. ^^")

}
*/
/*
package main

// channels

import (
	"fmt"
)

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)
}


*/

/* package main

// buffered channels
import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)
	messages <- "buffered"
	messages <- "channel"
	messages <- "kartelam"
	fmt.Println(<-messages)

	fmt.Println(<-messages)

	fmt.Println(<-messages)

}
*/

/*
package main
// sychnornized channels
import (
	"fmt"
	"time"
)

func worker(done chan bool) {

	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done

}
*/
/*
package main

// channels directions
import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
*/

/* package main

// select is powerful feature of go
import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)

		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
*/
/*
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	c2 := make(chan string)

	select {
	case res := <-c1:
		fmt.Println(res)
		fmt.Println("time out 1")
	}

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "one"
	}()

}
*/
/*
package main
// time outs
import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("time out 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("time out 2")
	}
}
*/
/*
package main
//non blocking channel
import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message : ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message : ", msg)
	default:
		fmt.Println(" no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signals", sig)
	default:
		fmt.Println("no activity")
	}
}
*/
/*
package main
// closing channels close channels
import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {

		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println(" sent job ", j)
	}
	close(jobs) // closing channel how to
	fmt.Println("sent all jobs")
	<-done
}
*/
/*
package main
// range over range rover channels :D
import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// kapatmazsan mesela bu satırı forun altına alırsan patlar deadlock error alırız...int

	for elem := range queue {
		fmt.Println(elem)
	}

}
*/
/*
package main

//timers
import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)

}
*/
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}
*/
/*
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}

}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
*/

/* package main

import (
	"fmt"
	"sync"
	"time"
)

// to wait for multiple go rouitnes to finish, we can use a wait group

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker done %d done\n", id)
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
*/
/*
package main
// rate limiting
import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request :", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}

	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {

		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
*/
/*
package main
//atomic waiters
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {

		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)

			}
			wg.Done()
		}()

		wg.Wait()
		fmt.Println("ops:", ops)
	}
}
*/
/*
package main
// mutex
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {

		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}

		}()
	}

	for w := 0; w < 10; w++ {

		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)

			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

}
*/
/* 
  package main

 import (
	 "fmt"
	 "math/rand"
	 "sync/atomic"
	 "time"
	)

	type readOp struct{
		key int
		resp chan int
	}


	type writeOp struct{
		key int
		val int
		resp chan bool
	}

	func main() {
		var readOps uint64
		var writeOps uint64

		reads:= make(chan readOp)
		writes:=make(chan writeOp)


		go func() {

			var state=make(map[int]int)
			for{
				select {
				case read:= <-reads:
					read.resp <-state[read.key]



				}
			}
		}

	} */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://freshman.tech/snippets/go/http-response-to-string/", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier

	if err != nil {

		log.Fatalln(err)

	}

	fmt.Println(string(b))
}