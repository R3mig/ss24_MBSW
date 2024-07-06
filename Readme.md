# ModelBasedSW Project SoSe24 P2

## Aufgaben

1. Apply both translation methods to the three examples below. The first example contains a sketch for each translation method.

2. Is it possible to apply both translation methods for each example? Why (not)?

3. What programming language features are required for target programs. Hint: The generic translation method might need to rely on structural subtyping and type assertions.

---

### Aufgabe 1

Siehe
> p2_1.go [Code](https://github.com/R3mig/ss24_MBSW/blob/main/p2_1.go)
---

### Aufgabe 2

1. Die Monomorphization-Methode kann man aus jeder generischen Funktion ableiten. Problem dabei ist nur, dass man für jeden unterstützten Datentypen eine extra Funktion schreiben muss, was zwar den Aufwand immens steigern kann, jedoch Typsicher wäre.

2. sum_G ist nur mit Type assertion umsetzbar, weil über ein interface{} mit for-range nicht einfach so itteriert werden kann. Es ist also nicht wirklich "generisch", aber soll wie sum() nur int und float32 abdecken.

```Go
func sum_G(xs interface{}) interface{} {
	switch xs := xs.(type) {
	case []int:
		var sum int
		for _, v := range xs {
			sum += v
		}
		return sum
	case []float32:
		var sum float32
		for _, v := range xs {
			sum += v
		}
		return sum
	default:
		return nil 
	}
}
```

Anders sieht es mit swap_G() aus.
Der Code

```Go
func swap_G(x interface{}, y interface{}) {
	tmp := *x
	*x = *y
	*y = tmp
}
```

hat aber folgenden Fehler zur Compile-Zeit geworfen:
> invalid operation: cannot indirect x (variable of type interface{})

Vlt hatte ich da einen Denkfehler, also habe ich es ohne Pointer versucht, was eigentlich Sinn ergeben würde, wenn es ja generisch sein soll.
Der Code

```Go
func swap_G(x interface{}, y interface{}) {
	tmp := x
	x = y
	y = tmp
}
```

hat aber folgenden fehler zur Laufzeit verursacht:
> Generic swap:
Before swap: World! Hello
After swap: World! Hello

Es sieht also so aus, dass man über ein interface keine Pointer übergeben kann, weswegen diese Art der übersetzung nicht für diesen Typ von Funktionen anwendbar ist.

---

### Aufgabe 3

Wie bei sum_G bereits angedeutet, werden Konzepte wie Structural subtyping und Type assertion für die Übersetzung benötigt um eine richtige generische behandlung sicherzustellen.

Gerade Type assertion ist essentiel um basierend auf dem Typen dynamisch zur Laufzeit zu entscheiden, wie man mit den Argumenten umgeht.
