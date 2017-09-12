package gofaker

import "fmt"

func ExampleCompass_Cardinal() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.Cardinal())
	// Output: east
}

func ExampleCompass_Ordinal() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.Ordinal())
	// Output: southeast
}

func ExampleCompass_HalfWind() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.HalfWind())
	// Output: east-northeast
}

func ExampleCompass_QuarterWind() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.QuarterWind())
	// Output: northeast by north
}

func ExampleCompass_Direction() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.Direction())
	// Output: northwest
}

func ExampleCompass_Abbreviation() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.Abbreviation())
	// Output: NW
}

func ExampleCompass_Azimuth() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.Azimuth())
	// Output: 315
}

func ExampleCompass_CardinalAbbreviation() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.CardinalAbbreviation())
	// Output: E
}

func ExampleCompass_OrdinalAbbreviation() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.OrdinalAbbreviation())
	// Output: SE
}

func ExampleCompass_HalfWindAbbreviation() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.HalfWindAbbreviation())
	// Output: ENE
}

func ExampleCompass_QuarterWindAbbreviation() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.QuarterWindAbbreviation())
	// Output: NEbN
}

func ExampleCompass_CardinalAzimuth() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.CardinalAzimuth())
	// Output: 90
}

func ExampleCompass_OrdinalAzimuth() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.OrdinalAzimuth())
	// Output: 135
}

func ExampleCompass_HalfWindAzimuth() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.HalfWindAzimuth())
	// Output: 67.5
}

func ExampleCompass_QuarterWindAzimuth() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Compass.QuarterWindAzimuth())
	// Output: 33.75
}
