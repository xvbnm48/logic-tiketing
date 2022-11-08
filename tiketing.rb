# // you are making a ticketing system, for children under 3 years old, they get in for free

# // this is the ticketing system

ages = [2, 5, 42, 24, 18]
total = 0

ages.each do |age|
    if age < 3
        total += 0
    elsif age < 12
        total += 50
    else
        total += 100
    end
    end

print total