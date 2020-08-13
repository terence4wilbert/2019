# function
print(f'Function: a function returns a value based on the given arguments.')
def add_one(number):
    return number + 1

print(add_one(2))

print()

print('Functions are First class object')
print('First-class object: functions can be passed around and used as arguments, just like any other object (string, int, float, list, and so on)')

# Regular functions
def say_hello(name):
    return f'Hello {name}'
def be_awesome(name):
    return f'Yo {name}, together we are the awesomest!'

#expects a function as an arguement, can pass it the two REGULAR functions from above

def greet_bob(greeter_func):
    return greeter_func('Bob')


print('greet_bob(say_hello), no parentheses on function')
print(greet_bob(say_hello))

print('greet_bob(be_awesome), no parentheses on function')
print(greet_bob(be_awesome))
print()
print('No parentheses on function means only REFERENCE to function is passed')
print()

print('Inner Functions')
print('two inner functions')
def parent():
    print("Printing from the parent() function")
# functions first_child and second child will only be available in the parent() function scope
    def first_child():
        print('Printing from the first child() function')

    def second_child():
        print('Printing from the second_child() function')

    second_child()
    first_child()

print('Default value of a function is \'None\'')
print('calling parent() function')
parent()

# will not work
# second_child()
print()
print('Functions can be used as return values as well')

def parent(num):
    def first_child():
        return "Hi, I am Emma"
    def second_child():
        return "Call me Liam"
    if num == 1:
        # no parentheses, passing a reference
        return first_child
    else:
        # no parentheses, passing a reference
        return second_child

first = parent(1)
second = parent(2)
print('References to objects function an memory location')
print(first)
print(second)
print('functions being executed return str text')
print(first())
print(second())

print()
print('Decorators')
def my_decorator(func):
    def wrapper():
        print("Something is happening before the function is called.")
        func()
        print("Something is happening after the function is called ")
    return wrapper

def say_whee():
    print('Whee!')

say_whee = my_decorator(say_whee)

say_whee()


def my_decorator(func):
    def wrapper():
        print("Something is happening before the function is called.")
        func()
        print("Something is happening after the function is called.")
    return wrapper

@my_decorator
def say_whee():
    print("Whee!")


# reusing decorators
def do_twice(func):
    def wrapper_do_twice():
        func()
        func()
    return wrapper_do_twice()

@do_twice
def say_whee():
    print("Whee!")

say_whee()

# Decorators with arguments
def do_twice(func):
    def wrapper_do_twice(*args, **kwargs):
        func(*args, **kwargs)
        func(*args, **kwargs)
    return wrapper_do_twice()

@do_twice
def greet(name):
    print(f'Hello {name}')
