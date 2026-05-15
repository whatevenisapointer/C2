#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <winsock2.h>
#include <windows.h>

#pragma comment(lib, "ws2_32.lib")



int main()
{
    WSADATA wsa;
    SOCKET s;
    struct sockaddr_in server;
    
    if(WSAStartup(MAKEWORD(2,2), &wsa) != 0)
    {
        printf("[-] Failed Error code: %d", WSAGetLastError());
        return 1;
    }

    if((s = socket(AF_INET, SOCK_STREAM, 0)) == INVALID_SOCKET)
    {
        printf("[-] Erorr creating socket: %d", WSAGetLastError());
        return 1;
    }

    server.sin_addr.s_addr = inet_addr("127.0.0.1");
    server.sin_family = AF_INET;
    server.sin_port = htons(4444);

    int conn = connect(s, (struct sockaddr *)&server, sizeof(server)); 
    if(conn != 0)
    {
        printf("[-] Connection error: %d", WSAGetLastError());
        return 1;
    }

    printf("Connected to teamserver");

    while(1)
    {
    }

    return 0;
}
