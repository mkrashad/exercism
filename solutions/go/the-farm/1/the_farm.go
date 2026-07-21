package thefarm
import ("errors"
        "fmt")

// TODO: define the 'DivideFood' function
func DivideFood(fl FodderCalculator, numCows int) (float64, error) {
    amount, err := fl.FodderAmount(numCows)
    if err != nil {
        return 0, err
    }
    factor, err := fl.FatteningFactor()
     if err != nil {
        return 0, err
    }
    
    result := float64(amount) / float64(numCows) * factor
    return result, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fl FodderCalculator, numCows int) (float64, error){
    if numCows <= 0 {
    	return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fl, numCows)
}
// TODO: define the 'ValidateNumberOfCows' function
type MyCustomErr struct {
    numCows int
    message string
}

func (e *MyCustomErr) Error() string {
    return  fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

func ValidateNumberOfCows(numCows int) error{
    if numCows < 0 {
    	return &MyCustomErr{
            numCows: numCows,
        	message: "there are no negative cows",
        }
    }

    if numCows == 0 {
    	return &MyCustomErr{
            numCows: numCows,
        	message: "no cows don't need food",
        }
    }
    return nil 
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
