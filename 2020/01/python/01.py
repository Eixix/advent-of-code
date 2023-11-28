with open('challenge.txt', 'r') as fd:
    reader = fd.readlines()
    for num1 in reader:
        for num2 in reader:
            if int(num1) + int(num2) == 2020:
                print(int(num1) * int(num2))