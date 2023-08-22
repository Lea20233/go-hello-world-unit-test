package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

//use table benchmark

func BenchmarkTableHello(b *testing.B) {

	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Mia",
			request: "Mia",
		},
		{
			name:    "Jonathan",
			request: "Jonathan",
		},
	}

	for _, bench := range benchmark {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}

//use sub benchmark test

func BenchmarkSub(b *testing.B) {

	b.Run("Lala", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lala")
		}
	})
	b.Run("Lui", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lui")
		}
	})

}

// use benchmark testing
func BenchmarkHelloWorldAlex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		HelloWorld("Alex")
	}
}

func BenchmarkHelloWorld(b *testing.B) {

	for i := 0; i < b.N; i++ {
		HelloWorld("Yuna")
	}
}

// use Table Test for SubTest Unit Testing
func TestTableHello(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Maya",
			request:  "Maya",
			expected: "Hello Maya",
		},
		{
			name:     "Lula",
			request:  "Lula",
			expected: "Hello Lula",
		},
		{
			name:     "Alex",
			request:  "Alex",
			expected: "Hello Alex",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := HelloWorld(test.request)
			assert.Equal(t, test.expected, r)
		})
	}

}

// use Sub Test unit testing
func TestSubTest(t *testing.T) {
	t.Run("Hello Luciana", func(t *testing.T) {
		r := HelloWorld("Lucy")
		assert.Equal(t, "Hello Lucy", r)
	})

	t.Run("Hello Aya", func(t *testing.T) {
		r := HelloWorld("Lala")
		assert.Equal(t, "Hello Lala", r)
	})
}

// for execute all unit test func in package helper
func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test is tested.")

	m.Run()

	fmt.Println("After Unit Test was tested.")

}

// use skip for cancel unit testing
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test can't run in Mac OS")
	}
	r := HelloWorld("Lucy")
	assert.Equal(t, "Hello Lucy", r)
}

// use assertion for testing
func TestHelloWorldAssert(t *testing.T) {

	r := HelloWorld("Lucy")
	assert.Equal(t, "Hello Lucy", r)
	fmt.Println("Testing using assertion equal is done.")
}

// use assertion for testing
func TestHelloWorldRequire(t *testing.T) {

	r := HelloWorld("Haha")
	require.Equal(t, "Hello Haha", r)
	fmt.Println("Testing using require equal is done.")
}

// use defer func when there's found error when test is running
func TestHelloWorld(t *testing.T) {
	//r is result
	r := HelloWorld("Luna")

	defer func() {
		if r != "Hello Luna" {
			//for error
			t.Errorf("Testing result is failed.")
		}
	}()
}

// don't use panic func if the test found error when its running
func TestHelloWorldBudi(t *testing.T) {

	r := HelloWorld("Budi")

	if r != "Hello Budi" {
		//for error
		panic("Test is failed.")
	}

}

//using func t.Fail() and t.FailNow() for error

func TestHelloWorld_Maya(t *testing.T) {
	//r is result
	r := HelloWorld("Maya")

	if r != "Hello Maya" {
		//for error
		t.Fail()
	}
	fmt.Println("TestHelloWorldMaya is done tested.")
}

func TestHelloWorld_Nana(t *testing.T) {
	//r is result
	r := HelloWorld("Nana")

	if r != "Hello Nana" {
		//for error
		t.FailNow()
	}
	fmt.Println("TestHelloWorldNana is done tested.")
}

// using func t.Fatal() and t.Error() for error
func TestHelloWorldMaya(t *testing.T) {
	//r is result
	r := HelloWorld("Maya")

	if r != "Hello Maya" {
		//for error
		t.Error("Test result is error. Try Again.")
	}
	fmt.Println("TestHelloWorldMaya is done tested.")
}

func TestHelloWorldNana(t *testing.T) {
	//r is result
	r := HelloWorld("Nana")

	if r != "Hello Nana" {
		//for error
		t.Fatal("Test result is error. Please try Again.")
	}
	fmt.Println("TestHelloWorldNana is done tested.")
}
