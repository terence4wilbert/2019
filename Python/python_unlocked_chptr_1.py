

i = "asdf"
j = "asdf"

id(i) == id(j)

ii = 100000000000
jj = 100000000000

id(jj) == id(ii)

class Kls:
    pass

k = Kls()
j = Kls()

id(k) == id(j)


# BAD
print('a' +' ' + 'simple' + ' ' + 'sentence' + ' ' +'')

#good
print(' '.join(['a', 'simple', 'sentence', '.']))


class C:
    def __init__(self, arg1, arg2):
        self.str = arg1
        self.lst = arg2

iC = C("arun", [1,2])
print("iC.str:")
print(iC.str)
print("iC.lst:")
print(iC.lst)
print("iC.lst.append(4):")
iC.lst.append(4)
print(iC.lst)

g = []

# should be the same
g.__class__
print(type(g))

#type is instance's class
class M:
    def __init__(self, d):
        self.d = d
    def square(self):
        return self.d * self.d

class N:
    def __init__(self, d):
        self.d = d
    def cube(self):
        return self.d * self.d * self.d

m = M(4)
print(type(m))
print(f'm squared = {m.square()}')
m.__class__ = N
print('set m.__class__ = N')
print(f'm cubed = {m.cube()}')
print(type(m))


