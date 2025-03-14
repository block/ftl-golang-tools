package rangeint

func _(i int, s struct{ i int }, slice []int) {
	for i := 0; i < 10; i++ { // want "for loop can be modernized using range over int"
		println(i)
	}
	for i = 0; i < f(); i++ { // want "for loop can be modernized using range over int"
	}
	for i := 0; i < 10; i++ { // want "for loop can be modernized using range over int"
		// i unused within loop
	}
	for i := 0; i < len(slice); i++ { // want "for loop can be modernized using range over int"
		println(slice[i])
	}

	// nope
	for i := 0; i < 10; { // nope: missing increment
	}
	for i := 0; i < 10; i-- { // nope: negative increment
	}
	for i := 0; ; i++ { // nope: missing comparison
	}
	for i := 0; i <= 10; i++ { // nope: wrong comparison
	}
	for ; i < 10; i++ { // nope: missing init
	}
	for s.i = 0; s.i < 10; s.i++ { // nope: not an ident
	}
	for i := 0; i < 10; i++ { // nope: takes address of i
		println(&i)
	}
	for i := 0; i < 10; i++ { // nope: increments i
		i++
	}
	for i := 0; i < 10; i++ { // nope: assigns i
		i = 8
	}
}

func f() int { return 0 }
