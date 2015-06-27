package cm

// Generate CM[string]int.
//go:generate goinline -package=github.com/sasha-s/go-inline/examples/blueprints/concurrentmap --target-package-name=cm_string_int -target-dir=cm_string_int -w Value->int Key->string

// Generate CM[string]CM[string]int.
//go:generate goinline -package=github.com/sasha-s/go-inline/examples/blueprints/concurrentmap --target-package-name=cm -target-dir=. -w Value->github.com/sasha-s/go-inline/examples/usage/cm/cm_string_int::*cm_string_int.CM Key->string

import (
	"fmt"
	"sync"

	"github.com/sasha-s/go-inline/examples/usage/cm/cm_string_int"
)

func ExampleCM() {
	m := New(5)
	var wg sync.WaitGroup
	wg.Add(100)
	// Hammer the map.
	for g := 0; g < 100; g++ {
		go func() {
			for j := 0; j < 10; j++ {
				for k := 10; k < 20; k++ {
					for i := 20; i < 30; i++ {
						inner, _ := m.InsertF(str(i), func() *cm_string_int.CM {
							return cm_string_int.New(5)
						})
						inner.Set(str(j), i+j+k)
					}
				}
			}
			wg.Done()
		}()
	}
	// Throw some concurrent reads into the mix.
	go func() {
		for i := 0; i < 100; i++ {
			inner, ok := m.Get(str(i))
			if !ok {
				continue
			}
			for _, k := range inner.Keys() {
				inner.Get(k)
			}
		}
	}()
	wg.Wait()
	for i := 20; i < 30; i++ {
		inner, _ := m.Get(str(i))
		for j := 0; j < 10; j++ {
			v, ok := inner.Get(str(j))
			fmt.Println(str(i), str(j), v, ok)
		}
	}

	// Output: 00000014 00000000 39 true
	//00000014 00000001 40 true
	//00000014 00000002 41 true
	//00000014 00000003 42 true
	//00000014 00000004 43 true
	//00000014 00000005 44 true
	//00000014 00000006 45 true
	//00000014 00000007 46 true
	//00000014 00000008 47 true
	//00000014 00000009 48 true
	//00000015 00000000 40 true
	//00000015 00000001 41 true
	//00000015 00000002 42 true
	//00000015 00000003 43 true
	//00000015 00000004 44 true
	//00000015 00000005 45 true
	//00000015 00000006 46 true
	//00000015 00000007 47 true
	//00000015 00000008 48 true
	//00000015 00000009 49 true
	//00000016 00000000 41 true
	//00000016 00000001 42 true
	//00000016 00000002 43 true
	//00000016 00000003 44 true
	//00000016 00000004 45 true
	//00000016 00000005 46 true
	//00000016 00000006 47 true
	//00000016 00000007 48 true
	//00000016 00000008 49 true
	//00000016 00000009 50 true
	//00000017 00000000 42 true
	//00000017 00000001 43 true
	//00000017 00000002 44 true
	//00000017 00000003 45 true
	//00000017 00000004 46 true
	//00000017 00000005 47 true
	//00000017 00000006 48 true
	//00000017 00000007 49 true
	//00000017 00000008 50 true
	//00000017 00000009 51 true
	//00000018 00000000 43 true
	//00000018 00000001 44 true
	//00000018 00000002 45 true
	//00000018 00000003 46 true
	//00000018 00000004 47 true
	//00000018 00000005 48 true
	//00000018 00000006 49 true
	//00000018 00000007 50 true
	//00000018 00000008 51 true
	//00000018 00000009 52 true
	//00000019 00000000 44 true
	//00000019 00000001 45 true
	//00000019 00000002 46 true
	//00000019 00000003 47 true
	//00000019 00000004 48 true
	//00000019 00000005 49 true
	//00000019 00000006 50 true
	//00000019 00000007 51 true
	//00000019 00000008 52 true
	//00000019 00000009 53 true
	//0000001a 00000000 45 true
	//0000001a 00000001 46 true
	//0000001a 00000002 47 true
	//0000001a 00000003 48 true
	//0000001a 00000004 49 true
	//0000001a 00000005 50 true
	//0000001a 00000006 51 true
	//0000001a 00000007 52 true
	//0000001a 00000008 53 true
	//0000001a 00000009 54 true
	//0000001b 00000000 46 true
	//0000001b 00000001 47 true
	//0000001b 00000002 48 true
	//0000001b 00000003 49 true
	//0000001b 00000004 50 true
	//0000001b 00000005 51 true
	//0000001b 00000006 52 true
	//0000001b 00000007 53 true
	//0000001b 00000008 54 true
	//0000001b 00000009 55 true
	//0000001c 00000000 47 true
	//0000001c 00000001 48 true
	//0000001c 00000002 49 true
	//0000001c 00000003 50 true
	//0000001c 00000004 51 true
	//0000001c 00000005 52 true
	//0000001c 00000006 53 true
	//0000001c 00000007 54 true
	//0000001c 00000008 55 true
	//0000001c 00000009 56 true
	//0000001d 00000000 48 true
	//0000001d 00000001 49 true
	//0000001d 00000002 50 true
	//0000001d 00000003 51 true
	//0000001d 00000004 52 true
	//0000001d 00000005 53 true
	//0000001d 00000006 54 true
	//0000001d 00000007 55 true
	//0000001d 00000008 56 true
	//0000001d 00000009 57 true

}

func str(i int) string {
	return fmt.Sprintf("%08x", i)
}
