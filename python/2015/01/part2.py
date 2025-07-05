BASEMENT_FLOOR = -1

def move_floor(char: str):
    if char == '(':
        return 1
    elif char == ')':
        return -1
    else:
        return 0

with open('input.txt', 'r') as f:
   line = f.read()
   # Index starts at 0 but floor starts at 1
   floor = 1
   for index, paren in enumerate(line):
       floor += move_floor(paren)   
       if floor == -1:
           print(f"Index: {index} Floor: {BASEMENT_FLOOR}")
           break
