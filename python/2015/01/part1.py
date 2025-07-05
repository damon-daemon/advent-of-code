def move_floor(char: str):
    if char == '(':
        return 1
    elif char == ')':
        return -1
    else:
        return 0

with open('input.txt', 'r') as f:
   line = f.read()
   floor = 0
   for paren in line:
       floor += move_floor(paren)
   print(f"Final floor: {floor}")
