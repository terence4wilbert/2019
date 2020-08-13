#! /usr/local/bin/python3

### EXAMPLE PYTHON MODULE
# Define some variables:

number_one = 1
age = 78

# define some functions
def print_hello():
    print("hello")

def times_four(input):
    print(input*4)


# define a class
class house:
    def __init__(self):
        self.type = input("What type of house?")
        self.height = input("What height (in feet)?")
        self.price = input("How much did it cost? ")
        self.age = input("How old is it (in years)?")

    def print_details(self):
        print(f'This house is a/an {self.height} foot')
        print(f'{self.type}, hose {self.age} years old and costing {self.price} dollars')