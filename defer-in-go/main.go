/* Deferred Functions
Deferred functions are something that are used a lot within Go when it comes to things like setting up connections to say,
 databases, or brokers. Typically we end up using a deferred function to ensure that connections are closed after we are
  done with them so we don’t have stale connections to the database.
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
)

type Engineer struct {
	Name string
}

func (e *Engineer) UpdateName(name string) {
	e.Name = name
}
func Crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// procces to rest of the body
	// do other stuff with the response
	// go nun sağladığı bu özellik sayesinde requesti atıp hemen kapıyoruz ve connection eski kalmıyor hep yenileniyor
}

/*


In this example, we’ve grouped the code to perform a http request as well as clean up and close our response body together.
The alternative to this would be to process the response body and then remember to close the body at the end of the function
 but this means that we’ve effectively scattered the code to make this request across the entire body of the function.
Doing it this way makes it easier to remember. When performing these HTTP requests,
 you quickly get into the habit of remembering to close if you always write your http requests using this same method.




*/
func doStuff(e *Engineer) {
	name := "Elliot Forbess"
	defer e.UpdateName(name)
	fmt.Println("Doing other exciting stuff")
	name = "Elliot M forbes"

	// defer ile name i atadığımız için name sonradan değişse bile yine ilk deferla aldığı değeri işliyip atıyor

}
func MyAwesomeFunction() (err error) {

	defer func() { // defer function  set the Error() value
		err = errors.New("I am an error")
	}()

	return nil
}

func ReturnValues() (x int) {

	defer func() {

		x = 10

	}()

	x = 5
	return
}

func main() {

	//   bu durumda sıra önemli ilk yazılan defer son  basılıyor
	defer func() {
		fmt.Println(" I will be executed last")
	}()

	defer func() {

		fmt.Println("Last in -but first out")
	}()

	fmt.Println("defer in go")

	err := MyAwesomeFunction()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ReturnValues())

	elliot := &Engineer{Name: "Elliot"}

	fmt.Printf("%+v\n", elliot)
	doStuff(elliot)
	fmt.Printf("%+v\n", elliot)

}
