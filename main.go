package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file1, err1 := os.Create("file1.bin") // Создаем file1.bin, и занасим в него данные, которые потом нужно будет прочитать и copy
	if err1 != nil {
		log.Fatal(err1)
	}
	bin := []byte("File is bin")
	file1.Write(bin)
	file1.Close()

	file2, err2 := os.Create("file2.bin") // создаем file2.bin, в который будем копировать данные из file1.bin
	if err2 != nil {
		log.Fatal(err2)
	}
	file11, err22 := os.Open("file1.bin") // открываем file1.bin и присваемваем значение переменной file11, тут вот долго тупил и понял, что если не открыть этот файл снова и не присвоить его переменной то нельзя будет скопирваоть его значения, я сначала думал, что открывать не надо и то что я создал выше уже считается и можно просто скопировать
	if err22 != nil {
		log.Fatal(err22)
	}
	io.Copy(file2, file11) // копируем данные из file1.bin в переменную file2
	file2.Close()

	file3, err3 := os.Open("file2.bin") // записываем данные из file2.bin в переменную file3
	if err3 != nil {
		log.Fatal(err3)
	}
	bin3 := make([]byte, 64) // создаем буффер, в который потмо будем читать данные из file3
	for {
		n, err := file3.Read(bin3)
		if err == io.EOF {
			break
		}
		fmt.Println(string(bin3[:n]))
	}
	file3.Close()
}
