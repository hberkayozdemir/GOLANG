// BENIM ÇIKARIMIM
/*
package main

import (
	"fmt"
	"time"
)
*/

/*
 strong type and compiled language and great concurency
 concurency gave us better speed
 sometimes i/o operation take time and maybe network in isn't stable and something happened is on
or web server off your target domain is slow depend on you simply like increase the time of our responsiveness of your application by waiting
 that web page download
 what ever we have go routunies in go language that simply tell us gives us the ability to write concurency & concurent applications
 it's very simple
 imagine that betwwen two guys I just had



şimdi baba şöyle hayal ediyoruz bizim fmt.Println inimz database den bir şey indirmek

senaryomuzda bizim bir web sitemiz var serverımızda yavaş o yüzden  hızlı değiliz amaaa


bu web sitemize 2 tane genco giriyor bizden gta vice city indircekler

birinci genco
 ve
 ikinici genco aynı anda download tuşuna bastı bu serverları birazda olsa zorlar
 biz go routine ile aralarına zamanlama koydurtuyoruz indir dedikten sonra 2 saniye 3 saniye daha bekliyor

 time.Sleep(2*time.second)

 ile
birinci genconun işi bitti indirdi 2 saniye sonra  diğer genco indiriyor
böylece servera daha az yük biniyor

*/
/*

func main() {

	go func() {
		fmt.Println("1. genco indiriyor")

		time.Sleep(2 * time.Second)

	}()
// verilen zamanı işletim sisteminin scheduling aslgorthimine göre verin
	time.Sleep(2 * time.Second)
	fmt.Println("2. genco indiriyor")
	// our main thread

}
*/

// TUTORIAL EDGE TEN

package main

import (
	"fmt"
	"time"
)

func Hello(name string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello :  %s\n", name)

}

func main() {
	/* 	go func() {
	   		fmt.Println("Hello world")
	   	}()

	   	time.Sleep(2 * time.Second)
	*/

	go Hello("Hilmi")
	// bu go routine main threadını blocklamıyor bloklasa ilk  alttaki fmt çalışmazdı
	fmt.Println("I should be printed first")
	time.Sleep(2 * time.Second)
}
