import socket
import struct

HOST = 'localhost'
PORT = 8080

def main():
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        
        while True:
            res = s.recv(4096).decode()
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
            result = sum(int(num_str) for num_str in nums)
            print("My response is", result)
            print("In byte:", result.to_bytes(8, "little"))
            
            s.sendall(struct.pack("<Q", result))

if __name__ == "__main__":
    main()
