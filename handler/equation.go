package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"web-calculator/calculator"
	"web-calculator/logger"
)

type EquationHandler struct {
	Logger *logger.Logger
}

// handle equation
func (h *EquationHandler) ProcessEquation(w http.ResponseWriter, r *http.Request) {

	// expected data structure
	var body struct {
		Equation string `json:"equation"`
	}

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&body); if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// init eq and failes
	eq := &calculator.Equation{}
	var failed bool

	// parse equation
	eq, failed = calculator.ParseInput(body.Equation); if failed {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid equation"))
		return
	}

	var result float64

	// calculate result
	switch eq.Op {
	case "+" :
		result = eq.Num1 + eq.Num2
	
	case "-" :
		result = eq.Num1 - eq.Num2
	
	case "*" :
		result = eq.Num1 * eq.Num2

	case "/" :
		result = eq.Num1 / eq.Num2
	}

	// log equation
	err = h.Logger.Log(r.Context(), body.Equation, result); if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return result
	w.Write([]byte(fmt.Sprintf("%f", result)))

}