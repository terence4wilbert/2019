import argparse
parser = argparse.ArgumentParser()
parser.add_argument('integers', metavar='N', type=int, nargs='+')
parser.add_argument('-f','--foo', help='foo bar')
parser.add_argument('-b','--bar',help='bar help')
parser.add_argument('-z', '--zar', help='zar help')
parser.add_argument('-t', '--toggle', action="store_true")
parser.add_argument('-x', '--exclude', action="store_false")
parser.add_argument('-s', '--start', action="store_true")
parser.add_argument('-xx','--kill',help='turn it all off')
parser.add_argument("square", help="display a square of a given number", type=int)
parser.add_argument("--verbosity", help="increase output verbosity")
parser.add_argument("-v", "--verbose", help="increase output verbosity with no arguements", action="store_true")
args = parser.parse_args()
answer = args.square**2
# print(args.echo)
# print(args.square**2)
if args.verbosity:
    print("verbosity turned on")
if args.verbose:
    print("verbose turned on")
    print(f'the square of {args.square} equals {answer}')
else:
    print(answer)
print(args)