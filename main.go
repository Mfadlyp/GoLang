package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if data == "" {
		return &validationError{Message: "Validation Error"}
	}
	if id != "fadly" {
		return &NotFoundError{Message: "Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// // terjadi error
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation error", validationError.Message)
		// } else if NotFoundError, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("NotFound Error", NotFoundError.Message)
		// } else {
		// 	fmt.Println("Unkown Error", err.Error())
		// }

		// mengunakan switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error", finalError.Error())
		case *NotFoundError:
			fmt.Println("NotFound Error", finalError.Error())
		default:
			fmt.Println("Unkown Error", err.Error())
		}
	} else {
		fmt.Println("SUKSES")
	}
}
