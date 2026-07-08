def me():
    print("Let's get the python hack started!")

me()

# Python Calculator Mini Program

opt = input("Enter operator (+ - * /): ")
num1 = float(input("Enter first number: "))
num2 = float(input("Enter second number: "))

if opt == "+":
    res = round(num1 + num2)
    print(res)
elif opt == "-":
    res = round(num1 - num2)
    print(res)
elif opt == "*":
    res = round(num1 * num2)
    print(res)
elif opt == "/":
    res = round(num1 / num2)
    print(res)
else :
    print(f"{opt} not valid operator.")