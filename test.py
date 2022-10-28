import os
arr = os.listdir()
arr.sort()
print(arr)
# name = name.replace("[SuperSliv.biz] ", "")

def rec(arr):
    for i in range(len(arr)):
        print(os.getcwd())
        try:
            arr_new = os.listdir(arr[i])
            os.chdir(arr[i])
            rec(arr_new)
            new_name = arr[i].replace("[SuperSliv.biz] ", "")
            os.rename(arr[i], new_name)
            print("Renamed FOLDER with FILES!")
        except:
            print("file found! name:", arr[i])
            if arr[i] != "test.py":
                new_name = arr[i].replace("[SuperSliv.biz] ", "")
                os.rename(arr[i], new_name)
                print("Renamed FILE or EMPTY FOLDER!")
        
    os.chdir("..")
    print(os.getcwd())
        # arr[i] = arr[i].replace("[SuperSliv.biz] ", "")


rec(arr)


# os.getcwd()
# os.chdir("..")
