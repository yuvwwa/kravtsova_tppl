file_name = open("test.txt", 'r').readlines()

lines = len(file_name)
chars = sum(len(line) for line in file_name)
empty = sum(1 for line in file_name if line.strip() == "")

freq_dict ={}
for line in file_name:
    for char in line:
        if char in freq_dict:
            freq_dict[char] +=1
        else:
            freq_dict[char] =1

print("1 - количество строк в файле")
print("2 - количество символов в файле")
print("3 - количество пустых строк")
print("4 - частотный словарь символов")

choice = input("Введите ваш выбор: ").split()
if "1" in choice:
    print(f"количество строк в файле: {lines}")
elif "2" in choice:
    print(f"количество символов в файле: {chars}")
elif "3" in choice:
    print(f"количество пустых строк: {empty}")
else:
    print(f"частотный словарь символов: ")
    for char, count in freq_dict.items():
        res = char if char!="\n" else "\\n"
        print(f"{res}: {count}")