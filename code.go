package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// сначала надо ввести количество файлов базы данных
// название файлов базы данных от 1.jpeg до 'k'.jpeg по пути images/
// название файлов обработки от 1.jpeg до 'k'.jpeg по пути imagesOb/
// все данные о них - var histogramall [10][16][4]int

func main() {
	jpeg()
	var k int = jpeg()
	jpeg2()
	var histogramall [][16][4]int
	for i := 0; i < k; i++ {

		a := strconv.Itoa(i + 1)
		name := "images/" + a + ".jpeg"
		reader, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()

		m, _, _ := image.Decode(reader)
		bounds := m.Bounds()
		var histogram [16][4]int
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				// A color's RGBA method returns values in the range [0, 65535].
				// Shifting by 12 reduces this to the range [0, 15].
				histogram[r>>12][0]++
				histogram[g>>12][1]++
				histogram[b>>12][2]++
				histogram[a>>12][3]++
			}
		}
		histogramall = append(histogramall, histogram)
	}
	// сравнивание данных изображений с базой
	fmt.Println(1)
	histogramall2 := main2()
	var j int
	var u int
	var f bool
	for j = 0; j < k; j++ {
		f = false
		for u = 0; u < k; u++ {
			if histogramall2[j] == histogramall[u] {
				f = true
			}
		}
		if f {
			fmt.Println("Отпечаток " + strconv.Itoa(j) + " найден в базе данных")
		} else {
			fmt.Println("Отпечаток " + strconv.Itoa(j) + " ненайден в базе данных")
		}
	}
	end := make(chan struct{})
	<-end

}

func main2() [][16][4]int {
	jpeg2()
	var k int = jpeg2()
	var histogramall [][16][4]int
	for i := 0; i < k; i++ {

		a := strconv.Itoa(i + 1)
		name := "imagesOb/" + a + ".jpeg"
		reader, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
		m, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		bounds := m.Bounds()
		var histogram [16][4]int
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				// A color's RGBA method returns values in the range [0, 65535].
				// Shifting by 12 reduces this to the range [0, 15].
				histogram[r>>12][0]++
				histogram[g>>12][1]++
				histogram[b>>12][2]++
				histogram[a>>12][3]++
			}
		}
		histogramall = append(histogramall, histogram)
	}
	return histogramall

}
func jpeg() int {
	files, err := ioutil.ReadDir("images")
	if err != nil {
		log.Fatal(err)
	}
	var k int
	var nam string
	for _, file := range files {
		nam = "images/" + file.Name()
		k++
		err := os.Rename(nam, "images/"+strconv.Itoa(k)+".jpeg")
		if err != nil {
			fmt.Println(err)
		}

		// imgIn, _ := os.Open("images/" + strconv.Itoa(k) + ".jpeg")
		// imgJpg, _, _ := image.Decode(imgIn)
		// imgIn.Close()
		// imgJpg = resize.Resize(440, 500, imgJpg, resize.Bicubic) // <-- Собственно изменение размера картинки
		// imgOut, _ := os.Create("images/" + strconv.Itoa(k) + ".jpeg")
		// imgOut.Close()
	}
	return k
}

func jpeg2() int {
	files, err := ioutil.ReadDir("imagesOb")
	if err != nil {
		log.Fatal(err)
	}
	var k int
	var nam string
	for _, file := range files {
		nam = "imagesOb/" + file.Name()
		k++
		err := os.Rename(nam, "imagesOb/"+strconv.Itoa(k)+".jpeg")
		if err != nil {
			fmt.Println(err)
		}
		// imgIn, _ := os.Open("imagesOb/" + strconv.Itoa(k) + ".jpeg")
		// imgJpg, _, _ := image.Decode(imgIn)
		// imgIn.Close()
		// imgJpg = resize.Resize(440, 500, imgJpg, resize.Bicubic) // <-- Собственно изменение размера картинки
		// imgOut, _ := os.Create("imagesOb/" + strconv.Itoa(k) + ".jpeg")
		// imgOut.Close()
	}
	return k
}
