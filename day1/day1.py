nums = []
with open("input.txt", "r") as file:
    for linea in file:
        linea = linea.strip()

        if linea == "":
            n = linea.replace("", ",")
            nums.append(n)
        else:
            nums.append(linea)

print(nums)
