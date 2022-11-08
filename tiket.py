# you are making a ticketing system, for children under 3 years old, they get in for free
# for children between 3 and 12 years old, they pay 14.00
# for adults, they pay 23.00
# create a program that asks the user how old they are, and then tell them the cost of their ticket



ages = {2, 5, 42, 24, 18}
total = 0
for age in ages:
    if age < 3:
        total += 0
        continue
    else:
        total += 100
print(total)

# using while loop

# while ages < 3:
#     total += 0
#     continue
# else:
#     total += 100
# print(total)
