package main

import "fmt"

func minichallange(n int) {
        for i := 1; i <= n; i++ {
		//Looping kondisi i = 1 hingga n
                if i%3 == 0 && i%5 == 0 {
                        fmt.Println("FizzBuzz")
		//Kondisi Jika angka n dapat dibagi dengan 3 dan 5 
		//Maka cetak FizzBuzz
                } else if i%3 == 0 {
                        fmt.Println("Fizz")
				// Jika angka dapat dibagi dengan 3 cetak Fizz
                } else if i%5 == 0 {
                        fmt.Println("Buzz")
				// Jika angka dapat dibagi angka 5 maka cetak Buzz 
                } else {
                        fmt.Println(i)
				//Kondisi else jika tidak memenuhi kondisi dibagi 3, 5, dan 3&5
				//Maka akan mencetak angka i 
                }
        }
}

func main() {
        minichallange(20)
		// program akan looping angka sebanyak 20 kali
}