/*
 * @Descripttion:
 * @Version: 1.0
 * @Author: ouyeye 1544763622@qq.com
 * @Date: 2024-12-18 15:53:55
 */
package proto

import "context"

type HelloServiceServerImp struct {
}

var _ HelloServiceServer = (*HelloServiceServerImp)(nil)

func (h *HelloServiceServerImp) SayHello(ctx context.Context, userName *String) (*String, error) {
	res := &String{Value: "Hello " + userName.GetValue()}
	return res, nil
}

func (h *HelloServiceServerImp) mustEmbedUnimplementedHelloServiceServer() {

}
