/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/5/8 16:59
*/

package contractInterface

import (
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	balance, err := GetErc20Balance("", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(balance.String())
}

func TestUploadOpenPosition(t *testing.T) {
	UploadOpenPosition("")

}

func TestUploadOutPosition(t *testing.T) {
	UploadOutPosition("")

}
