from pwn import *

def main():
    p = remote("localhost",8080)

    while True:
        res = p.recv().decode()
        print(res)
        if "FLAG" in res:
            break
        if "??" not in res:
            continue
        str_list = res.split(" ")
        parts = ""
        for part in str_list:
            if "+" in part:
                parts = part
                break
        nums = parts.split("+")
        result = 0
        for num_str in nums:
            num = int(num_str)
            result += num
        print("My response is",result)
        print("In byte : ",result.to_bytes(8,"little"))
        p.sendline(result.to_bytes(8,"little"))

if __name__ == "__main__":
    main()
