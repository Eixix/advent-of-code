def main():
  lines = []

  with open('input.csv', 'r') as file:
    lines = file.readlines()
  
  lines = [line.rstrip() for line in lines]

  # Rotate the list clockwise
  lines = list(reversed(list(reversed(list(zip(*lines))))))

  # Convert tuples to list
  lines = [list(elem) for elem in lines]

  gamma_rate_binary_string = ""

  for line in lines:
    frequency_dict = {}
    for character in line:
      if character in frequency_dict:
          frequency_dict[character] += 1
      else:
          frequency_dict[character] = 1

    gamma_rate_binary_string += max(frequency_dict, key=frequency_dict.get)

  gamma_rate = int(gamma_rate_binary_string, 2)

  epsilon_rate = int(''.join(['1' if i == '0' else '0' for i in gamma_rate_binary_string]), 2)
  print(gamma_rate * epsilon_rate)

if __name__ == '__main__':
  main()