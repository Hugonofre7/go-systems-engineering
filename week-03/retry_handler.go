package main

import "fmt"

func main() {
	unreliableOperation := makeUnreliableOperation()

	err := retry(unreliableOperation, 5)

	if err != nil {
		fmt.Printf("Retry failed: %v\n", err)
		return
	}

	fmt.Println("Retry succeeded")

	err = connectToDatabase("db-1")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = connectToDatabase("db-2")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = safeOperation()

	if err != nil {
		fmt.Printf("Recovered error: %v\n", err)
	}

	err = retry(dangerousOperation, 3)

	if err != nil {
		fmt.Printf("Final retry error: %v\n", err)
	}

}

func makeUnreliableOperation() func() error {
	attempts := 0
	return func() error {
		attempts++
		if attempts < 3 {
			return fmt.Errorf("intento %d falló", attempts)
		}
		return nil
	}
}

func retry(op func() error, attempts int) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		if err := runOperation(op); err != nil {
			lastErr = err
			fmt.Printf("intento %d: %v\n", i+1, err)
			continue
		}
		return nil
	}
	return lastErr
}

func runOperation(op func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recuperado: %v", r)
		}
	}()

	return op()
}

func connectToDatabase(name string) error {
	fmt.Printf("Conectando a %s...\n", name)
	defer fmt.Printf("Cerrando conexión a %s\n", name)
	defer fmt.Printf("Liberando lock para %s\n", name)

	if name == "db-2" {
		return fmt.Errorf("conexión rechazada")
	}
	return nil
}

func safeOperation() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recuperado de panic: %v", r)
		}
	}()

	panic("fallo catastrófico simulado")
}

func dangerousOperation() error {
	panic("fallo catastrófico en operación externa")
}
