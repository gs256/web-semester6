let register1 = ""
let register2 = ""
let expressionResult = ""
let selectedOperation = null

outputElement = document.getElementById("result")
digitButtons = document.querySelectorAll('[id ^= "btn_digit_"]')

function onDigitButtonClicked(digit) {
  if (!selectedOperation) {
    if (digit != "." || (digit == "." && !register1.includes(digit))) {
      register1 += digit
    }
    outputElement.innerHTML = register1
  } else {
    if (digit != "." || (digit == "." && !register2.includes(digit))) {
      register2 += digit
      outputElement.innerHTML = register2
    }
  }
}

digitButtons.forEach(button => {
  button.onclick = function () {
    const digitValue = button.innerHTML
    onDigitButtonClicked(digitValue)
  }
})

document.getElementById("btn_op_mult").onclick = () => {
  if (register1 === "") return
  selectedOperation = "x"
}

document.getElementById("btn_op_plus").onclick = () => {
  if (register1 === "") return
  selectedOperation = "+"
}

document.getElementById("btn_op_minus").onclick = () => {
  if (register1 === "") return
  selectedOperation = "-"
}

document.getElementById("btn_op_div").onclick = () => {
  if (register1 === "") return
  selectedOperation = "/"
}

document.getElementById("btn_op_clear").onclick = () => {
  register1 = ""
  register2 = ""
  selectedOperation = ""
  expressionResult = ""
  outputElement.innerHTML = 0
}

document.getElementById("btn_op_equal").onclick = () => {
  if (register1 === "" || register2 === "" || !selectedOperation) return

  switch (selectedOperation) {
    case "x":
      expressionResult = +register1 * +register2
      break
    case "+":
      expressionResult = +register1 + +register2
      break
    case "-":
      expressionResult = +register1 - +register2
      break
    case "/":
      expressionResult = +register1 / +register2
      break
  }

  register1 = expressionResult.toString()
  register2 = ""
  selectedOperation = null

  outputElement.innerHTML = register1
}

document.getElementById("btn_op_sign").onclick = () => {
  if (register1 === "") return
  if (register1 === "0") return
  if (register1 === ".") return

  if (register1.startsWith("-")) {
    register1 = register1.substring(1)
  } else {
    register1 = "-" + register1
  }
  outputElement.innerHTML = register1
}
