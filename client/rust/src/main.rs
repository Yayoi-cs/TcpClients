use std::io::{Read, Write};
use std::net::TcpStream;
use std::i32;

fn main() -> std::io::Result<()> {
    let mut stream = TcpStream::connect("localhost:8080")?;
    println!("Connected to Server : localhost:8080");
    let mut buffer = [0; 1024];
    let mut index = 0;
    loop {
        let mut bytes_read = stream.read(&mut buffer)?;
        let mut response = String::from_utf8_lossy(&buffer[..bytes_read]);
        println!("Response from server: {}", response);
        if index == 10 {
            break;
        }
        if !response.contains("??") {
            continue;
        }
        let result = calc(&response);
        println!("My response is {}", result);
        stream.write_all(&result.to_le_bytes())?;
        index +=1;
    }
    Ok(())
}

fn calc(response: &str) -> i32 {
    let tokens: Vec<&str> = response.split_whitespace().collect();
    let equation = tokens.iter().find(|&&token| token.contains('+')).unwrap();
    let operands: Vec<&str> = equation.split('+').collect();
    let left_operand: i32 = operands[0].parse().unwrap();
    let right_operand: i32 = operands[1].parse().unwrap();
    let result = left_operand + right_operand;
    return result;
}