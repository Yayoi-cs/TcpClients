#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>



int main(int argc, char const *argv[]){
  const int port_number = 8080;
  int sock = 0;
  if((sock = socket(AF_INET, SOCK_STREAM,0))< 0){
    printf("socket errorn\n");
    exit(1);
  }
  struct sockaddr_in server;
  server.sin_family = AF_INET;
  server.sin_addr.s_addr =inet_addr("127.0.0.1");
  server.sin_port = htons(port_number);
  if ((connect(sock, (struct sockaddr *)&server, sizeof(server))) < 0) {
    printf("connection failed\n");
    exit(1);
  }
  char buff[1024] = {0};
  while(1){
    recv(sock, buff, sizeof(buff),0);
    printf("Recved Message : %s\n",buff);
    if (strstr(buff, "FLAG") != NULL) {
      break;
    } else if (strstr(buff,"??") == NULL) {
      continue;
    }
    char *ptr = strtok(buff," ");
    char *parts = NULL;
    while (ptr != NULL) {
      if(strstr(ptr,"+")!=NULL) {
        parts = ptr;
        break;
      }
      ptr = strtok(NULL," ");
    }
    if (parts != NULL) {
      int result = 0;
      ptr = strtok(parts, "+");
      while (ptr != NULL) {
        result += atoi(ptr);
        ptr = strtok(NULL,"+");
      }
      printf("My response is %d\n",result);
      send(sock,&result,sizeof(result),0);
    }
  }
  close(sock);
  return 0;
}
