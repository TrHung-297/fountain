/* !!
 * File: client.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:54:03 am
 
 */

package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/middleware/examples/zproto"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	conn, err := grpc.Dial("127.0.0.1:22345", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := zproto.NewChatTestClient(conn)
	sess := &zproto.ChatSession{base.Int64ToString(rand.Int63())}
	fmt.Println("sessionId : ", sess.SessionId)

	var message string
	for {
		fmt.Print("> ")
		if n, err := fmt.Scanln(&message); err == io.EOF {
			return
		} else if n > 0 {
			if message == "quit" {
				return
			} else {
				_, err := client.SendChat(context.Background(), &zproto.ChatMessage{SenderSessionId: sess.SessionId, MessageData: message})
				if err != nil {
					fmt.Printf("%v.SendChat(_) = _, %v\n", client, err)
				}
			}
		}
	}
}
