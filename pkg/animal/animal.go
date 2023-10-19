package animal

var animalType = []string{
	"Bird",
	"Mammal",
	"Reptile",
	"Amphibian",
	"Fish",
	"Insect",
}

type AnimalType int

const (
	Bird AnimalType = 1 + iota
	Mammal
	Reptile
	Amphibian
	Fish
	Insect
)

// method String() disini dibuat berdasarkan kontrak atau interface dari
// fmt.Stringer, jadi beberapa library khusus nya fmt akan mencari dan menjalankan
// method ini untuk mendapatkan nilai sebenarnya (string) dari sebuah type atau nilai,
// ini juga berlaku untuk interface error. Untuk kasus sekarang method String() adalah implementasi
// method dari type AnimalType digunakan untuk mengembalikan nilai jenis hewan berdasarkan indeks yang
// diberikan pada parameter. Oleh karena itu diperlukan const dengan type yang sama dimana
// untuk mewakili angka indeks
func (a AnimalType) String() string {
	return animalType[a-1]
}

type Animal struct {
	Name string
	Type AnimalType
}

// implementasi method animal() sesuai dengan interface IAnimal
func (a Animal) animal() (name string, t AnimalType) {
	return a.Name, a.Type
}

type IAnimal interface {
	animal() (name string, t AnimalType)
}

func GetNameAndType(a IAnimal) (name string, t AnimalType) {
	return a.animal()
}
