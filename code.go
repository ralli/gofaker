package gofaker

import "fmt"

type Code struct {
	faker *Faker
}

func (code *Code) Base10ISBN() string {
	num := code.faker.Numerify("#########")
	sum := 0
	for i, v := range num {
		val := int(v - '0')
		idx := i + 1
		sum += idx * val
	}
	sum %= 11
	var remainder string
	if sum == 10 {
		remainder = "X"
	} else {
		remainder = fmt.Sprint(sum)
	}

	return num + remainder
}

func (code *Code) Base13ISBN() string {
	num := code.faker.Numerify("############")
	sum := 0
	for i, v := range num {
		val := int(v - '0')
		idx := i + 1
		if idx%2 == 0 {
			sum += val
		} else {
			sum += val * 3
		}
	}

	sum %= 10
	check := (10 - sum) % 10
	remainder := fmt.Sprint(check)
	return num + remainder
}

//
//  public static String base13ISBN() {
//    String num = Base.numerify("############");
//    int sum = 0;
//    for (int i = 0; i < 12; ++i) {
//      int val = num.charAt(i) - '0';
//      int idx = i + 1;
//      if (idx % 2 == 0)
//        sum += val;
//      else
//        sum += val * 3;
//    }
//    sum %= 10;
//    int check = (10 - sum) % 10;
//    String remainder = Integer.toString(check);
//    return num + remainder;
//  }
